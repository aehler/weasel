package organizations

import (
	"weasel/app/grid"
	"weasel/app/crypto"
	"weasel/app/registry"
	"reflect"
	"fmt"
)

func (r OrganizationRow) GridRows(cols []*grid.Column) []map[string]interface {}{

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
				"href" : fmt.Sprintf("/settings/organizations/edit/%s", crypto.EncryptUrl(row.ID)),
			},
			}}

		rows = append(rows, cr)

	}

	return rows
}

func GetOrganizations(oid uint) (OrganizationRow, error) {

	r := OrganizationRow{
		Items : []*StoredOrganization{},
	}

	if err := registry.Registry.Connect.Select(&r.Items, `select id, inn, kpp, organization_id, user_id, meta_info
		from weasel_main.organizations where organization_id = $1
		order by created_at desc`,
		oid,
	); err != nil {
		return r, err
	}

	return r, nil
}
