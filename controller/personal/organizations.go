package personal

import (
	"weasel/app"
	"weasel/app/grid"
	"weasel/lib/auth"
	"weasel/lib/organizations"
	"fmt"
)

func listOrgs(c *app.Context) {

	c.RenderHTML("/organizations/organizations.html", map[string]interface {} {

		"pageTitle" : "Организации",
		"gridURL" : "/settings/organizations/grid/",
	})

}

func organizationsGrid(c *app.Context) {

	user := c.Get("user").(auth.User)

	rows, err := organizations.GetOrganizations(user.OrganizationId)

	if err != nil {
		c.RenderJSON(map[string]interface {}{
		"Error" : err.Error(),
	})

		return
	}

	g := grid.New(rows)

	g.Column(
		&grid.Column{
		Name : "Shortname",
		Label: "Наименование",
		Cell : grid.CellTypeString,
		Order: 0,
	},
		&grid.Column{
		Name : "INN",
		Label: "ИНН",
		Cell : grid.CellTypeString,
		Order: 1,
	},
		&grid.Column{
		Name : "KPP",
		Label: "КПП",
		Cell : grid.CellTypeString,
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

func organizationsEdit(c *app.Context) {

	user := c.Get("user").(auth.User)

	o := &organizations.Organization{}

	f, err := o.Form(user)
	if err != nil {

		c.RenderHTML("/errors/500.html", map[string]interface {} {
			"Error" : err.Error(),
		})

		return
	}

	fmt.Println(f.Context())

	c.RenderHTML("/organizations/form.html", map[string]interface {} {

		"pageTitle" : "Организация",
		"form" : f.Context(),

	})

}
