package repository

import (
	"context"

	"github.com/Masher828/MessengerBackend/common-packages/constants"
	"github.com/Masher828/MessengerBackend/common-packages/system"
	"github.com/Masher828/MessengerBackend/messagesapp/models"
	"github.com/sirupsen/logrus"
)

func InsertMessage(message models.Message, log *logrus.Entry) error {

	client := system.SocialContext.MongoClient
	db := client.Database(constants.DatabaseSocialDB).Collection(constants.MessagesCollection)

	_, err := db.InsertOne(context.TODO(), message)
	if err != nil {
		log.Errorln(err)
	}

	return err
}
