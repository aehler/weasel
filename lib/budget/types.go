package budget

import (
	"weasel/middleware/auth"
	"weasel/lib/references"
)

type Fact struct {
	FactID uint `db:"id"`
	Sum float64 `db:"sum" weaselform:"sum"`
	Tags []string `db:"tags" weaselform:"tags"`
	Date string `db:"date_op" weaselform:"date"`
	Dimensions *references.Dimensions
	User *auth.User
}

type FactRows struct {

}

