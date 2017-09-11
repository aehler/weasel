package guest

import (
	"fmt"
	"weasel/app/crypto"
	"weasel/app"
	"weasel/app/session"
	"weasel/app/registry"
)

type GuestSession struct {
	ID string
	Lang string
}

func GuestSettings(c *app.Context) {

	gc := GuestSession{}

	if err := session.Get(c.Request, &gc, &session.Config{Keys : registry.Registry.SessionKeys}); err == nil {

		c.Set("ssid", gc.ID)
		c.Set("lang", gc.Lang)

		return

	}

	ssid := crypto.GenSessionId(0, "guest")

	gc.ID = ssid
	gc.Lang = "en"

	if err := session.Set(c.ResponseWriter, gc, &session.Config{Keys : registry.Registry.SessionKeys}); err != nil {

		fmt.Println("couldn't set cookie")

		c.RenderError(err.Error())

		c.Stop()

		return

	}

	c.Set("ssid", gc.ID)
	c.Set("lang", gc.Lang)

}

func ResetLanguage(c *app.Context) {

	gc := GuestSession{}

	lang := c.Params.ByName("lang")

	if err := session.Get(c.Request, &gc, &session.Config{Keys : registry.Registry.SessionKeys}); err == nil {

		gc.Lang = lang

		if err := session.Set(c.ResponseWriter, gc, &session.Config{Keys : registry.Registry.SessionKeys}); err != nil {

			fmt.Println("couldn't set cookie")

			c.RenderError(err.Error())

			c.Stop()

			return

		}

		c.Set("lang", gc.Lang)

	}

}