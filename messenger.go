// ADDED BY DROP - https://github.com/matryer/drop (v0.6)
//  source: github.com/maciekmm/messenger-platform-go-sdk (ca9227b956ad50bc8b6225a464f6c0146887f7c5)
//  update: drop -f github.com/maciekmm/messenger-platform-go-sdk
// license: The MIT License (MIT) (see repo for details)

package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	//GraphAPI specifies host used for API requests
	GraphAPI = "https://graph.facebook.com"
)

// MessageReceivedHandler is called when a new message is received
type MessageReceivedHandler func(Event, MessageOpts, ReceivedMessage)

// MessageDeliveredHandler is called when a message sent has been successfully delivered
type MessageDeliveredHandler func(Event, MessageOpts, Delivery)

// PostbackHandler is called when the postback button has been pressed by recipient
type PostbackHandler func(Event, MessageOpts, Postback)

// AuthenticationHandler is called when a new user joins/authenticates
type AuthenticationHandler func(Event, MessageOpts, *Optin)

// Messenger is the main service which handles all callbacks from facebook
// Events are delivered to handlers if they are specified
type Messenger struct {
	VerifyToken      string
	AppSecret        string
	AccessToken      string
	PageID           string
	MessageReceived  MessageReceivedHandler
	MessageDelivered MessageDeliveredHandler
	Postback         PostbackHandler
	Authentication   AuthenticationHandler
}

// Handler is the main HTTP handler for the Messenger service.
// It MUST be attached to some web server in order to receive messages
func (m *Messenger) Handler(rw http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		query := req.URL.Query()
		verifyToken := query.Get("hub.verify_token")
		log.Println("Handle", req.Method, " token=", verifyToken, " verify_token=", m.VerifyToken)
		log.Println("Equal:", verifyToken == m.VerifyToken)
		if verifyToken != m.VerifyToken {
			rw.WriteHeader(http.StatusUnauthorized)
			log.Println("StatusUnauthorized")
			return
		}
		rw.WriteHeader(http.StatusOK)
		log.Println("RET:", query.Get("hub.challenge"))
		rw.Write([]byte(query.Get("hub.challenge")))
	} else if req.Method == "POST" {
		m.handlePOST(rw, req)
	} else {
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (m *Messenger) handlePOST(rw http.ResponseWriter, req *http.Request) {
	read, err := ioutil.ReadAll(req.Body)

	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	//Message integrity check
	if m.AppSecret != "" {
		if req.Header.Get("x-hub-signature") == "" || !checkIntegrity(m.AppSecret, read, req.Header.Get("x-hub-signature")[5:]) {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	event := &upstreamEvent{}
	err = json.Unmarshal(read, event)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, entry := range event.Entries {
		for _, message := range entry.Messaging {
			if message.Delivery != nil {
				if m.MessageDelivered != nil {
					go m.MessageDelivered(entry.Event, message.MessageOpts, *message.Delivery)
				}
			} else if message.Message != nil {
				if m.MessageReceived != nil {
					go m.MessageReceived(entry.Event, message.MessageOpts, *message.Message)
				}
			} else if message.Postback != nil {
				if m.Postback != nil {
					go m.Postback(entry.Event, message.MessageOpts, *message.Postback)
				}
			} else if m.Authentication != nil {
				go m.Authentication(entry.Event, message.MessageOpts, message.Optin)
			}
		}
	}
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(`{"status":"ok"}`))
}

func checkIntegrity(appSecret string, bytes []byte, expectedSignature string) bool {
	mac := hmac.New(sha1.New, []byte(appSecret))
	mac.Write(bytes)
	if fmt.Sprintf("%x", mac.Sum(nil)) != expectedSignature {
		return false
	}
	return true
}

func (m *Messenger) doRequest(method string, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	query := req.URL.Query()
	query.Set("access_token", m.AccessToken)
	req.URL.RawQuery = query.Encode()
	return http.DefaultClient.Do(req)
}
