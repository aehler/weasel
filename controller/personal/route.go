package personal

import (
	"weasel/app"
	"weasel/middleware/auth"
)

func Route(ap *app.App) {

	ap.Get("/dashboard/", auth.Check, Dashboard)
	ap.Get("/personal/settings/", auth.Check, PersonalEdit)
	ap.Get("/settings/organizations/", auth.Check, listOrgs)
	ap.Get("/settings/organizations/grid/", auth.Check, organizationsGrid)
	ap.GetPost("/settings/organizations/edit/:orgId/", auth.Check, organizationsEdit)
	ap.GetPost("/login/register/", RegisterUser)

}
