package atmrepository

import "github.com/reinhardlinardi/atm-report/pkg/db"

type RepositoryImpl struct {
	conn db.DB
}

func New(conn db.DB) *RepositoryImpl {
	return &RepositoryImpl{conn: conn}
}

func (rp *RepositoryImpl) IsExist(atmId string) (bool, error) {
	var exist bool
	query := "SELECT EXISTS(SELECT * FROM atm WHERE atm_id = ?)"

	if err := rp.conn.QueryRow(&exist, query, atmId); err != nil {
		return false, err
	}
	return exist, nil
}
