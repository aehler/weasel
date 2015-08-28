package app

import (
	"html/template"
	"log"
)

var Templates = map[string]*template.Template{}

func InitTemplates(templates ...string) {

	for _, t := range templates {

		tmpl, err := template.ParseFiles(t)
		if err != nil {

			log.Fatal(err)

		}

		Templates[tmpl.Name()] = tmpl

	}

}
