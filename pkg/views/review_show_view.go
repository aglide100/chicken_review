package views

import (
	"io"

	"github.com/aglide100/chicken_review_webserver/pkg/models"
)

type reviewShowView struct {
	htmlctx   BaseHTMLContext
	review    *models.Review
	CheckUser *models.User
	APIKeys   *models.APIKeys
}

func NewReviewShowView(htmlctx BaseHTMLContext, review *models.Review, APIKeys *models.APIKeys, CheckUser *models.User) View {
	return &reviewShowView{htmlctx: htmlctx, review: review, APIKeys: APIKeys, CheckUser: CheckUser}
}

func (view reviewShowView) ContentType() string {
	return "text/html"
}

func (view reviewShowView) Render(w io.Writer) error {
	//view.review.GoogleMapsApi = view.APIKeys.GoogleMaps
	Content := &models.Content{
		APIKeys: view.APIKeys,
		Review:  view.review,
	}
	return view.htmlctx.RenderUsing(w, "ui/reviews/show.gohtml", Content, view.CheckUser)
}
