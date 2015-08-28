package db

import (
	"github.com/jmoiron/sqlx"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func New(config string)  *sqlx.DB {

	data, err := ioutil.ReadFile(config)

	if err != nil {

		log.Fatal(err.Error())
	}

	r := MySQL{}

	if err := yaml.Unmarshal(data, &r); err != nil {

		log.Fatal(err.Error())
	}

	s, err := r.Connect()

	if err != nil {

		log.Fatal(err.Error())

	}

	return s
}
