package index

import (
	"weasel/app"
	"weasel/middleware/guest"
	"fmt"
)

func Route(ap *app.App) {

	ap.Get("/", guest.GuestSettings, Index)
	ap.GetPost("/login/", Login)
	ap.Get("/logout/", Logout)

	ap.Get("/language/:lang", guest.ResetLanguage, func(c *app.Context){

		c.Redirect("/")

	})

	ap.Get("/about/", guest.GuestSettings, func(c *app.Context){

		c.RenderHTML(fmt.Sprintf("/%s/about.html", c.Get("lang")), map[string]interface {} {	})

	})

}