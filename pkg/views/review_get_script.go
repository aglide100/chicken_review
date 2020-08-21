package views

import (
	"io"
	"path/filepath"
)

type reviewScriptView struct {
	htmlctx BaseHTMLContext
	path    string
}

func NewReviewGetScriptView(htmlctx BaseHTMLContext, path string) View {
	return &reviewScriptView{htmlctx: htmlctx, path: path}
}

func (view reviewScriptView) ContentType() string {
	var contentType string

	view.path = view.path[16:]
	ext := filepath.Ext(view.path)
	switch ext {
	case ".js":
		contentType = "text/javascript"
	case ".css":
		contentType = "text/css"
	}

	return contentType
}

func (view reviewScriptView) Render(w io.Writer) error {
	return view.htmlctx.RenderImage(w, view.path)
}
