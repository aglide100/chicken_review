package views

import "io"

type RegisterView struct {
	htmlctx BaseHTMLContext
}

func (view RegisterView) ContentType() string {
	return "text/html"
}

func NewRegisterView(htmlctx BaseHTMLContext) View {
	return &RegisterView{htmlctx: htmlctx}
}

func (view RegisterView) Render(w io.Writer) error {
	return view.htmlctx.RenderUsing(w, "ui/reviews/register.gohtml", nil)
}