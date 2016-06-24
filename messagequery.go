// ADDED BY DROP - https://github.com/matryer/drop (v0.6)
//  source: github.com/maciekmm/messenger-platform-go-sdk (ca9227b956ad50bc8b6225a464f6c0146887f7c5)
//  update: drop -f github.com/maciekmm/messenger-platform-go-sdk
// license: The MIT License (MIT) (see repo for details)

package main

import (
	"errors"
)

type SendMessage struct {
	Text       string      `json:"text,omitempty"`
	Attachment *Attachment `json:"attachment,omitempty"`
}

// Recipient describes the person who will receive the message
// Either ID or PhoneNumber has to be set
type Recipient struct {
	ID          string `json:"id,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
}

// NotificationType describes the behavior phone will execute after receiving the message
type NotificationType string

const (
	// NotificationTypeRegular will emit a sound/vibration and a phone notification
	NotificationTypeRegular NotificationType = "REGULAR"
	// NotificationTypeSilentPush will just emit a phone notification
	NotificationTypeSilentPush NotificationType = "SILENT_PUSH"
	// NotificationTypeNoPush will not emit sound/vibration nor a phone notification
	NotificationTypeNoPush NotificationType = "NO_PUSH"
)

type MessageQuery struct {
	Recipient        Recipient        `json:"recipient"`
	Message          SendMessage      `json:"message"`
	NotificationType NotificationType `json:"notification_type,omitempty"`
}

func (mq *MessageQuery) RecipientID(recipientID string) error {
	if mq.Recipient.PhoneNumber != "" {
		return errors.New("Only one user identification (phone or id) can be specified.")
	}
	mq.Recipient.ID = recipientID
	return nil
}

func (mq *MessageQuery) RecipientPhoneNumber(phoneNumber string) error {
	if mq.Recipient.ID != "" {
		return errors.New("Only one user identification (phone or id) can be specified.")
	}
	mq.Recipient.PhoneNumber = phoneNumber
	return nil
}

func (mq *MessageQuery) Notification(notification NotificationType) *MessageQuery {
	mq.NotificationType = notification
	return mq
}

func (mq *MessageQuery) Text(text string) error {
	if mq.Message.Attachment == nil {
		mq.Message.Attachment = &Attachment{}
	}
	if mq.Message.Attachment != nil && mq.Message.Attachment.Type == AttachmentTypeTemplate {
		return errors.New("Can't set both text and template.")
	}
	mq.Message.Text = text
	return nil
}

func (mq *MessageQuery) resource(typ AttachmentType, url string) error {
	if mq.Message.Attachment == nil {
		mq.Message.Attachment = &Attachment{}
	}
	if mq.Message.Attachment.Payload != nil {
		return errors.New("Attachment already specified.")
	}
	mq.Message.Attachment.Type = typ
	mq.Message.Attachment.Payload = &Resource{URL: url}
	return nil
}

func (mq *MessageQuery) Audio(url string) error {
	return mq.resource(AttachmentTypeAudio, url)
}

func (mq *MessageQuery) Video(url string) error {
	return mq.resource(AttachmentTypeVideo, url)
}

func (mq *MessageQuery) Image(url string) error {
	return mq.resource(AttachmentTypeImage, url)
}
