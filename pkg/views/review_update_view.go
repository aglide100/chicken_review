package views

import (
	"io"

	"github.com/aglide100/chicken_review_webserver/pkg/models"
)

type reviewUpdateView struct {
	htmlctx   BaseHTMLContext
	review    *models.Review
	CheckUser *models.User
}

func NewReviewUpdateView(htmlctx BaseHTMLContext, review *models.Review, CheckUser *models.User) View {
	return &reviewUpdateView{htmlctx: htmlctx, review: review, CheckUser: CheckUser}
}

func (view reviewUpdateView) ContentType() string {
	return "text/html"
}

func (view reviewUpdateView) Render(w io.Writer) error {

	return view.htmlctx.RenderUsing(w, "ui/reviews/update.gohtml", view.review, view.CheckUser)
}
