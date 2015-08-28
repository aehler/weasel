package index

import (
	"weasel/app"
)

func Route(ap *app.App) {

	ap.Get("/", Index)
	ap.GetPost("/login/", Login)

}


