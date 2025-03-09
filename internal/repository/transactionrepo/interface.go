package transactionrepo

type Repository interface {
	InsertRows(data []Transaction) (int64, error)
}
