package database

import (
	"context"
	"fmt"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func GetMongoClient() (*mongo.Client, error) {

	host := viper.GetString("database.mongodb.host")
	port := viper.GetString("database.mongodb.port")
	prefix := viper.GetString("database.mongodb.prefix")
	username := viper.GetString("database.mongodb.username")
	password := viper.GetString("database.mongodb.password")

	uri := prefix + "://" + username + ":" + password + "@" + host

	if len(port) != 0 {
		uri += ":" + port
	}

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)

	clientOptions := options.Client().
		ApplyURI(uri).
		SetServerAPIOptions(serverAPIOptions)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println(err)
	}

	rp := readpref.Primary()
	err = client.Ping(context.TODO(), rp)
	if err != nil {
		fmt.Println(err)
	}

	return client, err
}
