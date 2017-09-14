package personal

import (
	"weasel/app"
	"weasel/app/form"
	"weasel/lib/auth"
)

func PersonalEdit(c *app.Context) {

	user := c.Get("user").(auth.User)
	ssid := c.Get("user").(string)

	f := form.New("Мои данные", "register", ssid)

	data := auth.RegisterForm{
		UserLastName : user.UserLastName,
		UserFirstName : user.UserFirstName,
		UserMiddleName : user.UserMiddleName,
		Email : user.Email,
	}

	f.Skip("Password", "Password2", "Login")

	if err := f.MapStruct(data); err != nil {

		c.RenderError(err.Error())

		c.Stop()

		return
	}

	c.RenderJSON(f)

}

func Dashboard(c *app.Context) {

	c.RenderHTML("/personal/dashboard.html", map[string]interface {} {
	})

}

func RegisterUser(c *app.Context) {

	f := form.New("Регистрация", "register", "login_salt")

	f.Skip("Email")

	data := auth.RegisterForm{}

	if err := f.MapStruct(data); err != nil {

		c.RenderError(err.Error())

		c.Stop()

		return
	}

	if c.IsPost() {

		if err := f.ParseForm(&data, c.Request); err != nil {

			c.RenderError(err.Error())

			c.Stop()

			return

		}

		_, err := auth.AddUser(data)
		if err != nil {

			c.RenderError(err.Error())

			c.Stop()

			return
		}

		c.RenderJSON(map[string]interface {} {
			"message" : "success",
			"label" : `Вы успешно зарегистрировались в системе <a href="/login/">Вход в систему</a>`,
		})

		c.Stop()

		return

	}

	c.RenderJSON(f)

}
