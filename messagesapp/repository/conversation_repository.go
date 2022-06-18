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

func CreateConversation(conversation *models.Conversation, log *logrus.Entry) error {

	client := system.SocialContext.MongoClient
	db := client.Database(constants.DatabaseSocialDB).Collection(constants.ConversationCollection)

	_, err := db.InsertOne(context.TODO(), conversation)
	if err != nil {
		log.Errorln(err)
	}
	return err
}

func AddUserToConversation(userConversations []interface{}, log *logrus.Entry) error {

	client := system.SocialContext.MongoClient
	db := client.Database(constants.DatabaseSocialDB).Collection(constants.UserConversationCollection)

	_, err := db.InsertMany(context.TODO(), userConversations)
	if err != nil {
		log.Errorln(err)
	}
	return err
}

func GetUserConversation(userId int64, offset, limit int64, log *logrus.Entry) ([]models.ResponseUserConversation, error) {

	client := system.SocialContext.MongoClient
	db := client.Database(constants.DatabaseSocialDB).Collection(constants.UserConversationCollection)

	var conversations []models.ResponseUserConversation
	pipeline := []bson.M{}

	conditionToMatchUserId := bson.M{"$match": bson.M{"userId": userId}}
	pipeline = append(pipeline, conditionToMatchUserId)

	conditionToGetConversationDetails := bson.M{"$lookup": bson.M{"from": constants.ConversationCollection, "localField": "conversationId", "foreignField": "_id", "as": "conversation"}}
	pipeline = append(pipeline, conditionToGetConversationDetails)

	sortByLastUpdated := bson.M{"$sort": bson.M{"conversation.lastUpdated": -1}}
	pipeline = append(pipeline, sortByLastUpdated)

	if offset != 0 {
		addOffset := bson.M{"$skip": offset}
		pipeline = append(pipeline, addOffset)
	}

	if limit != 0 {
		addLimit := bson.M{"$limit": limit}
		pipeline = append(pipeline, addLimit)
	}

	opts := options.Aggregate()
	cursor, err := db.Aggregate(context.TODO(), pipeline, opts)
	if err != nil {
		log.Errorln(err)
		return nil, err
	}

	err = cursor.All(context.TODO(), &conversations)
	if err != nil {
		log.Errorln(err)
		return nil, err
	}

	return conversations, nil
}

func IsUserPartOfConversation(userId int64, conversationId string, log *logrus.Entry) bool {
	client := system.SocialContext.MongoClient
	db := client.Database(constants.DatabaseSocialDB).Collection(constants.UserConversationCollection)

	where := bson.M{"conversationId": conversationId, "userId": userId}
	result := db.FindOne(context.TODO(), where)

	return result.Err() == nil
}

func UpdateConversation(conversationId string, dataToBeUpdated map[string]interface{}, log *logrus.Entry) error {
	client := system.SocialContext.MongoClient
	db := client.Database(constants.DatabaseSocialDB).Collection(constants.ConversationCollection)

	where := bson.M{"_id": conversationId}

	data := bson.M{"$set": dataToBeUpdated}

	_, err := db.UpdateOne(context.TODO(), where, data)
	if err != nil {
		log.Errorln(err)
	}

	return err
}

func AddMessageToStarredMessages(conversationId, messageId string, userId int64, log *logrus.Entry) error {

	client := system.SocialContext.MongoClient
	db := client.Database(constants.DatabaseSocialDB).Collection(constants.UserConversationCollection)

	update := bson.M{"$push": bson.M{"starredMessages": messageId}}

	where := bson.M{"conversationId": conversationId, "userId": userId}

	_, err := db.UpdateOne(context.TODO(), where, update)
	if err != nil {
		log.Errorln(err)
	}

	return err

}
