package budget

import (
	"weasel/app/grid"
	"weasel/app/crypto"
	"weasel/app/registry"
	"reflect"
	"fmt"
)

func (r FactRows) GridRows(cols []*grid.Column) []map[string]interface {}{

	rows := []map[string]interface {}{}

	for _, row := range r.Items {

		cr := map[string]interface {}{}

		for _, c := range cols {

			val := reflect.ValueOf(row).Elem().FieldByName(c.Name)

			if val.IsValid() {

				cr[c.Name] = val.Interface()

			} else {

				d := reflect.ValueOf(row).Elem().FieldByName("Dimensions").Elem()

				for i:=0; i<d.Len(); i++ {

					if d.Index(i).Elem().FieldByName("ReferenceAlias").String() == c.Name {

						cr[c.Name] = d.Index(i).Elem().FieldByName("Label").Interface()

					}

				}
			}

		}

		cr["actions"] = []map[string]interface {}{
			map[string]interface {}{"Редактировать" : map[string]interface{}{
				"href" : fmt.Sprintf("/budget/fact/edit/%s", crypto.EncryptUrl(row.FactID)),
				"target" : "jsForm",
			},
			}}

		rows = append(rows, cr)

	}

	return rows
}

func GetFactRows(oid uint) (FactRows, error) {

	r := FactRows{
		Items : []*Fact{},
	}

	if err := registry.Registry.Connect.Select(&r.Items, `select id, sum, date_op, dims_meta, user_meta
		from weasel_main.budget_operations where organization_id = $1
		order by date_op desc`,
		oid,
	); err != nil {
		return r, err
	}

	return r, nil
}
