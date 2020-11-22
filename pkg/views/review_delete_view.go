package views

import (
	"io"

	"github.com/aglide100/chicken_review_webserver/pkg/models"
)

type reviewDeleteView struct {
	htmlctx BaseHTMLContext
	review  *models.Review
}

func NewReviewDeleteView(htmlctx BaseHTMLContext, review *models.Review) View {
	return &reviewDeleteView{htmlctx: htmlctx, review: review}
}

func (view reviewDeleteView) ContentType() string {
	return "text/html"
}

func (view reviewDeleteView) Render(w io.Writer) error {
	CheckUser := &models.User{
		Name:   "",
		Email:  "",
		UserID: "",
	}
	return view.htmlctx.RenderUsing(w, "ui/reviews/delete.gohtml", view.review, CheckUser)
}
