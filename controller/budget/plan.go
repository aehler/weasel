package budget

import (
	"weasel/app"
	"weasel/app/form"
	"weasel/app/grid"
	"weasel/app/crypto"
	"weasel/lib/budget"
	"weasel/lib/references"
	"weasel/middleware/auth"
	"strconv"
	"fmt"
)

func plan(c *app.Context) {

	user := c.Get("user").(auth.User)

	dims := references.Dimensions{
		&references.Dimension{
			ReferenceAlias : "periods",
		},
	}

	references.DimOptions(&dims, user.OrganizationId)

	c.RenderHTML("/budget/plan.html", map[string]interface {} {

		"pageTitle" : "Бюджет - план",
		"gridURL" : "/budget/plan/grid/",
		"gridControls" : map[string]interface {}{
			"controls" : []grid.GridControl{grid.ControlPeriod},
			"context" : dims,
		},
	})
}

func planGrid(c *app.Context) {

	user := c.Get("user").(auth.User)

	periodId, err := strconv.ParseUint(c.Params.ByName("p"), 10, 64)
	if err != nil {

		periodId = 0

	}

	rows, err := budget.GetPlanRows(user.OrganizationId, uint(periodId))

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

func planForm(c *app.Context) {

	user := c.Get("user").(auth.User)

	id, _ := crypto.DecryptUrl(c.Params.ByName("rowId"))

	fact := budget.NewFact(&user)

	if err := fact.FactById(id); err != nil {

		c.RenderJSON(map[string]interface {}{
		"Error" : err.Error(),
	})

		return
	}

	post := form.New("Операция", "", user.SessionID)

	post.Action = fmt.Sprintf("/budget/plan/edit/%s/", c.Params.ByName("rowId"))

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

	if err := post.SetValues(fact); err != nil {

		c.RenderJSON(map[string]interface {}{
		"Error" : err.Error(),
	})

		return
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
