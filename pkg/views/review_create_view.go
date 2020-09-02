package views

import (
	"io"

	"github.com/aglide100/chicken_review_webserver/pkg/models"
)

type reviewCreateView struct {
	htmlctx BaseHTMLContext
	APIKeys *models.APIKeys
}

func NewReviewCreateView(htmlctx BaseHTMLContext, APIKeys *models.APIKeys) View {
	return &reviewCreateView{htmlctx: htmlctx, APIKeys: APIKeys}
}

func (view reviewCreateView) ContentType() string {
	return "text/html"
}

func (view reviewCreateView) Render(w io.Writer) error {
	return view.htmlctx.RenderUsing(w, "ui/reviews/create.gohtml", view.APIKeys)

}
