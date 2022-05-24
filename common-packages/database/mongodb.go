package database

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func GetMongoClient() (*mongo.Client, error) {
	// options := options.Client()

	// var timeOut = time.Second * 60

	// options.ServerSelectionTimeout = &timeOut

	host := viper.GetString("database.mongodb.host")
	port := viper.GetString("database.mongodb.port")
	prefix := viper.GetString("database.mongodb.prefix")
	username := viper.GetString("database.mongodb.username")
	password := viper.GetString("database.mongodb.password")

	uri := prefix + "://" + username + ":" + password + "@" + host

	if len(port) != 0 {
		uri += ":" + port
	}

	fmt.Println(uri)
	// client, err := mongo.Connect(context.TODO(), options.ApplyURI(uri))
	// if err != nil {
	// 	return nil, err
	// }

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://mongoose:bn00CkVZhgqkThWQ@cluster0.o8d1k.mongodb.net").
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)

	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		return nil, err
	}

	return client, err
}
