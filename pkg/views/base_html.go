package views

import (
	"fmt"
	"io"
	"io/ioutil"
	"text/template"

	"github.com/aglide100/chicken_review_webserver/pkg/models"
	"github.com/aglide100/chicken_review_webserver/ui"
)

var DefaultBaseHTMLContext = BaseHTMLContext{
	GlobPattern: "ui/*.gohtml",
	HTML: func(bodyContent interface{}, CheckUser *models.User) ui.HTML {
		return ui.HTML{
			Head: ui.Head{
				FavIcoURL:   "",
				Title:       "Chicken Review",
				Author:      "",
				Description: "We review chicken restaurants",
			},
			Body: ui.Body{Content: bodyContent, CheckUser: CheckUser},
		}
	},
}

type BaseHTMLContext struct {
	GlobPattern string
	HTML        func(bodyContent interface{}, CheckUser *models.User) ui.HTML
}

func (htmlctx *BaseHTMLContext) RenderFile(w io.Writer, path string) error {
	// /reviews/ -> 9번째
	content, err := ioutil.ReadFile(path[9:])
	if err != nil {
		return nil
	}

	w.Write(content)

	return nil
}

func (htmlctx *BaseHTMLContext) RenderImage(w io.Writer, path string) error {

	// remove to /reviews/0
	content, err := ioutil.ReadFile(path[9:])
	if err != nil {
		return nil
	}

	w.Write(content)

	return nil
}

func (htmlctx *BaseHTMLContext) RenderUsing(w io.Writer, contentPattern string, bodyContent interface{}, CheckUser *models.User) error {
	baseT, err := template.ParseGlob(htmlctx.GlobPattern)
	if err != nil {
		return fmt.Errorf("parsing base html: %v", err)
	}
	contentT, err := template.Must(baseT.Clone()).ParseGlob(contentPattern)
	if err != nil {
		return fmt.Errorf("parsing reviews html: %v", err)
	}

	html := htmlctx.HTML(bodyContent, CheckUser)
	if err := contentT.ExecuteTemplate(w, "html", html); err != nil {
		return fmt.Errorf("executing template: %v", err)
	}
	return nil
}
