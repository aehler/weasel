package storage

import (
	"weasel/app"
	"weasel/middleware/auth"
)

func Route(ap *app.App) {

	ap.Get("/storage/get/:fileId/", auth.Check, file)
	ap.Post("/storage/add/", auth.Check, fileForm)

}


