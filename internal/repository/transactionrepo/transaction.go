package transactionrepo

import (
	"fmt"
	"strings"

	"github.com/reinhardlinardi/atm-report/pkg/db"
)

const table = "transaction"

type RepositoryImpl struct {
	conn db.DB
}

func New(conn db.DB) *RepositoryImpl {
	return &RepositoryImpl{conn: conn}
}

func (rp *RepositoryImpl) InsertRows(data []Transaction) (int64, error) {
	if len(data) == 0 {
		return 0, nil
	}

	args := []any{}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("INSERT IGNORE INTO %s VALUES ", table))

	for idx, t := range data {
		if idx != 0 {
			sb.WriteString(", ")
		}
		sb.WriteString("(0, ?, ?, ?, ?, ?, ?, ?)")
		args = append(args, t.AtmId, t.TransactionId, t.Date, t.Type, t.Amount, t.CardNum, t.DestCardNum)
	}

	query := sb.String()
	rows, err := rp.conn.Exec(query, args...)

	if err != nil {
		return 0, err
	}
	return rows, nil
}
