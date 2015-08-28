package main

import (
	"weasel/controller/personal"
	"weasel/controller/index"
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
	config = flag.String("config", "./config.yml", "config file")
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

}
