package controllers

import (
	"encoding/gob"
	"log"
	"net/http"

	"github.com/aglide100/chicken_review_webserver/pkg/db"
	"github.com/aglide100/chicken_review_webserver/pkg/models"
	"github.com/aglide100/chicken_review_webserver/pkg/views"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth/gothic"
)

type LoginController struct {
	db    *db.Database
	store *sessions.CookieStore
}

func NewLoginController(db *db.Database, store *sessions.CookieStore) *LoginController {
	return &LoginController{db: db, store: store}
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
	session, _ := hdl.store.Get(req, "session-name")
	log.Println(session.Values["user"])
	if session.Values["user"] == nil {
		log.Println("[Login]: there are no session!")
		//http.Redirect(resp, req, "/login", 301)
	}

	view := views.NewReviewLoginView(views.DefaultBaseHTMLContext)
	resp.Header().Set("Contetn-Type", view.ContentType())
	err := view.Render(resp)
	if err != nil {
		log.Printf("faild to render : %v", err)
	}

}

/* Check Local User */
func (hdl *LoginController) LogIn(resp http.ResponseWriter, req *http.Request) {
	log.Printf("[login_func]: receive request to login view")

	// 로그인 채크 로직 ++
	UserID := req.PostFormValue("UserID")
	UserPWD := req.PostFormValue("UserPWD")

	gob.Register(&models.User{})

	session, _ := hdl.store.Get(req, "session-name")
	session.Values["user"] = &models.User{UserID}
	session.Save(req, resp)
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
func (hdl *LoginController) Register_goauth(resp http.ResponseWriter, req *http.Request) {
	log.Printf("[login_func]: receive request to goauth")

	usr, err := gothic.CompleteUserAuth(resp, req)
	if err != nil {
		log.Printf("Can't find Goauth User: %v", err)
	}

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
	}
}
