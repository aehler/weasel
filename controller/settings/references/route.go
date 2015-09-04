package references

import (
	"weasel/app"
	"weasel/middleware/auth"
)

func Route(ap *app.App) {

	ap.Get("/settings/references/", auth.Check, list)
	ap.GetPost("/settings/references/edit/:refId/", edit)

}
