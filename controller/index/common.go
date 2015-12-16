package index

import (
	"weasel/app"
	"weasel/app/crypto"
	"weasel/app/registry"
	"weasel/app/session"
	"weasel/lib/auth"
	"fmt"
)

func Index(c *app.Context) {

	user := c.Get("user")

	if user != nil {

		app.Redirect("/dashboard/", c, 302)

		return

	}

	c.RenderHTML("/public.html", map[string]interface {} {

	})

}

func Logout(c *app.Context) {

	var sd string

	if err := session.Get(c.Request, &sd, &session.Config{Keys : registry.Registry.SessionKeys}); err != nil {

		fmt.Println(err)

		app.Redirect("/login/", c, 302)

		return
	}

	registry.Registry.Session.Kill(sd)

	app.Redirect("/login/", c, 302)

	return
}

func Login(c *app.Context) {

	if c.IsPost() {

		c.Request.ParseForm()

		c.Request.PostFormValue("email")

		u, err := auth.AuthUser(c.Request.PostFormValue("email"), crypto.Encrypt(c.Request.PostFormValue("password"), ""))
		if err != nil {

			if err.Error() == `sql: no rows in result set` {

				c.RenderJSON(map[string]interface {} { "loginError" : "Логин или пароль не верны"})

				return

			} else {

				c.RenderError(err.Error())
			}

			fmt.Println("couldn't login", err)

			c.Stop()

			return

		}

		ssid := crypto.GenSessionId(u.UserID, u.UserLastName)

		if err := registry.Registry.Session.Add(ssid, u); err != nil {

			c.RenderError(err.Error())
		}

		if err := session.Set(c.ResponseWriter, ssid, &session.Config{Keys : registry.Registry.SessionKeys}); err != nil {

			fmt.Println("couldn't set cookie")

			c.RenderError(err.Error())

		}

		app.Redirect("/dashboard/", c, 302)

//		c.RenderJSON(map[string]interface {} { "redirect" : "/dashboard/"})
//
//		c.Stop()

		return

	}

	c.RenderHTML("/login.html", map[string]interface {} {
	})

}
