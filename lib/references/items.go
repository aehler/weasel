package references

import (
	"weasel/app/grid"
)

type Item struct {
	ID uint
	Label string
	Version uint
	Parents string
}

type Items struct {
	Items []Item
	grid.Grider
}

func ItemsList(oid, refid uint) (*Items, error) {

	its := Items{}

	return &its, nil
}

func (r *Reference) ReferenceItems() {

}

func (i *Items) GridRows(cols []*grid.Column) []map[string]interface {}{

	return []map[string]interface {}{}
}
