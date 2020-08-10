package views

import (
	"io"
	"log"
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

	log.Printf("view.path before: %v", view.path)
	view.path = view.path[16:]
	ext := filepath.Ext(view.path)
	log.Printf("view.path after: %v", view.path)
	switch ext {
	case ".js":
		log.Printf("It is javascript, ", ext)
		contentType = "text/javascript"
	case ".css":
		log.Printf("It is css,", ext)
		contentType = "text/css"
	}

	return contentType
}

func (view reviewGetScriptView) Render(w io.Writer) error {
	return view.htmlctx.RenderImage(w, view.path)
}
