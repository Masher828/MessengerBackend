package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Masher828/MessengerBackend/messagesapp/models"
	"github.com/Masher828/MessengerBackend/messagesapp/services"
	"github.com/sirupsen/logrus"
	"github.com/zenazn/goji/web"
)

func (controller *Controller) SendMessage(c web.C, w http.ResponseWriter, r *http.Request, log *logrus.Entry) ([]byte, error) {

	var message models.Message

	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		log.Errorln(err)
		return nil, err
	}

	err = services.SendMessage(message, log)
	if err != nil {
		log.Errorln(err)
		return nil, err
	}

	response := map[string]string{"success": "ok"}
	return json.Marshal(response)

}
