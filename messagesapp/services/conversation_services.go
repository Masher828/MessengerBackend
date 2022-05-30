package services

import (
	"reflect"

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

	exists, err := authapprepository.CheckIfUsersExist(log, conversation.MemberIds)
	if err != nil {
		log.Errorln(err)
		return err
	} else if !exists { // to check if all the users exists in the the database
		err = system.InvalidPayloadData
		log.Errorln(err)
		return err
	}

	for _, userId := range conversation.MemberIds {
		userConversation := models.UserConversation{
			Id:             string(uuid.New().String()),
			ConversationId: conversation.Id,
			UserId:         userId,
			IsArchived:     false,
			IsMuted:        false,
			UpdatedOn:      now,
			CreatedOn:      now,
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

	userIdsMap := map[int64][]int{}

	for i := range conversations {
		if conversations[i].Conversation[0].Type == constants.ConversationTypePersonal {
			var friendId int64
			if conversations[i].Conversation[0].MemberIds[0] == id {
				friendId = conversations[i].Conversation[0].MemberIds[1]
			} else {
				friendId = conversations[i].Conversation[0].MemberIds[0]
			}

			if _, ok := userIdsMap[friendId]; ok {
				userIdsMap[friendId] = append(userIdsMap[friendId], i)
			} else {
				userIdsMap[friendId] = []int{i}
			}
		}
	}

	userids := []int64{}

	keys := reflect.ValueOf(userIdsMap).MapKeys()
	for i := range keys {
		userids = append(userids, int64(keys[i].Int()))
	}

	userMap, _, err := authapprepository.GetUsersById(log, userids)
	if err != nil {
		log.Errorln(err)
		return nil, err
	}

	for userId := range userIdsMap {
		for _, i := range userIdsMap[userId] {
			conversations[i].Conversation[0].Name = userMap[userId].FullName
		}
	}

	return conversations, err
}
