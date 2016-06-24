// ADDED BY DROP - https://github.com/matryer/drop (v0.6)
//  source: github.com/maciekmm/messenger-platform-go-sdk (ca9227b956ad50bc8b6225a464f6c0146887f7c5)
//  update: drop -f github.com/maciekmm/messenger-platform-go-sdk
// license: The MIT License (MIT) (see repo for details)

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Profile struct holds data associated with Facebook profile
type Profile struct {
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	ProfilePicture string `json:"profile_pic,omitempty"`
	Locale         string `json:"locale,omitempty"`
	Timezone       int    `json:"timezone,omitempty"`
	Gender         string `json:"gender,omitempty"`
}

// GetProfile fetches the recipient's profile from facebook platform
// Non empty UserID has to be specified in order to receive the information
func (m *Messenger) GetProfile(userID string) (*Profile, error) {
	resp, err := m.doRequest("GET", fmt.Sprintf(GraphAPI+"/v2.6/%s?fields=first_name,last_name,profile_pic,locale,timezone,gender", userID), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	read, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		er := new(rawError)
		json.Unmarshal(read, er)
		return nil, errors.New("Error occured: " + er.Error.Message)
	}
	profile := new(Profile)
	return profile, json.Unmarshal(read, profile)
}
