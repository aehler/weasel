package app

import (
	"github.com/flosch/pongo2"
	"fmt"
	"io/ioutil"
	"log"
)

var Templates = map[string]*pongo2.Template{}

var dir = "/srv/src/weasel/templates/pages"

func InitTemplates() {

	parseDir("")

}

func parseDir(dirname string) {

	cdir := fmt.Sprintf("%s%s", dir, dirname)

	fi, err := ioutil.ReadDir(fmt.Sprintf("%s%s", dir, dirname))
	if err != nil {
		log.Fatal("Cannot access template dir", cdir)
	}

	for _, file := range fi {

		if file.IsDir() {

			parseDir(fmt.Sprintf("%s/%s", dirname, file.Name()))

		} else {

			tmpl := pongo2.Must(pongo2.FromFile(fmt.Sprintf("%s/%s", cdir, file.Name())))

			Templates[fmt.Sprintf("%s/%s", dirname, file.Name())] = tmpl

			fmt.Println("Added template", fmt.Sprintf("%s/%s", dirname, file.Name()))
		}
	}
}
