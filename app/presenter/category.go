package presenter

import (
	"html/template"
	"net/http"

	"github.com/bmf-san/gobel-client-example/app/model"
)

// CategoryIndex is a data for index template.
type CategoryIndex struct {
	Categories *model.Categories
	Pagination *model.Pagination
}

// ExecuteCategoryIndex responses a index template.
func (p *Presenter) ExecuteCategoryIndex(w http.ResponseWriter, c *CategoryIndex) error {
	tpl := template.Must(template.ParseFS(tpls, "template/layout/base.tpl", "template/category/index.tpl", "template/partial/pagination.tpl", "template/partial/pagination.tpl"))
	if err := tpl.ExecuteTemplate(w, "base", c); err != nil {
		return err
	}
	return nil
}
