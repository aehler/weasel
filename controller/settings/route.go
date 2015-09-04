package settings

import (
	"weasel/app"
	"weasel/controller/settings/references"
)

func Route(ap *app.App) {

	references.Route(ap)

}


