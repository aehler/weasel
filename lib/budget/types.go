package budget

import (
	"weasel/app/grid"
	"weasel/middleware/auth"
	"weasel/lib/references"
	"time"
)

type Fact struct {
	FactID uint `db:"id"`
	Sum float64 `db:"sum" weaselform:"sum"`
	Tags []string `db:"-" weaselform:"tags"`
	Date time.Time `db:"date_op" weaselform:"date"`
	Dimensions *references.Dimensions `db:"dims_meta"`
	User *auth.User `db:"user_meta"`
}

type FactRows struct {
	Items []*Fact
	grid.Grider
}
