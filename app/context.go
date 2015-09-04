package app

import (
	"github.com/julienschmidt/httprouter"
	"github.com/flosch/pongo2"
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

func (c *Context) Set(key string, value interface{}) {

	c.mutex.Lock()

	c.values[key] = value

	c.mutex.Unlock()
}

func (c *Context) Get(key string) interface{} {

	defer c.mutex.Unlock()

	c.mutex.Lock()

	if v, found := c.values[key]; found {

		return v
	}

	return nil
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

	c.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
	c.ResponseWriter.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")

	return json.NewEncoder(c.ResponseWriter).Encode(value)

}

func (c *Context) RenderHTML(tmplName string, context map[string]interface {}) {

	if Templates[tmplName] == nil {

		c.RenderError("Template not found")
	}

	c.ResponseWriter.Header().Set("Content-Type", "text/html; charset=utf-8")
	c.ResponseWriter.Header().Set("Cache-Control", "no-cache,no-store,must-revalidate")
	c.ResponseWriter.Header().Set("Pragma", "no-cache")
	c.ResponseWriter.Header().Set("Connection", "keep-alive")
	c.ResponseWriter.Header().Set("Expires", "0")

	context["currentUser"] = c.Get("user")

	if err := Templates[tmplName].ExecuteWriter(pongo2.Context(context), c.ResponseWriter); err != nil {

		c.RenderError(err.Error())
	}

}

func (c *Context) RenderError(e string) error {

	return json.NewEncoder(c.ResponseWriter).Encode(map[string]string{"error" : e})

}

func (c *Context) RenderLayout() {

	c.ResponseWriter.Header().Set("Content-Type", "text/html; charset=utf-8")
	c.ResponseWriter.Header().Set("Cache-Control", "no-cache,no-store,must-revalidate")
	c.ResponseWriter.Header().Set("Pragma", "no-cache")
	c.ResponseWriter.Header().Set("Connection", "keep-alive")
	c.ResponseWriter.Header().Set("Expires", "0")

	http.ServeFile(c.ResponseWriter, c.Request, "/srv/src/weasel/static/main.html")

}
