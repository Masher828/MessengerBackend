package system

import "errors"

var (

	// Authentication
	InvalidNameErr           = errors.New("please enter a valid name of minimum length of 2")
	InvalidEmailErr          = errors.New("please enter a valid email")
	InvalidPasswordFormatErr = errors.New("please enter a password of length (8,20)")
	InvalidContactNumberErr  = errors.New("please enter a valid contact number")
	InvalidCredentialsErr    = errors.New("invalid email id & password")

	InvalidPayloadData = errors.New("Invalid Data")

	EmailAlreadyExists = errors.New("Email already exists")

	// Conversation
	InvalidGroupType                  = errors.New("please select a valid group type")
	InvalidGroupName                  = errors.New("please enter a valid group name")
	InvalidPersonalChatName           = errors.New("personal chats cannot have name")
	InvalidPersonalConversationMember = errors.New("you cannot add more than 1 user in personal conversation")
	InvalidGroupMembersLimit          = errors.New("size of conversation can be between 2 - 200")

	UserNotPartOfConversation = errors.New("conversation does not exist")

	UnauthorizedErr            = errors.New("user is not authorized to perfrom the task")
	AccessTokenDoesNotExistErr = errors.New("access token does not exist in the redis")
	InternalServerErr          = errors.New("internal server error")
)
