package references

import (
	"weasel/app/form"
	"weasel/app/grid"
	"weasel/app/crypto"
	"weasel/app/registry"
	"reflect"
	"fmt"
)

type Item struct {
	ID uint `db:"id"`
	RefId uint `db:"reference_id"`
	Label string `db:"name"`
	Alias string `db:"alias"`
	Parents string `db:"parents"`
	Pid uint `db:"pid"`
	Fields string `db:"fields"`
	IsGroup bool `db:"is_group"`
	Version uint `db:"ver"`
	Depth uint `db:"depth"`
}

type Items struct {
	Items []*Item
	grid.Grider
	form.Selecter
}

func GetItem(id uint) (*Item, error) {

	i := Item{}

	if err := registry.Registry.Connect.Get(&i, `select id, reference_id, name, alias, parents, pid, fields, is_group, ver from weasel_classifiers.items where id = $1`, id); err != nil {
		return &i, err
	}

	return &i, nil
}

func (i *Item) save() error {

	var r uint

	return registry.Registry.Connect.Get(&r, `select 1 from weasel_classifiers.save_classifier_item($1, $2, $3, $4, $5, $6, $7)`,
		i.ID,
		i.RefId,
		i.Label,
		i.Alias,
		i.Pid,
		i.Fields,
		i.IsGroup,
	)

}

func (i *Items) Opts() form.Options {

	opts := form.Options{}

	for _, it := range i.Items {

		opts = append(opts, &form.Option{
				Value : it.ID,
				Label : it.Label,
			})

	}

	return opts
}

func (i *Items) GridRows(cols []*grid.Column) []map[string]interface {}{

	rows := []map[string]interface {}{}

	for _, row := range i.Items {

		cr := map[string]interface {}{}

		for _, c := range cols {

			val := reflect.ValueOf(row).Elem().FieldByName(c.Name)

			if val.IsValid() {

				cr[c.Name] = val.Interface()
			}
		}

		acts := []map[string]interface {}{
			map[string]interface {}{"Редактировать" : map[string]interface{}{
				"href" : fmt.Sprintf("/settings/references/item_edit/%s/%s/", crypto.EncryptUrl(row.RefId), crypto.EncryptUrl(row.ID)),
				"target" : "jsForm",
			}},
		}
		if row.IsGroup {
			acts = append(acts,
				map[string]interface {}{"Добавить в группу" :  map[string]interface{}{
					"href" : fmt.Sprintf("/settings/references/item_add/%s/%s/", crypto.EncryptUrl(row.RefId), crypto.EncryptUrl(row.ID)),
					"target" : "jsForm",
				}},
			)
		}

		cr["actions"] = acts
		cr["depth"] = row.Depth


		rows = append(rows, cr)

	}

	return rows
}

func (its *Items)sortTree(i uint) []*Item {

	ni := []*Item{}

	for _, it := range its.Items {

		if i == it.Pid {

			ni = append(ni, it)

			ni = append(ni, its.sortTree(it.ID)...)

		}
	}

	return ni
}

func (its *Items) Len() int {
	return len(its.Items)
}

func (its *Items) Swap(i, j int) {
	its.Items[i], its.Items[j] = its.Items[j], its.Items[i]
}

func (its *Items) Less(i, j int) bool {
	return its.Items[i].Label < its.Items[j].Label && its.Items[i].Pid == its.Items[j].Pid
}
