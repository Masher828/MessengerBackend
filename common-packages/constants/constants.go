package constants

import "time"

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

	//Auth
	IncorrectPasswordCountEmail       = "IncorrectPasswordCountEmail:%s"
	IncorrectPasswordCountEmailExpiry = 30 * 24 * time.Hour // 30 days
	MaxIncorrectPasswordAllowed       = 5

	//Access Token Constants
	UserToAccessTokenList = "UserToAccessTokenKeys:%s"
	AccessTokenToUser     = "AccessTokenTouser:%s"
	AccessTokenExpiry     = 3 * 24 * time.Hour // 3 days
)
