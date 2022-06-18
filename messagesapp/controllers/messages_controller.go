package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Masher828/MessengerBackend/common-packages/constants"
	commonpackagesmodel "github.com/Masher828/MessengerBackend/common-packages/models"
	"github.com/Masher828/MessengerBackend/common-packages/system"
	"github.com/Masher828/MessengerBackend/messagesapp/models"
	"github.com/Masher828/MessengerBackend/messagesapp/services"
	"github.com/sirupsen/logrus"
	"github.com/zenazn/goji/web"
)

func (controller *Controller) SendMessage(c web.C, w http.ResponseWriter, r *http.Request, log *logrus.Entry) ([]byte, error) {

	var message models.MessageRequest

	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		log.Errorln(err)
		return nil, err
	}

	userContext := c.Env[constants.UserContext].(commonpackagesmodel.UserModelContext)

	err = services.SendMessage(message, userContext.Id, log)
	if err != nil {
		log.Errorln(err)
		return nil, err
	}

	response := map[string]string{"success": "ok"}
	return json.Marshal(response)

}

func (Controller *Controller) GetMessagesForConversation(c web.C, w http.ResponseWriter, r *http.Request, log *logrus.Entry) ([]byte, error) {

	conversationId := c.URLParams["conversationId"]

	if len(conversationId) == 0 {
		err := system.UserNotPartOfConversation
		log.Errorln(err)
		return nil, err
	}

	offset, limit := system.GetOffsetAndLimit(r.URL.Query()["offset"], r.URL.Query()["limit"], constants.DefaultOffset, constants.DefaultLimit, log)

	userContext := c.Env[constants.UserContext].(commonpackagesmodel.UserModelContext)

	messages, err := services.GetMessagesForConversation(conversationId, userContext.Id, offset, limit, log)
	if err != nil {
		log.Errorln(err)
		return nil, err
	}

	response := map[string]interface{}{"success": "ok", "data": messages}
	return json.Marshal(response)
}

func (Controller *Controller) StartConversationMessage(c web.C, w http.ResponseWriter, r *http.Request, log *logrus.Entry) ([]byte, error) {

	conversationId := c.URLParams["conversationId"]

	messageId := c.URLParams["messageId"]

	if len(conversationId) == 0 {
		err := system.InvalidConversationId
		log.Errorln(err)
		return nil, err
	}

	if len(messageId) == 0 {
		err := system.InvalidMessageId
		log.Errorln(err)
		return nil, err
	}

	userContext := c.Env[constants.UserContext].(commonpackagesmodel.UserModelContext)

	err := services.StarConversationMessage(conversationId, messageId, userContext.Id, log)
	if err != nil {
		log.Errorln(err)
		return nil, err
	}

	response := map[string]interface{}{"success": "ok"}
	return json.Marshal(response)
}
