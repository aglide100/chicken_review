package views

import (
	"io"
	"path/filepath"
)

type reviewGetScriptView struct {
	htmlctx BaseHTMLContext
	path    string
}

func NewReviewGetScriptView(htmlctx BaseHTMLContext, path string) View {
	return &reviewGetScriptView{htmlctx: htmlctx, path: path}
}

func (view reviewGetScriptView) ContentType() string {
	var contentType string

	view.path = view.path[9:]
	ext := filepath.Ext(view.path)

	switch ext {
	case ".js":
		contentType = "text/css"
	case ".cs":
		contentType = "text/javascript"
	}

	return contentType
}

func (view reviewGetScriptView) Render(w io.Writer) error {
	return view.htmlctx.RenderImage(w, view.path)
}
