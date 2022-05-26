package models

import (
	"github.com/Masher828/MessengerBackend/common-packages/system"
)

type Message struct {
	UserId         int64   `json:"userId" bson:"userId"`
	ConversationId string  `json:"conversationId" bson:"conversationId"`
	Type           string  `json:"type" bson:"type"`
	Body           string  `json:"body" bson:"body"`
	Status         string  `json:"status" bson:"status"`
	DeletedFor     []int64 `json:"deletedFor" bson:"deletedFor"`
	SentOn         int64   `json:"sentOn" bson:"sentOn"`
	DeliveredOn    int64   `json:"deliveredOn" bson:"deliveredOn"`
}

func (message *Message) Isvalid() (bool, error) {
	if message.UserId == 0 {
		return false, system.InvalidPayloadData
	}

	if len(message.Body) == 0 {
		return false, system.InvalidPayloadData
	}

	if len(message.ConversationId) == 0 {
		return false, system.InvalidPayloadData
	}

	// repository.GetUserConversation()
	//validate conversationId

	return true, nil
}
