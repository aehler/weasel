package index

import (
	"weasel/app"
	"weasel/app/crypto"
	"weasel/app/form"
	"weasel/lib/auth"
	"fmt"
)

func Index(c *app.Context) {

	c.RenderHTML("blank.html", map[string]interface {} {

	})

}

func Login(c *app.Context) {

	f := form.New("login", "login", "login_salt")

	data := auth.LoginForm{}

	if err := f.MapStruct(data); err != nil {

		c.RenderError(err.Error())
	}

	if c.IsPost() {

		u, err := auth.AuthUser("login", crypto.Encrypt("password", "llpass"))

		fmt.Println(u)
		fmt.Println(err)

	}

	c.RenderHTML("blank.html", map[string]interface {} {
		"form" : f,
	})

}
