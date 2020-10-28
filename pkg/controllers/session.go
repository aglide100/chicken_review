package controllers

import (
	"fmt"
	"net/http"

	"github.com/aglide100/chicken_review_webserver/pkg/models"
	"github.com/gorilla/sessions"
)

type SessionController struct {
	store *sessions.CookieStore
}

func NewSessionController(store *sessions.CookieStore) *SessionController {
	return &SessionController{store: store}
}

func (hdl *SessionController) SaveSession(resp http.ResponseWriter, req *http.Request, user *models.User, providerUser *models.ProviderUser, UserType string) error {
	session, err := hdl.store.Get(req, "session-name")
	if err != nil {
		return fmt.Errorf("Can't get session", err)
	}
	if UserType == "Goth" {
		session.Values["UserID"] = providerUser.UserID
		session.Values["Name"] = providerUser.Name
		session.Values["Email"] = providerUser.Email
	} else {
		session.Values["UserID"] = user.UserID
		session.Values["Name"] = user.Name
		session.Values["Email"] = user.Email
	}
	session.Save(req, resp)
	return nil
}

func (hdl *SessionController) GetSession(resp http.ResponseWriter, req *http.Request) string {
	session, err := hdl.store.Get(req, "session-name")
	if err != nil {
		return "err"
	}
	if (session.Values["Name"] != nil) || (session.Values["UserID"] != nil) || (session.Values["Email"] != nil) {
		return "true"
	} else {
		return "nil"
	}
}

func (hdl *SessionController) GetUserDataInSession(resp http.ResponseWriter, req *http.Request) *models.User {
	session, err := hdl.store.Get(req, "session-name")
	if err != nil {
		fmt.Errorf("Can't get session!")
	}
	user := &models.User{
		UserID: session.Values["UserID"].(string),
		Name:   session.Values["Name"].(string),
		Email:  session.Values["Email"].(string),
	}

	return user
}

func (hdl *SessionController) RemoveSession(resp http.ResponseWriter, req *http.Request) error {
	session, err := hdl.store.Get(req, "session-name")
	if err != nil {
		return fmt.Errorf("Can't get session", err)
	}
	delete(session.Values, "UserID")
	delete(session.Values, "Name")
	delete(session.Values, "Email")
	session.Save(req, resp)
	return nil
}
