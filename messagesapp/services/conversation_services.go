package services

import (
	authapprepository "github.com/Masher828/MessengerBackend/authapp/repository"
	"github.com/Masher828/MessengerBackend/common-packages/constants"
	"github.com/Masher828/MessengerBackend/common-packages/system"
	"github.com/Masher828/MessengerBackend/messagesapp/models"
	messagesapprepository "github.com/Masher828/MessengerBackend/messagesapp/repository"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func CreateConversation(conversation *models.Conversation, log *logrus.Entry) error {

	if _, err := conversation.IsValid(); err != nil {
		log.Errorln(err)
		return err
	}

	conversation.Id = uuid.New().String()

	now := system.GetUTCTime().Unix()
	conversation.UpdatedOn = now
	conversation.CreatedOn = now

	err := messagesapprepository.CreateConversation(conversation, log)
	if err != nil {
		log.Errorln(err)
		return err
	}

	var userConversations []interface{}

	userMap, count, err := authapprepository.GetUsersById(log, conversation.MemberIds)
	if err != nil {
		log.Errorln(err)
		return err
	} else if count != len(conversation.MemberIds) { // to check if all the users exists in the the database
		err = system.InvalidPayloadData
		log.Errorln(err)
		return err
	}

	for _, userId := range conversation.MemberIds {
		userConversation := models.UserConversation{
			Id:               string(uuid.New().String()),
			ConversationId:   conversation.Id,
			ConversationName: conversation.Name,
			UserId:           userId,
			IsArchived:       false,
			IsMuted:          false,
			UpdatedOn:        now,
			CreatedOn:        now,
		}
		if conversation.Type == constants.ConversationTypeGroup {
			userConversation.ConversationName = conversation.Name
		} else {
			userConversation.ConversationName = userMap[userId].FullName
		}

		userConversations = append(userConversations, &userConversation)
	}

	err = messagesapprepository.AddUserToConversation(userConversations, log)
	if err != nil {
		log.Errorln(err)
	}

	return err
}

func GetuserConversation(id int64, offset, limit int64, log *logrus.Entry) ([]models.ResponseUserConversation, error) {

	conversations, err := messagesapprepository.GetUserConversation(id, offset, limit, log)
	if err != nil {
		log.Errorln(err)
		return nil, err
	}

	return conversations, err
}

func GetConversationByName(id int64, pattern string, log *logrus.Entry) ([]models.UserConversation, error) {
	conversations, err := messagesapprepository.GetConversationByName(id, pattern, log)
	if err != nil {
		log.Errorln(err)
	}

	return conversations, err
}
