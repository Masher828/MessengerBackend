package models

import "time"

type UserModel struct {
	Id          int64
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Contact     string    `json:"contact"`
	CountryCode string    `json:"contry_code"`
	Country     string    `json:"country"`
	DateOfBirth time.Time `json:"date_of_birth"`
	DateCreated time.Time `json:"date_created"`
	LastUpdated time.Time `json:"last_updated"`
}