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

	//log.Printf("view path :%v ////////// %v", view.path, view.path[16:])
	// /reviews/ui/css/common.css ////////// common.css
	// 16자가 최소한의 경로
	ext := filepath.Ext(view.path[16:])
	switch ext {
	case ".js":
		contentType = "text/javascript"
	case ".css":
		contentType = "text/css"
	}

	return contentType
}

func (view reviewScriptView) Render(w io.Writer) error {
	//log.Printf("view.path : %v", view.path)
	return view.htmlctx.RenderFile(w, view.path)
}
