package references

import (
	"weasel/app/grid"
)

type Items struct {
	Items []Item
	grid.Grider
}

func (r *Reference) ReferenceItems() {

}

func (i *Items) GridRows(cols []*grid.Column) []map[string]interface {}{

	return []map[string]interface {}{}
}
