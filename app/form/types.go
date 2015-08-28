package form

const (
	Text uint = iota + 1
	TextArea
	Select
	MultipleSelect
	CheckboxGroup
	Date
	Password
)

var elementType = map[string]uint {
	"text" : Text,
	"textarea" : TextArea,
	"select" : Select,
	"multipleselect" : MultipleSelect,
	"checkbox" : CheckboxGroup,
	"date" : Date,
	"password" : Password,
}
