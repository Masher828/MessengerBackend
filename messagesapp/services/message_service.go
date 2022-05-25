package services

import (
	"github.com/Masher828/MessengerBackend/authapp/repository"
	"github.com/Masher828/MessengerBackend/common-packages/system"
	"github.com/Masher828/MessengerBackend/messagesapp/models"
	"github.com/sirupsen/logrus"
)

func SendMessage(message models.Message, log *logrus.Entry) error {

	message.SentOn = system.GetUTCTime().Unix()
	if 
	err := repository.InsertMessage(message, log)
	if err != nil {
		log.Errorln(err)
		return err
	}

	dataToupdate := map[string]interface{}{"recentMessage": message.Body}

	err = repository.UpdateConversationRecentMessage(dataToupdate, log)
	if err != nil {
		log.Errorln(err)
	}

	return err

}
