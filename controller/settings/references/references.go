package references

import (
	"weasel/app"
	"weasel/app/crypto"
	"weasel/app/grid"
	"weasel/app/form"
	"weasel/lib/references"
	"weasel/middleware/auth"
	"fmt"
)

func list(c *app.Context) {

	c.RenderHTML("/references/references.html", map[string]interface {} {

		"pageTitle" : "Справочники",
		"gridURL" : "/settings/references/grid/",

	})

}

func gridJSON(c *app.Context) {

	user := c.Get("user").(auth.User)

	refList, err := references.ReferenceList(user.OrganizationId)
	if err != nil {

		c.RenderJSON(map[string]interface {}{
			"Error" : err.Error(),
		})

		c.Stop()

		return
	}

	g := grid.New(refList)

	g.Column(
		&grid.Column{
			Name : "ID",
			Label: "#",
			Cell : grid.CellTypeInt,
			Order: 0,
		},
		&grid.Column{
			Name : "Name",
			Label: "Наименование справочника",
			Cell : grid.CellTypeString,
			Order: 1,
		},
		&grid.Column{
			Name : "Total",
			Label: "Элементов",
			Cell : grid.CellTypeInt,
			Order: 2,
		},
		&grid.Column{
			Name : "actions",
			Label: "..",
			Cell : grid.CellTypeActions,
			Order: 999,
		},
	)

	c.RenderJSON(g.Context())
}

func edit(c *app.Context) {

	c.RenderHTML("/blank.html", map[string]interface {} {

	})

}

func itemsList(c *app.Context) {

	c.RenderHTML("/references/references.html", map[string]interface {} {

		"pageTitle" : "Элементы справочника",
		"gridURL" : fmt.Sprintf("/settings/references/grid/%s", c.Params.ByName("refId")),

	})
}

func itemsGridJSON(c *app.Context) {

	var (
		user = c.Get("user").(auth.User)
		refid, _ = crypto.DecryptUrl(c.Params.ByName("refId"))
	)

	itemsList, err := references.ItemsList(user.OrganizationId, refid)
	if err != nil {

		c.RenderJSON(map[string]interface {}{
		"Error" : err.Error(),
	})

		c.Stop()

		return
	}

	g := grid.New(itemsList)

	g.Column(
		&grid.Column{
		Name : "ID",
		Label: "#",
		Cell : grid.CellTypeInt,
		Order: 0,
	},
		&grid.Column{
		Name : "Key",
		Label: "Код",
		Cell : grid.CellTypeInt,
		Order: 1,
	},
		&grid.Column{
		Name : "Label",
		Label: "Значение",
		Cell : grid.CellTypeString,
		Order: 2,
	},
		&grid.Column{
		Name : "Ord",
		Label: "Порядок в списке",
		Cell : grid.CellTypeInt,
		Order: 3,
	},
		&grid.Column{
		Name : "actions",
		Label: "..",
		Cell : grid.CellTypeActions,
		Order: 999,
	},
	)

	c.RenderJSON(g.Context())

}

func editItems(c *app.Context) {

	f := form.New("Элементы справочника", "register", "login_salt")

	c.RenderJSON(f)

}
