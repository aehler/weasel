package references

import (
	"weasel/app"
	"weasel/app/crypto"
	"weasel/app/grid"
	"weasel/app/form"
	"weasel/lib/references"
	"weasel/lib/auth"
	"fmt"
)

func list(c *app.Context) {

	c.RenderHTML("/references/references.html", map[string]interface {} {

		"pageTitle" : "Справочники",
		"gridURL" : "/settings/references/grid/",

	})

}

func gridJSON(c *app.Context) {

	user := c.Get("user").(auth.User)

	refList, err := references.ReferenceList(user.OrganizationId)
	if err != nil {

		c.RenderJSON(map[string]interface {}{
			"Error" : err.Error(),
		})

		c.Stop()

		return
	}

	g := grid.New(refList)

	g.Column(
		&grid.Column{
			Name : "ID",
			Label: "#",
			Cell : grid.CellTypeInt,
			Order: 0,
		},
		&grid.Column{
			Name : "Name",
			Label: "Наименование справочника",
			Cell : grid.CellTypeString,
			Order: 1,
		},
		&grid.Column{
			Name : "Total",
			Label: "Элементов",
			Cell : grid.CellTypeInt,
			Order: 2,
		},
		&grid.Column{
			Name : "actions",
			Label: "..",
			Cell : grid.CellTypeActions,
			Order: 999,
		},
	)

	c.RenderJSON(g.Context())
}

func edit(c *app.Context) {

	c.RenderHTML("/blank.html", map[string]interface {} {

	})

}

func itemsList(c *app.Context) {

	var (
		user = c.Get("user").(auth.User)
		refid, _ = crypto.DecryptUrl(c.Params.ByName("refId"))
	)

	ref, err := references.New(refid, user.OrganizationId)
	if err != nil {

		c.RenderHTML("/errors/500.html", map[string]interface {}{
		"Error" : err.Error(),
	})

		c.Stop()

		return
	}

	c.RenderHTML("/references/reference_items.html", map[string]interface {} {

		"pageTitle" : "Элементы справочника",
		"gridURL" : fmt.Sprintf("/settings/references/items_grid/%s/", c.Params.ByName("refId")),
		"ref" : ref,

	})
}

func itemsGridJSON(c *app.Context) {

	var (
		user = c.Get("user").(auth.User)
		refid, _ = crypto.DecryptUrl(c.Params.ByName("refId"))
	)

	ref, err := references.New(refid, user.OrganizationId)
	if err != nil {

		c.RenderJSON(map[string]interface {}{
		"Error" : err.Error(),
	})

		c.Stop()

		return
	}

	itemsList, err := ref.ItemsList()
	if err != nil {

		c.RenderJSON(map[string]interface {}{
		"Error" : err.Error(),
		})

		c.Stop()

		return
	}

	g := grid.New(itemsList)

	g.Column(
		&grid.Column{
		Name : "ID",
		Label: "#",
		Cell : grid.CellTypeInt,
		Order: 0,
	},
		&grid.Column{
		Name : "Label",
		Label: "Значение",
		Cell : grid.CellTypeStringWithOffset,
		Order: 2,
	},
		&grid.Column{
		Name : "actions",
		Label: "..",
		Cell : grid.CellTypeActions,
		Order: 999,
	},
	)

	c.RenderJSON(g.Context())

}

func editItem(c *app.Context) {

	var (
		user = c.Get("user").(auth.User)
		ssid = c.Get("ssid").(string)
		rid, _ = crypto.DecryptUrl(c.Params.ByName("refId"))
		id, _ = crypto.DecryptUrl(c.Params.ByName("itemId"))
	)

	f := form.New("Элемент", "register", ssid)

	f.Action = fmt.Sprintf("/settings/references/item_edit/%s/%s/", c.Params.ByName("refId"), c.Params.ByName("itemId"))

	ref, err := references.New(rid, user.OrganizationId)
	if err != nil {

		c.RenderJSON(map[string]interface {}{
		"Error" : err.Error(),
	})

		c.Stop()

		return
	}

	item, err := references.GetItem(id)
	if err != nil {

		c.RenderJSON(map[string]interface {}{
			"Error" : err.Error(),
		})

		c.Stop()

		return
	}

	f.Fields(ref.Meta.Fields...)

	if ref.Meta.Type == "tree" {
		f.Fields(&form.Element{
			Name : "is_group",
			Label : "Это группа",
			Order : 990,
			TypeName : "bool",
			Type : form.Checkbox,
	},
		)
	}

	if err := f.UnmarshalValues(item.Fields); err != nil {

		c.RenderJSON(map[string]interface{}{
			"Error" : err.Error(),
		})

		c.Stop()

		return
	}

	if c.IsPost() {

		if err := f.ParseForm(nil, c.Request); err != nil {

		c.RenderJSON(map[string]interface{}{
			"Error" : err.Error(),
		})

			c.Stop()

			return

		}

		fv := f.Values()

		if err := ref.SaveItem(item, fv); err != nil {

			c.RenderJSON(map[string]interface{}{
			"Error" : err.Error(),
		})

			c.Stop()

			return

		}

		c.RenderJSON(map[string]interface{}{
		"s" : true,
	})

		c.Stop()

		return
	}

	c.RenderJSON(f.Context())

}

func addItem(c *app.Context) {

	var (
		user = c.Get("user").(auth.User)
		ssid = c.Get("ssid").(string)
		rid, _ = crypto.DecryptUrl(c.Params.ByName("refId"))
		pid, _ = crypto.DecryptUrl(c.Params.ByName("itemId"))
	)

	ref, err := references.New(rid, user.OrganizationId)
	if err != nil {

		c.RenderJSON(map[string]interface {}{
			"Error" : err.Error(),
		})

		c.Stop()

		return
	}

	f := form.New("Новый элемент", "ref_element", ssid)

	f.Action = fmt.Sprintf("/settings/references/item_add/%s/%s/", c.Params.ByName("refId"), c.Params.ByName("itemId"))

	f.Fields(ref.Meta.Fields...)

	if ref.Meta.Type == "tree" {
			f.Fields(&form.Element{
				Name : "is_group",
				Label : "Это группа",
				Order : 990,
				TypeName : "bool",
				Type : form.Checkbox,
			},
		)
	}

	if c.IsPost() {

		if err := f.ParseForm(nil, c.Request); err != nil {

		c.RenderJSON(map[string]interface{}{
			"Error" : err.Error(),
		})

			c.Stop()

			return

		}

		fv := f.Values()

		fv["pid"] = fmt.Sprintf("%d",pid)

		if err := ref.SaveItem(&references.Item{ID : 0}, fv); err != nil {

			c.RenderJSON(map[string]interface{}{
				"Error" : err.Error(),
			})

			c.Stop()

			return

		}

		c.RenderJSON(map[string]interface{}{
			"s" : true,
		})

		c.Stop()

		return
	}

	c.RenderJSON(f.Context())

}
