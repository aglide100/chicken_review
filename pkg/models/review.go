package models

type Review struct {
	Title             string
	Author            string
	StoreName         string
	Date              string
	PhoneNumber       string
	Comment           string
	Score             int
	ID                int64
	PictureURLs       []string
	DefaultPictureURL string
}

type User struct {
	UserID      string
	UserPWD     string
	Name        string
	Addr        string
	PhoneNumber string
	Language    string
}
