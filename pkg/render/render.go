package render

import (
	"bytes"
	"html/template"
	"log"
	"morningEat/pkg/config"
	"morningEat/pkg/models"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}
var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a

}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	tmpl = filepath.Base(tmpl)

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
		return
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println("Error writing template to browser", err)

	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myChache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myChache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myChache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myChache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myChache, err
			}
		}
		myChache[name] = ts

	}
	return myChache, nil
}
