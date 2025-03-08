package db

import "database/sql"

func (db *DBImpl) QueryRow(dest any, query string, args ...any) error {
	return db.conn.Get(dest, query, args...)
}

func (db *DBImpl) InsertRow(query string, args ...any) (int64, error) {
	res, err := db.exec(query, args...)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (db *DBImpl) Query(dest any, query string, args ...any) error {
	return db.conn.Select(dest, query, args)
}

func (db *DBImpl) Exec(query string, args ...any) (int64, error) {
	res, err := db.exec(query, args...)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func (db *DBImpl) exec(query string, args ...any) (sql.Result, error) {
	return db.conn.Exec(query, args...)
}
