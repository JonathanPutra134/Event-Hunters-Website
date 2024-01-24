package dto

import "time"

type UserRegistrationRequest struct {
	FullName    string
	Email       string
	Password    string
	PhoneNumber string
	Address     string
	Latitude    string
	Longitude   string
}

type Session struct {
	ID         string
	Data       string
	Expiration time.Time
}