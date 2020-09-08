package views

import (
	"io"
	"path/filepath"
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

	ext := filepath.Ext(view.path)
	switch ext {
	case ".js":
		contentType = "text/javascript"
	case ".css":
		contentType = "text/css"
	}

	return contentType
}

func (view reviewAssetsView) Render(w io.Writer) error {
	return view.htmlctx.RenderImage(w, view.path)
}
