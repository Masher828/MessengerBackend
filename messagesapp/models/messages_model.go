package models

import (
	"github.com/Masher828/MessengerBackend/common-packages/system"
)

type Message struct {
	Id             string  `json:"id" bson:"_id"`
	UserId         int64   `json:"userId" bson:"userId"`
	ConversationId string  `json:"conversationId" bson:"conversationId"`
	Type           string  `json:"type" bson:"type"`
	Body           string  `json:"body" bson:"body"`
	Status         string  `json:"status" bson:"status"`
	DeletedFor     []int64 `json:"deletedFor" bson:"deletedFor"`
	SentOn         int64   `json:"sentOn" bson:"sentOn"`
	DeliveredOn    int64   `json:"deliveredOn" bson:"deliveredOn"`
}

type MessageRequest struct {
	ConversationId string `json:"conversationId" bson:"conversationId"`
	Type           string `json:"type" bson:"type"`
	Body           string `json:"body" bson:"body"`
}

func (messageRequest *MessageRequest) GetMessage() *Message {
	var message Message

	message.Body = messageRequest.Body
	message.Type = messageRequest.Type
	message.ConversationId = messageRequest.ConversationId

	return &message
}

func (message *MessageRequest) Isvalid() (bool, error) {

	if len(message.Body) == 0 {
		return false, system.InvalidPayloadData
	}

	if len(message.ConversationId) == 0 {
		return false, system.InvalidPayloadData
	}

	return true, nil
}
