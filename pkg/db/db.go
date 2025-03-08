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
	conf := db.config
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s", conf.User, conf.Pass, conf.Host, conf.Port, conf.Schema)

	conn, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return err
	}

	db.conn = conn
	return nil
}

func (db *DBImpl) Close() {
	db.conn.Close()
}
