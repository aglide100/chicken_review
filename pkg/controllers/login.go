package controllers

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/aglide100/chicken_review_webserver/pkg/db"
	"github.com/aglide100/chicken_review_webserver/pkg/models"
	"github.com/aglide100/chicken_review_webserver/pkg/views"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
)

type LoginController struct {
	db        *db.Database
	store     *sessions.CookieStore
	Providers *goth.Providers
}

func NewLoginController(db *db.Database, store *sessions.CookieStore) *LoginController {

	return &LoginController{db: db, store: store}
}

func findProvider(resp http.ResponseWriter, req *http.Request) string {
	var matches []string

	var authProviderPattern = regexp.MustCompile("^/auth/[A-Za-z]")

	matches = authProviderPattern.FindStringSubmatch(req.URL.Path)

	result := strings.Join(matches, "")

	switch result {
	case "naver":
		return result
	case "github":
		return result
	default:
		return "No match providers"
	}
}

func (hdl *LoginController) Register_Page(resp http.ResponseWriter, req *http.Request) {
	log.Printf("[login_func]: receive request to register_page")

	view := views.NewRegisterView(views.DefaultBaseHTMLContext)
	resp.Header().Set("Content-Type", view.ContentType())
	err := view.Render(resp)
	if err != nil {
		log.Printf("failed to render: %v", err)
	}
}

func (hdl *LoginController) Register(resp http.ResponseWriter, req *http.Request) {
	log.Printf("[login_func]: receive request to register view")

	NewUser := &models.User{
		UserID:      req.PostFormValue("UserID"),
		UserPWD:     req.PostFormValue("UserPWD"),
		Name:        req.PostFormValue("Name"),
		Addr:        req.PostFormValue("Addr"),
		PhoneNumber: req.PostFormValue("PhoneNumber"),
		Language:    req.PostFormValue("Language"),
	}

	err := hdl.db.RegisterNewUser(NewUser)
	if err != nil {
		log.Printf("failed to register : %v", err)
	}

	view := views.NewRegisterView(views.DefaultBaseHTMLContext)
	resp.Header().Set("Content-Type", view.ContentType())
	err = view.Render(resp)
	if err != nil {
		log.Printf("failed render : %v", err)
	}

}

func (hdl *LoginController) LoginCheck(resp http.ResponseWriter, req *http.Request) {
	log.Printf("[login_func]: receive request to LoginCheck")

	if SessionControl(hdl, resp, req, nil, "check") == "nil" {

	}

	view := views.NewReviewLoginView(views.DefaultBaseHTMLContext)
	resp.Header().Set("Contetn-Type", view.ContentType())
	err := view.Render(resp)
	if err != nil {
		log.Printf("faild to render : %v", err)
	}

}

func SessionControl(hdl *LoginController, resp http.ResponseWriter, req *http.Request, user *models.User, set string) string {
	if set == "save" {
		session, _ := hdl.store.Get(req, "session-name")
		session.Values["user"] = &models.User{
			UserID:  user.UserID,
			UserPWD: user.UserPWD}
		session.Save(req, resp)
		return "saved"
	} else if set == "check" {
		session, _ := hdl.store.Get(req, "session-name")
		log.Println(session.Values["user"])
		if session.Values["user"] == nil {
			log.Println("[Login]: there are no session!")

			return "nil"
		}
		return "checked"
	}
	return "default"
}

/* Check Local User */
func (hdl *LoginController) LogIn(resp http.ResponseWriter, req *http.Request) {
	log.Printf("[login_func]: receive request to login view")

	// 로그인 채크 로직 ++
	UserID := req.PostFormValue("UserID")
	UserPWD := req.PostFormValue("UserPWD")

	User := &models.User{
		UserID:  UserID,
		UserPWD: UserPWD,
	}
	//gob.Register(&models.User{})

	SessionControl(hdl, resp, req, User, "save")
	log.Printf("save session, id: %v pwd: %v", UserID, UserPWD)

	// Goauth 와 로컬 유저 체크 하는 로직 넣기

	view := views.NewReviewLoginView(views.DefaultBaseHTMLContext)
	resp.Header().Set("Content-Type", view.ContentType())
	err := view.Render(resp)
	if err != nil {
		log.Printf("faild to render : %v", err)
	}
}

func (hdl *LoginController) LogOut(resp http.ResponseWriter, req *http.Request) {
	log.Printf("[login_func]: receive request to LogOut")

	// 세션 지우기
}

/* Check Provider(Goauth) User */
func (hdl *LoginController) AuthGoth(resp http.ResponseWriter, req *http.Request) {
	log.Printf("[login_func]: receive request to goauth")

	usr, err := gothic.CompleteUserAuth(resp, req)
	if err == nil {
		// parse user data

		user := &models.ProviderUser{
			RawData:           usr.RawData,
			Provider:          usr.Provider,
			Email:             usr.Email,
			Name:              usr.Name,
			FirstName:         usr.FirstName,
			LastEName:         usr.LastName,
			NickName:          usr.NickName,
			Description:       usr.Description,
			UserID:            usr.UserID,
			AvatarURL:         usr.AvatarURL,
			Location:          usr.Location,
			AccessToken:       usr.AccessToken,
			AccessTokenSecret: usr.AccessTokenSecret,
			RefreshToken:      usr.RefreshToken,
			ExpiresAt:         usr.ExpiresAt,
			IDToken:           usr.IDToken,
		}

		err = hdl.db.RegisterNewGoauthUser(user)
		if err != nil {
			log.Printf("Can't register Goauth User: %v", err)
			User := &models.User{
				UserID: user.UserID,
			}
			SessionControl(hdl, resp, req, User, "save")

			http.Redirect(resp, req, "/reviews", 301)
			//resp.Header().Set("Location", "/")
			//resp.WriteHeader(http.StatusTemporaryRedirect)
		} else {

			http.Redirect(resp, req, "/reviews", 301)
			//resp.Header().Set("Location", "/")
		}
	} else {
		gothic.BeginAuthHandler(resp, req)

		//log.Printf("Can't find Goauth User: %v", err)

		//resp.Header().Set("Location", "/")
		//resp.WriteHeader(http.StatusTemporaryRedirect)
	}

}

/*
func (hdl *LoginController) GothLogIn(resp http.ResponseWriter, req *http.Request) {

}
*/

func (hdl *LoginController) GothLogOut(resp http.ResponseWriter, req *http.Request) {
	log.Printf("[login_func]: receive request to logout gothUser")
	gothic.Logout(resp, req)

	resp.Header().Set("Location", "/reviews")
	resp.WriteHeader(http.StatusTemporaryRedirect)
}

func (hdl *LoginController) GothCallBack(resp http.ResponseWriter, req *http.Request) {
	log.Printf("[login_func]: receive request to callback goth")

	user, err := gothic.CompleteUserAuth(resp, req)
	if err != nil {
		fmt.Fprintln(resp, err)
	}

	resp.Header().Set("Location", "/")
	log.Printf("User :%v", user)
}
