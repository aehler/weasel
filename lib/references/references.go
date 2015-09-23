package references

import (
	"weasel/app/registry"
	"weasel/app/form"
	"weasel/app/crypto"
	"weasel/app/grid"
	"encoding/json"
	"strings"
	"reflect"
	"strconv"
	"sort"
	"fmt"
)

type Reference struct {
	ID uint `db:"id"`
	OrganizationId uint `db:"organization_id"`
	Name string `db:"name" weaselform:"name" formLabel:"Наименование справочника"`
	Alias string `db:"alias"`
	Total uint `db:"total"`
	Blocked bool `db:"blocked" weaselform:"blocked" formLabel:"Заблокирован"`
	Items Items `db:"-"`
	Meta referenceMeta `db:"-"`
}

type referenceMeta struct {
	Type string
	Fields []*form.Element
}

type References struct {
	R []*Reference
	grid.Grider
}

func New(rid, oid uint) (*Reference, error) {

	r := &Reference{}

	if err := registry.Registry.Connect.Get(r, `select id, organization_id, name, alias, total, blocked
	from weasel_classifiers."references"
	left join weasel_classifiers."counter" on reference_id = id
	where organization_id = $1 and id = $2`,
		oid,
		rid,
	); err != nil {
		return r, err
	}

	if err := r.mapConfig(); err != nil {
		return r, err
	}

	defer func(ref *Reference) {
		ref = nil
	}(r)

	return r, nil
}

func ByAlias(oid uint, alias... string) (map[string]*Reference, error) {

	r := []*Reference{}
	res := map[string]*Reference{}

	if err := registry.Registry.Connect.Select(&r, `select
	id, name, organization_id, blocked, total, alias from weasel_classifiers."references"
      left join weasel_classifiers.counter on id = reference_id
      where organization_id = $1 and alias = any($2::text[])`,
		oid,
		fmt.Sprintf(`{"%s"}`, strings.Join(alias, `","`)),
	); err != nil {

		fmt.Println(err)

		return res, err

	}

	for _, al := range alias {

		res[al] = &Reference{}

		for _, ref := range r {

			if ref.Alias == al {

				res[al] = ref

			}

		}

	}

	return res, nil
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

func (r *Reference) ItemsList() (*Items, error) {

	its := Items{
		Items : []*Item{},
	}

	if err := registry.Registry.Connect.Select(&its.Items, `WITH RECURSIVE search_graph(id, pid, name, reference_id, fields, is_group , depth, path, cycle) AS (
        SELECT g.id, g.pid, g.name,
		  g.reference_id, g.fields, g.is_group,
		   1 as depth,
          string_to_array(g.parents, '.')::bigint[] as path,
          false
        FROM weasel_classifiers.items g
        where g.reference_id = $1 and g.pid = 0 and ver = 0
      UNION ALL
        SELECT g.id, g.pid, g.name, g.reference_id, g.fields, g.is_group, sg.depth + 1,
           string_to_array(g.parents, '.')::bigint[],
          g.id = ANY(sg.path)
        FROM weasel_classifiers.items g, search_graph sg
        WHERE g.pid = sg.id and g.ver = 0 AND NOT cycle
)
SELECT id, pid, name, reference_id, fields, is_group , depth, path as parents FROM search_graph;`,
		r.ID,
	); err != nil {
		return &its, err
	}

	sort.Sort(&its)

	its.Items = its.sortTree(uint(0))

	return &its, nil
}

func (r *Reference) SaveItem(item *Item, vals map[string]string) error {

	fb, err := json.Marshal(vals)
	if err != nil {
		return err
	}

	if item.ID == 0 {

		pid, err := strconv.ParseUint(vals["pid"], 10, 64)
		if err != nil {
			return err
		}

		it := &Item{
			ID : 0,
			RefId : r.ID,
			Label : vals["title"],
			Alias : crypto.Unique(),
			Parents : "",
			Pid : uint(pid),
			Fields : string(fb),
			IsGroup : vals["is_group"] == "true",
			Version : 0,
		}

		if err := it.save(); err != nil {
			return err
		}

		return nil
	}

	item.Fields = string(fb)
	item.IsGroup = vals["is_group"] == "true"
	item.Label = vals["title"]

	if err := item.save(); err != nil {
		return err
	}

	return nil
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
			map[string]interface {}{"Редактировать" : map[string]interface{}{
				"href" : fmt.Sprintf("/settings/references/items/%s", crypto.EncryptUrl(row.ID))},

			}}

		rows = append(rows, cr)

	}

	return rows
}
