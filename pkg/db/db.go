package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DBImpl struct {
	conn   *sqlx.DB
	config *Config
}

type Config struct {
	User   string
	Pass   string
	Host   string
	Port   uint16
	Schema string
}

func New(config *Config) *DBImpl {
	return &DBImpl{config: config}
}

func (db *DBImpl) Open() error {
	c := db.config

	conn, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@(%s:%d)/%s", c.User, c.Pass, c.Host, c.Port, c.Schema))
	if err != nil {
		fmt.Printf("err connect db: %s\n", err.Error())
		return err
	}

	db.conn = conn
	return nil
}

func (db *DBImpl) Close() {
	db.conn.Close()
}
