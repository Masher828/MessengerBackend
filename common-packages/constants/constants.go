package constants

var (

	//Conversation Type
	ConversationTypeGroup    = "GROUP"
	ConversationTypePersonal = "PERSONAL"

	//Mongo Database
	DatabaseSocialDB = "social_db"

	//Collection Name
	ConversationCollection     = "conversation"
	UserConversationCollection = "user_conversation"
	MessagesCollection         = "messages"

	//Conversation defaults
	DefaultOffset int64 = 0
	DefaultLimit  int64 = 10

	//Message Status
	MessageStatusSent      = "sent"
	MessageStatusFailed    = "failed"
	MessageStatusRead      = "read"
	MessageStatusDelivered = "delivered"
	MessageStatusDeleted   = "deleted"

	MongoNoDocumentErro = "mongo: no documents in result"

	//Middleware services constants
	AuthFailed  = "AuthFailed"
	UserContext = "UserContext"
)
