package atmrepo

import (
	"fmt"

	"github.com/reinhardlinardi/atm-report/pkg/db"
)

const table = "atm"

type RepositoryImpl struct {
	conn db.DB
}

func New(conn db.DB) *RepositoryImpl {
	return &RepositoryImpl{conn: conn}
}

func (rp *RepositoryImpl) IsExist(atmId string) (bool, error) {
	var exist bool

	query := fmt.Sprintf("SELECT EXISTS(SELECT * FROM %s WHERE atm_id = ?)", table)
	err := rp.conn.QueryRow(&exist, query, atmId)

	if err != nil {
		return false, err
	}
	return exist, nil
}
