package index

import (
	"weasel/app"
	"weasel/middleware/guest"
	"fmt"
	"weasel/middleware/auth"
	"time"
)

func Route(ap *app.App) {

	ap.Get("/", guest.GuestSettings, auth.GetAuthUser, Index)
	ap.GetPost("/login/", Login)
	ap.Get("/logout/", Logout)

	ap.Get("/language/:lang", guest.ResetLanguage, func(c *app.Context){

		c.Redirect("/")

	})

	ap.Get("/about/", guest.GuestSettings, func(c *app.Context){

		c.RenderHTML(fmt.Sprintf("/%s/about.html", c.Get("lang")), map[string]interface {} {	})

	})

	ap.Get("/topics/", guest.GuestSettings, Topics)

	ap.Get("/pong/", func(c *app.Context){

		time.Sleep(time.Millisecond * 50)

		c.RenderJSON(c.Params.ByName("v"))

	})

}