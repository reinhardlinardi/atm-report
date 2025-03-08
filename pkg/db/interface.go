package db

type DB interface {
	QueryRow(dest any, query string, args ...any) error
	InsertRow(query string, args ...any) (id int64, err error)
	Query(dest any, query string, args ...any) error
	Exec(query string, args ...any) (rowsAffected int64, err error)
}
