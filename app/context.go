package app

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"encoding/json"
	"sync"
)

type Context struct {
	http.ResponseWriter
	mutex    *sync.Mutex
	values   map[string]interface{}
	Request  *http.Request
	Params   httprouter.Params
	handlers []Handler
	index    int
	stop     bool
}

func (c *Context) run() {

	for c.index < len(c.handlers) {

		c.handlers[c.index](c)

		c.index++

		if c.stop {

			return
		}
	}

}

func (c *Context) Stop() {

	c.stop = true
}

func (c *Context) IsPost() bool {

	if c.Request.Method == "POST" {

		return true

	}

	return false
}

func (c *Context) RenderJSON(value interface {}) error {

	return json.NewEncoder(c.ResponseWriter).Encode(value)

}

func (c *Context) RenderHTML(tmplName string, context interface {}) {

	if Templates[tmplName] == nil {

		c.RenderError("Template not found")
	}

	c.ResponseWriter.Header().Set("Content-Type", "text/html; charset=utf-8")

	jcontext, err := json.Marshal(context)

	if err != nil {
		jcontext = []byte(`{"error" : "Couldn't marshal JSON context"}`)
	}

	if err := Templates[tmplName].Execute(c.ResponseWriter, map[string]interface {}{
			"content" : string(jcontext),
			"current_user" : "{}",
	}); err != nil {

		c.RenderError(err.Error())
	}

}

func (c *Context) RenderError(e string) error {

	return json.NewEncoder(c.ResponseWriter).Encode(map[string]string{"error" : e})

}

func (c *Context) RenderLayout() {

	c.ResponseWriter.Header().Set("Content-Type", "text/html; charset=utf-8")

	http.ServeFile(c.ResponseWriter, c.Request, "/srv/src/weasel/static/main.html")

}
