package budget

import (
	"weasel/lib/references"
	"weasel/middleware/auth"
)

type Plan struct {
	PlanSum float64
	FactSum float64
	Dimensions *references.Dimensions
}

func MonthlyPlan(user *auth.User) ([]*Plan, error) {

	res := []*Plan{}

	return res, nil
}
