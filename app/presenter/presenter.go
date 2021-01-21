package presenter

import (
	"embed"
	"html/template"
)

// Presenter represents the singular of presenter.
type Presenter struct{}

// NewPresenter creates a Presenter.
func NewPresenter() *Presenter {
	return &Presenter{}
}

//go:embed template
var tpls embed.FS

func (p *Presenter) unescape(text string) template.HTML {
	return template.HTML(text)
}
