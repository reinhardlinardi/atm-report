package history

type Repository interface {
	IsExist(atmId, date string) (bool, error)
	Insert(atmId, date string) (int64, error)
}
