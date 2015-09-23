package references

import (
	"weasel/app"
	"weasel/middleware/auth"
)

func Route(ap *app.App) {

	ap.Get("/settings/references/", auth.Check, list)
	ap.Get("/settings/references/grid/", auth.Check, gridJSON)
	ap.GetPost("/settings/references/edit/:refId/", edit)
	ap.Get("/settings/references/items/:refId/", auth.Check, itemsList)
	ap.Get("/settings/references/items_grid/:refId/", auth.Check, itemsGridJSON)
	ap.GetPost("/settings/references/item_edit/:refId/:itemId/", auth.Check, editItem)
	ap.GetPost("/settings/references/item_add/:refId/:itemId/", auth.Check, addItem)

}
