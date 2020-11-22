package ui

import "github.com/aglide100/chicken_review_webserver/pkg/models"

type HTML struct {
	Head Head
	Body Body
}

type Head struct {
	Title       string
	FavIcoURL   string
	Author      string
	Description string
	OG          *OGData
	Twitter     *Twitter
	ContentHead interface{}
}

type Body struct {
	Content   interface{}
	CheckUser *models.User
}

type Lang struct {
	Data interface{}
}

type OGData struct {
	Title        string
	Type         string
	ImageURL     string
	CanonicalURL string

	SiteName    string
	Description string
}

type Twitter struct {
	ImageURL    string
	Title       string
	Description string
}
