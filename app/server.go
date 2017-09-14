package app

import (
	"weasel/app/registry"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"runtime/debug"
	"log"
	"github.com/flosch/pongo2"
	"fmt"
)

type App struct {
	Router *httprouter.Router
	handlers []Handler
}

type Handler404 struct {}

func New(config string) *App {

	a := App{
		Router: httprouter.New(),
	}

	pathes := registry.ReadPathConf(config)

	InitTemplates(pathes.Templates)

	fmt.Printf("Serve static on /%s/*filepath\n", pathes.HTTPStatic)

	a.Router.ServeFiles(fmt.Sprintf("/%s/*filepath", pathes.HTTPStatic), http.Dir(pathes.Static))

	a.Router.NotFound = Handler404{}.ServeHTTP

	a.Router.PanicHandler = func(rw http.ResponseWriter, _ *http.Request, err interface{}) {

		rw.Header().Set("Content-Type", "text/html")

		rw.WriteHeader(http.StatusInternalServerError)

		Templates["/errors/500.html"].ExecuteWriter(pongo2.Context{"Error" : "DON'T PANIC"}, rw)

		log.Printf("PANIC: %s\n", debug.Stack())
	}

	registry.Init(config)

	a.Get("/metrics/", metricsHandler)

	return &a
}

func (e Handler404) ServeHTTP(rw http.ResponseWriter, _ *http.Request) {

	Templates["/errors/404.html"].ExecuteWriter(pongo2.Context{}, rw)

}

func (a *App) Get(route string, handlers ...Handler) {

	handler := handler(append(a.handlers, handlers...))

	a.Router.GET(route, handler)
	a.Router.HEAD(route, handler)

}

func (a *App) Post(pattern string, handlers ...Handler) {

	a.Router.POST(pattern, handler(append(a.handlers, handlers...)))
}

func (a *App) GetPost(pattern string, handlers ...Handler) {

	handler := handler(append(a.handlers, handlers...))

	a.Router.GET(pattern, handler)
	a.Router.POST(pattern, handler)
}

func (a *App) Handler(h Handler) {

	a.handlers = append(a.handlers, h)
}

func Redirect(url string, c *Context, code int) {

	c.Stop()

	http.Redirect(c.ResponseWriter, c.Request, url, code)

}
