package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DB struct {
	Conn   *sqlx.DB
	Config *Config
}

type Config struct {
	User   string
	Pass   string
	Host   string
	Port   uint16
	Schema string
}

func New(config *Config) *DB {
	return &DB{Config: config}
}

func (db *DB) Open() error {
	c := db.Config

	conn, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@(%s:%d)/%s", c.User, c.Pass, c.Host, c.Port, c.Schema))
	if err != nil {
		fmt.Printf("err connect db: %s\n", err.Error())
		return err
	}

	db.Conn = conn
	return nil
}

func (db *DB) Close() {
	db.Conn.Close()
}
