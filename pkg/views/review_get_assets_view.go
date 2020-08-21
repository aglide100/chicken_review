package views

import (
	"io"
)

type reviewAssetsView struct {
	htmlctx BaseHTMLContext
	path    string
}

func NewReviewGetAssetsView(htmlctx BaseHTMLContext, path string) View {
	return &reviewAssetsView{htmlctx: htmlctx, path: path}
}

func (view reviewAssetsView) ContentType() string {
	var contentType string

	//view.path = view.path[16:]

	return contentType
}

func (view reviewAssetsView) Render(w io.Writer) error {
	return view.htmlctx.RenderImage(w, view.path)
}
