package models

import "time"

type User struct {
	UserID      string
	UserPWD     string
	Name        string
	Addr        string
	PhoneNumber string
	Language    string
}

type ProviderUser struct {
	RawData           map[string]interface{}
	Provider          string
	Email             string
	Name              string
	FirstName         string
	LastEName         string
	NickName          string
	Description       string
	UserID            string
	AvatarURL         string
	Location          string
	AccessToken       string
	AccessTokenSecret string
	RefreshToken      string
	ExpiresAt         time.Time
	IDToken           string
}
