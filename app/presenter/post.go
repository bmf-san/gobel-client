package presenter

import (
	"html/template"
	"net/http"

	"github.com/bmf-san/gobel-client-example/app/model"
)

// PostIndex is a data for index template.
type PostIndex struct {
	Posts      *model.Posts
	Pagination *model.Pagination
}

// PostIndexByCategory is a data for index template by category.
type PostIndexByCategory struct {
	CategoryName string
	Posts        *model.Posts
	Pagination   *model.Pagination
}

// PostIndexByTag is a data for index template by tag.
type PostIndexByTag struct {
	TagName    string
	Posts      *model.Posts
	Pagination *model.Pagination
}

// PostShow is a data for show template.
type PostShow struct {
	Post *model.Post
}

// ExecutePostIndex responses a index template.
func (pt *Presenter) ExecutePostIndex(w http.ResponseWriter, p *PostIndex) error {
	tpl := template.Must(template.ParseFS(tpls, "template/layout/base.tpl", "template/post/index.tpl", "template/partial/pagination.tpl"))
	if err := tpl.ExecuteTemplate(w, "base", p); err != nil {
		return err
	}
	return nil
}

// ExecutePostIndexByCategory responses a index template by category.
func (pt *Presenter) ExecutePostIndexByCategory(w http.ResponseWriter, p *PostIndexByCategory) error {
	tpl := template.Must(template.ParseFS(tpls, "template/layout/base.tpl", "template/post/category.tpl", "template/partial/pagination.tpl"))
	if err := tpl.ExecuteTemplate(w, "base", p); err != nil {
		return err
	}
	return nil
}

// ExecutePostIndexByTag responses a index template by tag.
func (pt *Presenter) ExecutePostIndexByTag(w http.ResponseWriter, p *PostIndexByTag) error {
	tpl := template.Must(template.ParseFS(tpls, "template/layout/base.tpl", "template/post/tag.tpl", "template/partial/pagination.tpl"))
	if err := tpl.ExecuteTemplate(w, "base", p); err != nil {
		return err
	}
	return nil
}

// ExecutePostShow responses a show template by tag.
func (pt *Presenter) ExecutePostShow(w http.ResponseWriter, p *PostShow) error {
	fm := template.FuncMap{
		"unescape": pt.unescape,
	}
	tpl := template.Must(template.New("base").Funcs(fm).ParseFS(tpls, "template/layout/base.tpl", "template/post/show.tpl"))
	if err := tpl.ExecuteTemplate(w, "base", p); err != nil {
		return err
	}
	return nil
}
