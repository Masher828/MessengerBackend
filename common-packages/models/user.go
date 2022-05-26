package models

type UserModelContext struct {
	Id          int64  `json:"id"`
	FullName    string `json:"name" column:"name"`
	Email       string `json:"email" column:"email"`
	AccessToken string `json:"accessToken"`
}
