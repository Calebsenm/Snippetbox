package main

import (
	"html/template"
	"path/filepath"
	"time"
    "io/fs"
	"github.com/calebsenm/snippetbox/internal/models"
    "github.com/calebsenm/snippetbox/ui"
)

type templateData struct {
	CurrentYear     int
	Snippet         *models.Snippet
	Snippets        []*models.Snippet
	Form            any
	Flash           string
	IsAuthenticated bool 
    CSRFToken       string
}

func humanDate(t time.Time) string {
    if t.IsZero() {
        return ""
    }
    return t.UTC().Format("02 Jan 2006 at 15:04")
 }


// Initialize a template.FuncMap object and store it in a global variable. This is
// essentially a string-keyed map which acts as a lookup between the names of our
// custom template functions and the functions themselves.
var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache() (map[string]*template.Template, error) {
	
    cache := map[string]*template.Template{}
    pages, err := fs.Glob(ui.Files, "html/pages/*.tmpl")

	if err != nil {
		return nil, err
	}

	for _, page := range pages {
        name := filepath.Base(page)

        patterns := []string{
            "html/base.tmpl",
            "html/partials/*.tmpl",
            page, 
        }

        ts, err := template.New(name).Funcs(functions).ParseFS(ui.Files, patterns...)
        if err != nil {
            return nil, err
        }

        cache[name] = ts

    }
	return cache, nil
}
