package personal

import (
	"weasel/app"
	"weasel/middleware/auth"
)

func Route(ap *app.App) {

	ap.Get("/dashboard/", auth.Check, Dashboard)
	ap.Get("/personal/settings/", auth.Check, PersonalEdit)
	ap.Get("/settings/organizations/", auth.Check, Organizations)
	ap.GetPost("/login/register/", RegisterUser)

}
