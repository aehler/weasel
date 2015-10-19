package budget

import (
	"weasel/lib/references"
	"weasel/middleware/auth"
	"weasel/app/registry"
	"encoding/json"
	"time"
)


func NewFact(user *auth.User) *Fact {

	return &Fact{
		FactID : 0,
		Sum : 0,
		Tags : []string{},
		Date : time.Now(),
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

func (f *Fact) FactById(id uint) error {

	return registry.Registry.Connect.Get(f, `select id, sum, date_op, dims_meta, user_meta
		from weasel_main.budget_operations where id = $1
		order by date_op desc`, id)
}

func (f *Fact) DimensionOptions() error {

	return references.DimOptions(f.Dimensions, f.User.OrganizationId)
}

func (f *Fact) Save() error {

	var (
		fd, fu []byte
		r uint
	)


	fd, err := json.Marshal(f.Dimensions)

	if err != nil {

		return err

	}

	fu, err = json.Marshal(f.User)

	if err != nil {

		return err

	}

	if err := registry.Registry.Connect.Get(&r, `select * from weasel_main.save_budget_operation($1, $2, $3, $4, $5, $6, $7)`,
		f.FactID,
		f.User.UserID,
		f.User.OrganizationId,
		f.Sum,
		f.Date.Format("2006-01-02"),
		string(fu),
		string(fd),
	); err != nil {

		return err

	}

	return nil

}
