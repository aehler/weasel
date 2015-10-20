package main

import (
	"weasel/controller/budget"
	"weasel/controller/settings"
	"weasel/controller/personal"
	"weasel/controller/index"
	"weasel/controller/storage"
	"weasel/app"
	"log"
	"flag"
	"net/http"
	"fmt"
)

type Router interface {

	Route(a *app.App)
}

var (
	port   = flag.Uint("port", 80, "the port to listen on")
	config = flag.String("config", "conf.d", "config directory")
)

func main() {

	flag.Parse()

	a := app.New(*config)

	collect(a)

	fmt.Println("Starting server on port", *port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), a.Router))

}

func collect(a *app.App) {

	personal.Route(a)
	index.Route(a)
	settings.Route(a)
	budget.Route(a)
	storage.Route(a)
}
