package budget

import (
	"weasel/lib/references"
	"weasel/middleware/auth"
	"weasel/app/registry"
	"weasel/app/grid"
	"encoding/json"
	"time"
)


func (f *FactRows)GridRows(c []*grid.Column) []map[string]interface {} {

	return []map[string]interface {}{}
}

func NewFact(user *auth.User) *Fact {

	return &Fact{
		FactID : 0,
		Sum : 0,
		Tags : []string{},
		Date : "",
		User : user,
		Dimensions : &references.Dimensions{
			&references.Dimension{
				ReferenceAlias : "currency",
			},
			&references.Dimension{
				ReferenceAlias : "cost_items",
			},
			&references.Dimension{
				ReferenceAlias : "importance",
			},
			&references.Dimension{
				ReferenceAlias : "accounts",
			},
			&references.Dimension{
				ReferenceAlias : "project",
			},
		},
	}

}

func (f *Fact) DimensionOptions() error {

	return references.DimOptions(f.Dimensions, f.User.OrganizationId)
}

func (f *Fact) Save() error {

	var (
		fd, fu []byte
		r uint
		t time.Time
	)


	fd, err := json.Marshal(f.Dimensions)

	if err != nil {

		return err

	}

	fu, err = json.Marshal(f.User)

	if err != nil {

		return err

	}

	if t, err = time.Parse("02.01.2006", f.Date); err != nil {

		t = time.Time{}

	}

	if err := registry.Registry.Connect.Get(&r, `select * from weasel_main.save_budget_operation($1, $2, $3, $4, $5, $6, $7)`,
		f.FactID,
		f.User.UserID,
		f.User.OrganizationId,
		f.Sum,
		t.Format("2006-01-02"),
		string(fu),
		string(fd),
	); err != nil {

		return err

	}

	return nil

}
