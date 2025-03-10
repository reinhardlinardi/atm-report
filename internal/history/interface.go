package history

type Repository interface {
	Check(atmId, date string, seq int) (exist bool, err error)
	Append(atmId, date string, seq int) (int64, error)
}
