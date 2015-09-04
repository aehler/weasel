package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type MySQL struct {
	Address  string `yaml:"address"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

func (p *MySQL) ConnectString() string {

	return fmt.Sprintf("%s:%s@%s/%s", p.Username, p.Password, p.Address, p.Database)
}

func (p *MySQL) Connect() (*sqlx.DB, error) {

	return sqlx.Open("mysql", p.ConnectString())
}
