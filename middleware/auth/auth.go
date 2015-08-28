package auth

import (
	"weasel/app"
	"weasel/app/session"
	"weasel/lib/auth"
)

type Auth struct {
	User *auth.User
	SSID string
}

func Check(c *app.Context) {

	var sd interface {}

	if err := session.Get(c.Request, &sd, &session.Config{}); err != nil {

		app.Redirect("/login/", c, 301)

		return
	}

}
