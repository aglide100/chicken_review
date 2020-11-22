package views

import (
	"io"

	"github.com/aglide100/chicken_review_webserver/pkg/models"
)

type RegisterView struct {
	htmlctx BaseHTMLContext
}

func (view RegisterView) ContentType() string {
	return "text/html"
}

func NewRegisterView(htmlctx BaseHTMLContext) View {
	return &RegisterView{htmlctx: htmlctx}
}

func (view RegisterView) Render(w io.Writer) error {
	CheckUser := &models.User{
		Name:   "",
		Email:  "",
		UserID: "",
	}
	return view.htmlctx.RenderUsing(w, "ui/reviews/register.gohtml", nil, CheckUser)
}
