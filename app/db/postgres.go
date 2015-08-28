package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgreSQL struct {
	Address  string `yaml:"address"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	ShardID  int16  `yaml:"shard_id"`
}

func (p *PostgreSQL) ConnectString() string {

	return fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable&connect_timeout=1", p.Username, p.Password, p.Address, p.Database)
}

func (p *PostgreSQL) Connect() (*sqlx.DB, error) {

	return sqlx.Open("postgres", p.ConnectString())

}
