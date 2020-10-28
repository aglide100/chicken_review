package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aglide100/chicken_review_webserver/pkg/db"
)

type User struct {
	UserID string
	Name   string
	Email  string
}

type AjaxController struct {
	db          *db.Database
	sessionCtrl *SessionController
}

func NewAjaxController(db *db.Database, sessionCtrl *SessionController) *AjaxController {
	return &AjaxController{db: db, sessionCtrl: sessionCtrl}
}

func (hdl *AjaxController) AjaxHandler(resp http.ResponseWriter, req *http.Request) {
	//parse request to struct
	log.Printf("receive ajax handler")
	//var d User
	//err := json.NewDecoder(req.Body).Decode(&d)
	//if err != nil {
	//	http.Error(resp, err.Error(), http.StatusInternalServerError)
	//}

	// create json response from struct
	//var s = `{"Name":"test","Email":"test","UserID":"test"}`
	var s = `nope`
	a, err := json.Marshal(s)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
	}
	resp.Write(a)
}
