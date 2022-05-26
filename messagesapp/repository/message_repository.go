package repository

import (
	"context"

	"github.com/Masher828/MessengerBackend/common-packages/constants"
	"github.com/Masher828/MessengerBackend/common-packages/system"
	"github.com/Masher828/MessengerBackend/messagesapp/models"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InsertMessage(message *models.Message, log *logrus.Entry) error {

	client := system.SocialContext.MongoClient
	db := client.Database(constants.DatabaseSocialDB).Collection(constants.MessagesCollection)

	_, err := db.InsertOne(context.TODO(), message)
	if err != nil {
		log.Errorln(err)
	}

	return err
}

func GetMessagesForConversation(conversationId string, userId, offset, limit int64, log *logrus.Entry) ([]*models.Message, error) {

	client := system.SocialContext.MongoClient
	db := client.Database(constants.DatabaseSocialDB).Collection(constants.MessagesCollection)

	opts := options.Find()

	opts.SetSort(bson.M{"senton": -1})
	opts.SetSkip(offset)
	opts.SetLimit(limit)

	where := bson.M{"conversationId": conversationId, "deletedFor": bson.M{"$nin": []int64{userId}}}

	result, err := db.Find(context.TODO(), where, opts)
	if err != nil {
		log.Errorln(err)
		return nil, err
	}

	var messages []*models.Message

	err = result.All(context.TODO(), &messages)
	if err != nil {
		log.Errorln(err)
	}
	return messages, err
}
