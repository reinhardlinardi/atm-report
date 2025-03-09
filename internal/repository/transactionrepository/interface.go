package transactionrepository

import "github.com/reinhardlinardi/atm-report/model"

type Repository interface {
	InsertRows(data []model.Transaction) (int64, error)
}
