package templates

import (
	"html/template"
	"log"
	"path/filepath"
)

var Tmpl *template.Template

func init() {
	var err error
	Tmpl, err = template.ParseGlob(filepath.Join("templates", "*.html"))
	if err != nil {
		log.Fatal(err)
	}
}
