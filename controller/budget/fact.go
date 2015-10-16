package budget

import (
	"weasel/app"
	"weasel/app/form"
	"weasel/app/grid"
	"weasel/lib/budget"
	"weasel/middleware/auth"
	"fmt"
)


func fact(c *app.Context) {

	var (
		//user = c.Get("user").(auth.User)
	)

	c.RenderHTML("/budget/grid.html", map[string]interface {} {

		"pageTitle" : "Бюджет",
		"gridURL" : "/budget/fact/grid/",

	})
}

func factGrid(c *app.Context) {

	user := c.Get("user").(auth.User)

	rows, err := budget.GetFactRows(user.OrganizationId)

	if err != nil {
		c.RenderJSON(map[string]interface {}{
			"Error" : err.Error(),
		})

		return
	}

	g := grid.New(rows)

	g.Column(
		&grid.Column{
		Name : "Date",
		Label: "Дата операции",
		Cell : grid.CellTypeString,
		Order: 0,
	},
		&grid.Column{
		Name : "Sum",
		Label: "Сумма операции",
		Cell : grid.CellTypeString,
		Order: 1,
	},
		&grid.Column{
		Name : "cost_items",
		Label: "Статья бюджета",
		Cell : grid.CellTypeString,
		Order: 2,
	},
		&grid.Column{
		Name : "importance",
		Label: "Приоритет платежа",
		Cell : grid.CellTypeString,
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

func factForm(c *app.Context) {

	user := c.Get("user").(auth.User)
	//id, _ = crypto.DecryptUrl(c.Params.ByName("rowId"))

	post := form.New("Операция", "", user.SessionID)

	post.Action = fmt.Sprintf("/budget/fact/edit/%s/", c.Params.ByName("rowId"))

	post.Fields(
		&form.Element{
			Name : "Sum",
			Label : "Сумма операции",
			Order : 0,
			TypeName : "number",
			Type : form.Number,
		},
		&form.Element{
			Name : "Date",
			Label : "Дата операции",
			Order : 2,
			TypeName : "date",
			Type : form.Date,
		},
		&form.Element{
			Name : "Tags",
			Label : "Тэги",
			Order : 990,
			TypeName : "taglist",
			Type : form.TagList,
		},
	)

	fact := budget.NewFact(&user)

	if err := fact.DimensionOptions(); err != nil {

		c.RenderJSON(map[string]interface {}{
			"Error" : err.Error(),
		})

		return

	}

	for i, dim := range *fact.Dimensions {

		post.Fields(
			&form.Element{
				Name : dim.ReferenceAlias,
				Label : dim.ReferenceLabel,
				Order : uint(i+100),
				TypeName : "select",
				Type : form.Select,
				Options : dim.Options,
			},
		)

	}

	if c.IsPost() {

		if err := post.ParseForm(fact, c.Request); err != nil {

			c.RenderJSON(map[string]interface{}{
				"Error" : err.Error(),
			})

			c.Stop()

			return

		}

		fact.Dimensions.MapValues(post.Values())

		if err := fact.Save(); err != nil {

			c.RenderJSON(map[string]interface{}{
			"Error" : err.Error(),
		})

			c.Stop()

			return

		}

	}

	c.RenderJSON(post.Context())
}
