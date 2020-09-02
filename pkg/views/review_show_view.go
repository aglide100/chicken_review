package views

import (
	"io"

	"github.com/aglide100/chicken_review_webserver/pkg/models"
)

type reviewShowView struct {
	htmlctx BaseHTMLContext
	review  *models.Review
	APIKeys *models.APIKeys
}

func NewReviewShowView(htmlctx BaseHTMLContext, review *models.Review, APIKeys *models.APIKeys) View {
	return &reviewShowView{htmlctx: htmlctx, review: review, APIKeys: APIKeys}
}

func (view reviewShowView) ContentType() string {
	return "text/html"
}

func (view reviewShowView) Render(w io.Writer) error {
	//view.review.GoogleMapsApi = view.APIKeys.GoogleMaps
	return view.htmlctx.RenderUsing(w, "ui/reviews/show.gohtml", view.review)
}
