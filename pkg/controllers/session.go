package controllers

import (
	"fmt"
	"log"
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
		if providerUser.UserID == "" {
			session.Values["UserID"] = "<empty>"
		} else {
			session.Values["UserID"] = providerUser.UserID
		}

		if providerUser.Name == "" {
			session.Values["Name"] = "<empty>"
		} else {
			session.Values["Name"] = providerUser.Name
		}

		if providerUser.Email == "" {
			session.Values["Email"] = "<empty>"
		} else {
			session.Values["Email"] = providerUser.Email
		}
	} else {
		if user.UserID == "" {
			session.Values["UserID"] = "<empty>"
		} else {
			session.Values["UserID"] = user.UserID
		}

		if user.Name == "" {
			session.Values["Name"] = "<empty>"
		} else {
			session.Values["Name"] = user.Name
		}

		if user.Email == "" {
			session.Values["Email"] = "<empty>"
		} else {
			session.Values["Email"] = user.Email
		}
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

func (hdl *SessionController) GetUserDataInSession(req *http.Request) *models.User {
	log.Printf("Get User data in session!")
	session, err := hdl.store.Get(req, "session-name")
	if err != nil {
		fmt.Errorf("Can't get session!")
	}

	var (
		UserID string
		Name   string
		Email  string
	)

	if session.Values["UserID"] == nil {
		UserID = "<empty>"
	} else {
		UserID = session.Values["UserID"].(string)
	}

	if session.Values["Name"] == nil {
		Name = "<empty>"
	} else {
		Name = session.Values["Name"].(string)
	}

	if session.Values["Email"] == nil {
		Email = "<empty>"
	} else {
		Email = session.Values["Email"].(string)
	}

	user := &models.User{
		UserID: UserID,
		Name:   Name,
		Email:  Email,
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
