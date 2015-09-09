package references

import (
	"weasel/app"
	"weasel/app/grid"
	"weasel/app/form"
	"weasel/lib/references"
	"weasel/middleware/auth"
)

func list(c *app.Context) {

	c.RenderHTML("/references/references.html", map[string]interface {} {

		"pageTitle" : "Справочники",

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

func editItems(c *app.Context) {

	f := form.New("Элементы справочника", "register", "login_salt")

	c.RenderJSON(f)

}
