package views

import (
	"io"

	"github.com/aglide100/chicken_review_webserver/pkg/models"
)

type reviewListView struct {
	htmlctx   BaseHTMLContext
	reviews   []*models.Review
	CheckUser *models.User
}

func NewReviewListView(htmlctx BaseHTMLContext, reviews []*models.Review, CheckUser *models.User) View {
	return &reviewListView{htmlctx: htmlctx, reviews: reviews, CheckUser: CheckUser}
}

func (view reviewListView) ContentType() string {
	return "text/html"
}

func (view reviewListView) Render(w io.Writer) error {
	return view.htmlctx.RenderUsing(w, "ui/reviews/list.gohtml", view.reviews, view.CheckUser)
}
