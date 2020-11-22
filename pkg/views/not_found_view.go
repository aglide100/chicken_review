package views

import (
	"io"

	"github.com/aglide100/chicken_review_webserver/pkg/models"
)

//NotFoundView struct
type NotFoundView struct {
	htmlctx BaseHTMLContext
}

func (view NotFoundView) ContentType() string {
	return "text/html"
}

func NewNotFoundView(htmlctx BaseHTMLContext) View {
	return &NotFoundView{htmlctx: htmlctx}
}

// sended notFoundHTML
func (view NotFoundView) Render(w io.Writer) error {
	CheckUser := &models.User{
		Name:   "",
		UserID: "",
		Email:  "",
	}
	return view.htmlctx.RenderUsing(w, "ui/defaults/404.gohtml", nil, CheckUser)

}
