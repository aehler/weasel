package registry

import (
	"github.com/jmoiron/sqlx"
	"weasel/app/session"
	"weasel/app/db"
)

var Registry struct {
		Connect *sqlx.DB
		Session *session.SessionStorage
		SessionKeys []*[32]byte
		ReferenceConf map[string]*refConf
	}

func Init(config string) {

	Registry.Connect = db.New(config)

	Registry.Session = session.Init()

	Registry.SessionKeys = append(Registry.SessionKeys, &[32]byte{
			'm',
		} )

	readRefConf(config)
}
