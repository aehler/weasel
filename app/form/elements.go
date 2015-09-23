package form

import (
	"strings"
)

type Element struct {
	Name string
	HashName string `json:"n"`
	Value []string `json:"v"`
	Label string `json:"l"`
	PlaceHolder string
	Description string
	Type uint `json:"-"`
	TypeName string `json:"t"`
	Order uint
	Options Options `json:"o"`
}

func (e *Element) GetValue() string {
	
	return strings.Join(e.Value, ", ")
	
}
