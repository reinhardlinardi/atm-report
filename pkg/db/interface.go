package db

type DB interface {
	QueryRow(dest any, query string, args ...any) error
	InsertRow(query string, args ...any) (int64, error)
	Query(dest any, query string, args ...any) error
	Exec(query string, args ...any) (int64, error)
}
