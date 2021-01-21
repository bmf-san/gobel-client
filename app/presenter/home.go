package presenter

import (
	"html/template"
	"net/http"
)

// ExecuteHomeIndex responses a index template.
func (p *Presenter) ExecuteHomeIndex(w http.ResponseWriter) error {
	tpl := template.Must(template.ParseFS(tpls, "template/layout/base.tpl", "template/home/index.tpl"))
	if err := tpl.ExecuteTemplate(w, "base", ""); err != nil {
		return err
	}
	return nil
}
