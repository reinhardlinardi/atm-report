package fileloadrepository

import (
	"github.com/reinhardlinardi/atm-report/pkg/db"
)

type RepositoryImpl struct {
	conn db.DB
}

func New(conn db.DB) *RepositoryImpl {
	return &RepositoryImpl{conn: conn}
}

func (rp *RepositoryImpl) IsExist(atmId, date string) (bool, error) {
	var exist bool
	query := "SELECT EXISTS(SELECT * FROM file_load WHERE atm_id = ? AND date = ?)"

	if err := rp.conn.QueryRow(&exist, query, atmId, date); err != nil {
		return false, err
	}
	return exist, nil
}
