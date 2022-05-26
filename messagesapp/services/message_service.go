package services

import (
	"github.com/Masher828/MessengerBackend/common-packages/constants"
	"github.com/Masher828/MessengerBackend/common-packages/system"
	"github.com/Masher828/MessengerBackend/messagesapp/models"
	"github.com/Masher828/MessengerBackend/messagesapp/repository"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func SendMessage(messageRequest models.MessageRequest, userId int64, log *logrus.Entry) error {

	now := system.GetUTCTime().Unix()

	message := messageRequest.GetMessage()

	message.Id = uuid.New().String()
	message.SentOn = now
	message.Status = constants.MessageStatusSent
	message.UserId = userId

	if _, err := messageRequest.Isvalid(); err != nil {
		log.Errorln(err)
		return err
	}

	if !repository.IsUserPartOfConversation(message.UserId, message.ConversationId, log) {
		err := system.UserNotPartOfConversation
		log.Errorln(err)
		return err
	}

	err := repository.InsertMessage(message, log)
	if err != nil {
		log.Errorln(err)
		return err
	}

	dataToupdate := map[string]interface{}{"recentMessage": message.Body, "updatedOn": now}

	err = repository.UpdateConversation(message.ConversationId, dataToupdate, log)
	if err != nil {
		log.Errorln(err)
	}

	return err

}

func GetMessagesForConversation(conversationId string, userId int64, offset, limit int64, log *logrus.Entry) ([]*models.Message, error) {

	if !repository.IsUserPartOfConversation(userId, conversationId, log) {
		err := system.UserNotPartOfConversation
		log.Errorln(err)
		return nil, err
	}

	messages, err := repository.GetMessagesForConversation(conversationId, userId, offset, limit, log)
	if err != nil {
		log.Errorln(err)
		return nil, err
	}

	return messages, nil

}
