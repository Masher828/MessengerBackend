package models

type Message struct {
	UserId      int64  `json:"userId" bson:"userId"`
	Type        string `json:"type" bson:"type"`
	Body        string `json:"body" bson:"body"`
	SentOn      int64  `json:"sentOn" bson:"sentOn"`
	DeliveredOn int64  `json:"deliveredOn" bson:"deliveredOn"`
}

func (message *Message) Isvalid() (bool, error) {
	if message.UserId == 0 {
		return false, 
	}
}
