package budget

import (
	"weasel/app"
	"weasel/middleware/auth"
)

func Route(ap *app.App) {

	ap.Get("/budget/fact/", auth.Check, fact)
	ap.Get("/budget/fact/grid/", auth.Check, factGrid)
	ap.GetPost("/budget/fact/edit/:rowId/", auth.Check, factForm)
	ap.Get("/budget/plan/", auth.Check, plan)
}


