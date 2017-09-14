package auth

import (
	"weasel/app"
	"weasel/app/session"
	"weasel/app/registry"
	"weasel/lib/auth"
	"encoding/json"
	"fmt"
	"weasel/middleware/guest"
)

func GetAuthUser(c *app.Context) {

	var sd string

	u := auth.Auth{}

	if err := session.Get(c.Request, &sd, &session.Config{Keys : registry.Registry.SessionKeys, Name:"auth"}); err != nil {

		fmt.Println(err)

		return
	}

	v, err := registry.Registry.Session.Get(sd)

	if err != nil {

		fmt.Println(err)

		return
	}

	if err := json.Unmarshal(v, &u); err != nil {

		guest.GuestSettings(c)

		fmt.Println(err)

		return
	}

	u.SSID = sd

}

func Check(c *app.Context) {

	var sd string

	if err := session.Get(c.Request, &sd, &session.Config{Keys : registry.Registry.SessionKeys, Name:"auth"}); err != nil {

		fmt.Println(err)

		app.Redirect("/login/", c, 302)

		return
	}

	v, err := registry.Registry.Session.Get(sd)

	if err != nil {

		fmt.Println(err)

		app.Redirect("/login/", c, 302)

		return
	}

	u := auth.Auth{}

	if err := json.Unmarshal(v, &u); err != nil {

		fmt.Println(err)

		app.Redirect("/login/", c, 302)

		return
	}

	u.SSID = sd

	c.Set("user", u)
}