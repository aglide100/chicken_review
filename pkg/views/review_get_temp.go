package views

import (
	"io"
)

type reviewTempView struct {
	htmlctx BaseHTMLContext
}

func NewReviewGetTempView(htmlctx BaseHTMLContext) View {
	return &reviewTempView{htmlctx: htmlctx}
}

func (view reviewTempView) ContentType() string {
	return "text/html"
}

func (view reviewTempView) Render(w io.Writer) error {
	return view.htmlctx.RenderFile(w, "ui/reviews/temp.gohtml")
}
