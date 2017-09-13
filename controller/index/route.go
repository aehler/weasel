package index

import (
	"weasel/app"
	"weasel/middleware/guest"
	"fmt"
	"weasel/middleware/auth"
)

func Route(ap *app.App) {

	ap.Get("/", auth.GetAuthUser, Index)
	ap.GetPost("/login/", Login)
	ap.Get("/logout/", Logout)

	ap.Get("/language/:lang", guest.ResetLanguage, func(c *app.Context){

		c.Redirect("/")

	})

	ap.Get("/about/", auth.GetAuthUser, func(c *app.Context){

		c.RenderHTML(fmt.Sprintf("/%s/about.html", c.Get("lang")), map[string]interface {} {	})

	})

	ap.Get("/topics/", auth.GetAuthUser, Topics)

}