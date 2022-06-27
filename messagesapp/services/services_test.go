package services

import (
	"context"
	"fmt"
	"testing"

	"github.com/Masher828/MessengerBackend/common-packages/conf"
	"github.com/Masher828/MessengerBackend/common-packages/log"
	"github.com/Masher828/MessengerBackend/common-packages/system"
	messagesapprepository "github.com/Masher828/MessengerBackend/messagesapp/repository"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Entry

func init() {
	err := conf.LoadConfigFile()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = system.PrepareSocialContext()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	logger = log.GetDefaultLogger(0, "", "")
}

func Test_Check(t *testing.T) {
	fmt.Println(messagesapprepository.GetUserConversation(1, 0, 10, "", logger))
}

func Test_red(t *testing.T) {
	system.SocialContext.Redis.LPush(context.TODO(), "checking", "ddd12344")

	fmt.Println(system.SocialContext.Redis.LRange(context.TODO(), "checking", 0, 2).Result())
}

func Test_Conversation(t *testing.T) {
	fmt.Println(messagesapprepository.GetConversationByName(5, "pump", logger))
}
