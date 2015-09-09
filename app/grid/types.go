package grid

type CellType string

const (
	CellTypeString CellType = "string"
	CellTypeInt CellType = "integer"
	CellTypeFloat CellType = "number"
	CellTypeDate CellType = "date"
	CellTypeUri CellType = "uri"
	CellTypeActions CellType = "actions"
)

type Column struct {
	Name string `json:"name"`
	Label string `json:"label"`
	Editable string `json:"editable"`
	Cell CellType `json:"cell"`
	Order int16   `json:"-"`
}
