package views

import (
	"io"

	"github.com/aglide100/chicken_review_webserver/pkg/models"
)

type reviewLoginView struct {
	htmlctx BaseHTMLContext
}

func NewReviewLoginView(htmlctx BaseHTMLContext) View {
	return &reviewLoginView{htmlctx: htmlctx}
}

func (view reviewLoginView) ContentType() string {
	return "text/html"
}

func (view reviewLoginView) Render(w io.Writer) error {
	CheckUser := &models.User{
		UserID: "",
		Name:   "",
		Email:  "",
	}

	return view.htmlctx.RenderUsing(w, "ui/reviews/login.gohtml", nil, CheckUser)
	// if (check status) -> view.htmlctx.RenderUsing(w, "ui/reviews/login_status.gohtml", nil)
}
