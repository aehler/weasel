package registry

import (
	"github.com/jmoiron/sqlx"
	"weasel/app/session"
)

var Registry struct {
		Connect *sqlx.DB
		Session *session.SessionStorage
	}
