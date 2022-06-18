package routes

import (
	"github.com/Masher828/MessengerBackend/common-packages/system"
	"github.com/Masher828/MessengerBackend/messagesapp/controllers"
	"github.com/zenazn/goji"
)

func PrepareRoutes(application *system.Application) {

	//conversation
	goji.Post("/messages/conversation", application.Route(&controllers.Controller{}, "CreateConversation", false))
	goji.Get("/messages/conversation", application.Route(&controllers.Controller{}, "GetConversation", false))
	goji.Get("/messages/converstaion/:conversationId", application.Route(&controllers.Controller{}, "GetConversationById", false))

	//messages
	goji.Post("/messages/send", application.Route(&controllers.Controller{}, "SendMessage", false))
	goji.Get("/messages/conversation/:conversationId/messages", application.Route(&controllers.Controller{}, "GetMessagesForConversation", false))
	goji.Post("/messages/conversation/:conversationId/message/:messageId", application.Route(&controllers.Controller{}, "StartConversationMessage", false))
	goji.Delete("/messages/conversation/:conversationId/message/:messageId", application.Route(&controllers.Controller{}, "DeleteMessage", false))

}
