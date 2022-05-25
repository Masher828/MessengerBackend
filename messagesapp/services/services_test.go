package services

import (
	"fmt"
	"testing"

	"github.com/Masher828/MessengerBackend/common-packages/conf"
	"github.com/Masher828/MessengerBackend/common-packages/log"
	"github.com/Masher828/MessengerBackend/common-packages/system"
	"github.com/Masher828/MessengerBackend/messagesapp/repository"
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
	fmt.Println(repository.IsUserPartOfConversation(1, "d416731a-e49c-453a-8a91-0f6370ed818", logger))
}
