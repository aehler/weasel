package index

import (
	"weasel/app"
	"fmt"
	"weasel/lib/articles"
)

func Topics(c *app.Context) {

	lng := c.Get("lang").(string)

	t, err := articles.ListTopicsByLang(lng)

	if err != nil {
		c.RenderHTML("/errors/500.html", map[string]interface {} {
			"Error" : err.Error(),
		})

		c.Stop()

		return
	}

	c.RenderHTML(fmt.Sprintf("/%s/topics.html", lng), map[string]interface {} {
		"topics" : t,
	})

	return

}
