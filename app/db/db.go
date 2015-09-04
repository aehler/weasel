package db

import (
	"github.com/jmoiron/sqlx"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type dbcreds struct {
	Address  string `yaml:"address"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

func New(config string)  *sqlx.DB {

	data, err := ioutil.ReadFile(config)

	if err != nil {

		log.Fatal(err.Error())
	}

	rr := map[string]dbcreds{}

	if err := yaml.Unmarshal(data, &rr); err != nil {

		log.Fatal(err.Error())
	}

	r := &PostgreSQL{
		Address : rr["postgresql"].Address,
		Username : rr["postgresql"].Username,
		Password : rr["postgresql"].Password,
		Database : rr["postgresql"].Database,
	}

	s, err := r.Connect()

	if err != nil {

		log.Fatal(err.Error())

	}

	return s
}
