package models

import (
	"regexp"
	"strings"
	"time"

	"github.com/Masher828/MessengerBackend/common-packages/system"
)

type UserModel struct {
	Id          int64     `json:"id"`
	FullName    string    `json:"name" column:"name,omitempty"`
	Email       string    `json:"email" column:"email,omitempty"`
	Password    string    `json:"password" column:"password,omitempty"`
	Contact     string    `json:"contact" column:"contact,omitempty"`
	CountryCode string    `json:"contry_code" column:"country_code,omitempty"`
	Country     string    `json:"country" column:"country,omitempty"`
	DateOfBirth time.Time `json:"date_of_birth" column:"date_of_birth,omitempty"`
	DateCreated time.Time `json:"date_created" column:"date_created,omitempty"`
	LastUpdated time.Time `json:"last_updated" column:"last_updated,omitempty"`
	LastLogin   time.Time `json:"last_login" column:"last_login,omitempty"`
}

type UserSearchDetails struct {
	Id       int64  `json:"id"`
	FullName string `json:"name"`
	Email    string `json:"email"`
}

type UserLoginModel struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserDetails struct {
	Id          int64  `json:"id"`
	FullName    string `json:"name" column:"name"`
	Email       string `json:"email" column:"email"`
	AccessToken string `json:"accessToken"`
}

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func (user *UserModel) IsValid() (bool, error) {
	if len(user.FullName) < 2 {
		return false, system.InvalidNameErr
	}

	if !emailRegex.MatchString(user.Email) {
		return false, system.InvalidEmailErr
	}

	if len(user.Password) < 8 || len(user.Password) > 20 {
		return false, system.InvalidPasswordFormatErr
	}

	if len(user.Contact) != 10 {
		return false, system.InvalidContactNumberErr
	}

	user.Email = strings.ToLower(user.Email)

	return true, nil
}

func (user *UserModel) CreateUserContext() *UserDetails {
	var userDetails UserDetails
	userDetails.Email = user.Email
	userDetails.FullName = user.FullName
	userDetails.Id = user.Id
	return &userDetails
}

func (user *UserLoginModel) IsValid() (bool, error) {
	if !emailRegex.MatchString(user.Email) {
		return false, system.InvalidEmailErr
	}

	if len(user.Password) < 8 || len(user.Password) > 20 {
		return false, system.InvalidPasswordFormatErr
	}

	user.Email = strings.ToLower(user.Email)

	return true, nil
}
