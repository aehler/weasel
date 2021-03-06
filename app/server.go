package app

import (
	"weasel/app/registry"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type App struct {
	Router *httprouter.Router
	handlers []Handler
}

func New(config string) *App {

	a := App{
		Router: httprouter.New(),
	}

	pathes := registry.ReadPathConf(config)

	InitTemplates(pathes.Templates)

	a.Router.ServeFiles("/static/*filepath", http.Dir(pathes.Static))

	//a.Router.NotFound = http.FileServer(http.Dir("static/404.html")).ServeHTTP

	registry.Init(config)

	//a.Handler( func (c *Context) { c.RenderLayout() })

	return &a
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
