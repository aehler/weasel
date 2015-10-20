package registry

import (
	"github.com/jmoiron/sqlx"
	"weasel/app/session"
	"weasel/app/db"
)

type registry struct {
	Connect *sqlx.DB
	Session *session.SessionStorage
	SessionKeys []*[32]byte
	ReferenceConf map[string]*refConf
	storage map[string]*storage
}

var Registry registry

func Init(config string) {

	Registry.Connect = db.New(config)

	Registry.Session = session.Init()

	Registry.SessionKeys = append(Registry.SessionKeys, &[32]byte{
			'm',
		} )

	readRefConf(config)

	readStorageConf(config)

}

func (r *registry) Storage(key string) *storage {

	return Registry.storage[key]

}
