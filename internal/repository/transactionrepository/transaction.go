package transactionrepository

import (
	"strings"

	"github.com/reinhardlinardi/atm-report/model"
	"github.com/reinhardlinardi/atm-report/pkg/db"
)

type RepositoryImpl struct {
	conn db.DB
}

func New(conn db.DB) *RepositoryImpl {
	return &RepositoryImpl{conn: conn}
}

func (rp *RepositoryImpl) InsertRows(data []model.Transaction) (int64, error) {
	if len(data) == 0 {
		return 0, nil
	}

	args := []any{}

	var sb strings.Builder
	sb.WriteString("INSERT IGNORE INTO transaction VALUES ")

	for idx, t := range data {
		if idx != 0 {
			sb.WriteString(", ")
		}
		sb.WriteString("(0, ?, ?, ?, ?, ?, ?, ?)")
		args = append(args, t.AtmId, t.TransactionId, t.TransactionDate, t.TransactionType, t.Amount, t.CardNum, t.DestCardNum)
	}

	query := sb.String()

	rows, err := rp.conn.Exec(query, args...)
	if err != nil {
		return 0, err
	}

	return rows, nil
}
