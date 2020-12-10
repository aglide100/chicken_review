package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aglide100/chicken_review_webserver/pkg/db"
	"github.com/aglide100/chicken_review_webserver/pkg/models"
	"github.com/aglide100/chicken_review_webserver/pkg/views"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
)

type LoginController struct {
	db          *db.Database
	sessionCtrl *SessionController
	Providers   *goth.Providers
}

func NewLoginController(db *db.Database, sessionCtrl *SessionController) *LoginController {

	return &LoginController{db: db, sessionCtrl: sessionCtrl}
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

	if hdl.sessionCtrl.GetSession(resp, req) == "true" {
		log.Printf("find Session data")
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

	User := &models.User{
		UserID:  UserID,
		UserPWD: UserPWD,
	}
	//gob.Register(&models.User{})

	hdl.sessionCtrl.SaveSession(resp, req, User, nil, "Local")
	log.Printf("save session, id: %v pwd: %v", UserID, UserPWD)

	view := views.NewReviewLoginView(views.DefaultBaseHTMLContext)
	resp.Header().Set("Content-Type", view.ContentType())
	err := view.Render(resp)
	if err != nil {
		log.Printf("faild to render : %v", err)
	}
}

func (hdl *LoginController) LogOut(resp http.ResponseWriter, req *http.Request) {
	log.Printf("[login_func]: receive request to LogOut")

	hdl.sessionCtrl.RemoveSession(resp, req)
	resp.Header().Set("Location", "/reviews")
	resp.WriteHeader(http.StatusTemporaryRedirect)
}

/* Check Provider(Goauth) User */
func (hdl *LoginController) AuthGoth(resp http.ResponseWriter, req *http.Request) {
	log.Printf("[login_func]: receive request to goauth")

	usr, err := gothic.CompleteUserAuth(resp, req)
	if err == nil {
		// parse user data
		user := GothUserChangeToPuser(&usr)

		err = hdl.db.RegisterNewProviderUser(user)
		if err != nil {
			log.Printf("Can't register Goauth User: %v", err)
			http.Redirect(resp, req, "/reviews", 301)
		}

	} else {
		gothic.BeginAuthHandler(resp, req)
	}

}

func GothUserChangeToPuser(user *goth.User) *models.ProviderUser {
	PUser := &models.ProviderUser{
		RawData:           user.RawData,
		Provider:          user.Provider,
		Email:             user.Email,
		Name:              user.Name,
		FirstName:         user.FirstName,
		LastName:          user.LastName,
		NickName:          user.NickName,
		Description:       user.Description,
		UserID:            user.UserID,
		AvatarURL:         user.AvatarURL,
		Location:          user.Location,
		AccessToken:       user.AccessToken,
		AccessTokenSecret: user.AccessTokenSecret,
		RefreshToken:      user.RefreshToken,
		ExpiresAt:         user.ExpiresAt,
		IDToken:           user.IDToken,
	}

	return PUser
}

func (hdl *LoginController) GothLogOut(resp http.ResponseWriter, req *http.Request) {
	log.Printf("[login_func]: receive request to logout gothUser")
	gothic.Logout(resp, req)

	hdl.sessionCtrl.RemoveSession(resp, req)
	resp.Header().Set("Location", "/reviews")
	resp.WriteHeader(http.StatusTemporaryRedirect)
}

func (hdl *LoginController) GothCallBack(resp http.ResponseWriter, req *http.Request) {
	log.Printf("[login_func]: receive request to callback goth")

	user, err := gothic.CompleteUserAuth(resp, req)
	if err != nil {
		fmt.Fprintln(resp, err)
	}

	log.Printf("User :%v, %v, %v", user.Name, user.NickName, user.Email)
	//resp.Header().Set("Location", "/")
	pUser := GothUserChangeToPuser(&user)

	err, ok := hdl.db.CheckProviderUser(pUser)
	if err != nil {
		// err check
	}
	if !ok {
		hdl.db.RegisterNewProviderUser(pUser)
	}
	hdl.sessionCtrl.RemoveSession(resp, req)
	hdl.sessionCtrl.SaveSession(resp, req, nil, pUser, "Goth")
	http.Redirect(resp, req, "/reviews", 301)
	//resp.Header().Set("Location", "/reviews")
}
