package references

import (
	"weasel/app"

)

func list(c *app.Context) {

	c.RenderHTML("/references/references.html", map[string]interface {} {

		"pageTitle" : "Справочники",

	})

}

func edit(c *app.Context) {

	c.RenderHTML("/blank.html", map[string]interface {} {

	})

}
