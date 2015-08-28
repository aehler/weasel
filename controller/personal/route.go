package personal

import (
	"weasel/app"
	"weasel/middleware/auth"
)

func Route(ap *app.App) {

	ap.Get("/settings/users/", auth.Check, PersonalEdit)
	ap.Get("/settings/organizations/", auth.Check, Organizations)
	ap.GetPost("/login/register/", RegisterUser)

}
