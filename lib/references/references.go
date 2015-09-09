package references

import (
	"weasel/app/registry"
	"weasel/app/crypto"
	"weasel/app/grid"
	"reflect"
	"fmt"
)

type Reference struct {
	ID uint `db:"id"`
	OrganizationId uint `db:"organization_id"`
	Name string `db:"name" weaselform:"name" formLabel:"Наименование справочника"`
	Alias string `db:"alias"`
	Total uint `db:"total"`
	Blocked bool `db:"blocked" weaselform:"blocked" formLabel:"Заблокирован"`
}

type References struct {
	R []*Reference
	grid.Grider
}

func ReferenceList(oid uint) (References, error) {

	r := References{}

	if err := registry.Registry.Connect.Select(&r.R, `select
	id, name, organization_id, blocked, total, alias from weasel_classifiers."references"
      left join weasel_classifiers.counter on id = reference_id
      where organization_id = $1`,
		oid,
	); err != nil {

		return r, err

	}

	return r, nil

}

func (r References) GridRows(cols []*grid.Column) []map[string]interface {}{

	rows := []map[string]interface {}{}

	for _, row := range r.R {

		cr := map[string]interface {}{}

		for _, c := range cols {

			val := reflect.ValueOf(row).Elem().FieldByName(c.Name)

			if val.IsValid() {

				cr[c.Name] = val.Interface()
			}
		}

		cr["actions"] = []map[string]interface {}{
			map[string]interface {}{"Редактировать" : fmt.Sprintf("/settings/references/edit_items/%s", crypto.EncryptUrl(row.ID))},
		}

		rows = append(rows, cr)

	}

	return rows
}
