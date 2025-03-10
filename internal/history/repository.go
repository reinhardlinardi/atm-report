package history

import (
	"fmt"

	"github.com/reinhardlinardi/atm-report/pkg/db"
)

type RepositoryImpl struct {
	conn db.DB
}

func New(conn db.DB) *RepositoryImpl {
	return &RepositoryImpl{conn: conn}
}

func (r *RepositoryImpl) Check(atmId, date string, seq int) (bool, error) {
	var exist bool

	query := fmt.Sprintf("SELECT EXISTS(SELECT * FROM %s WHERE atm_id = ? AND date = ? AND seq = ?)", table)
	err := r.conn.QueryRow(&exist, query, atmId, date, seq)

	if err != nil {
		return false, err
	}
	return exist, nil
}

func (r *RepositoryImpl) Append(atmId, date string, seq int) (int64, error) {
	query := fmt.Sprintf("INSERT INTO %s VALUES (0, ?, ?, ?)", table)
	id, err := r.conn.InsertRow(query, atmId, date, seq)

	if err != nil {
		return 0, err
	}
	return id, nil
}
