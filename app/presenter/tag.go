package presenter

import (
	"html/template"
	"net/http"

	"github.com/bmf-san/gobel-client-example/app/model"
)

// TagIndex is a data for index view.
type TagIndex struct {
	Tags       *model.Tags
	Pagination *model.Pagination
}

// ExecuteTagIndex responses a index view.
func (p *Presenter) ExecuteTagIndex(w http.ResponseWriter, t *TagIndex) error {
	tpl := template.Must(template.ParseFiles("view/layout/base.tpl", "view/tag/index.tpl", "view/partial/pagination.tpl", "view/partial/pagination.tpl"))
	if err := tpl.ExecuteTemplate(w, "base", t); err != nil {
		return err
	}
	return nil
}
