package budget

import (
	"weasel/app"
	"weasel/lib/budget"
	"weasel/middleware/auth"
)

func plan(c *app.Context) {

	user := c.Get("user").(auth.User)

	plan, _ := budget.MonthlyPlan(&user)

	c.RenderHTML("/budget/plan.html", map[string]interface {} {

		"pageTitle" : "Бюджет - план",
		"plan" : plan,
	})
}
