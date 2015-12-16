package budget

import (
	"weasel/lib/references"
	"weasel/app/registry"
)

type Plan struct {
	PlanSum float64
	FactSum float64
	Dimensions *references.Dimensions
}

func GetPlanRows(oid, periodId uint) (FactRows, error) {

	r := FactRows{
		Items : []*Fact{},
	}

	if err := registry.Registry.Connect.Select(&r.Items, `select * from weasel_operations.get_plan_items($1, $2)`,
		oid,
		periodId,
	); err != nil {
		return r, err
	}

	return r, nil
}
