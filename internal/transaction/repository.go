package transaction

import (
	"fmt"
	"strings"

	"github.com/reinhardlinardi/atm-report/pkg/db"
)

type RepositoryImpl struct {
	conn db.DB
}

func New(conn db.DB) *RepositoryImpl {
	return &RepositoryImpl{conn: conn}
}

func (r *RepositoryImpl) Load(data []Transaction) (int64, error) {
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
	rows, err := r.conn.Exec(query, args...)

	if err != nil {
		return 0, err
	}
	return rows, nil
}

func (r *RepositoryImpl) CountDaily() ([]DailyCount, error) {
	res := []DailyCount{}
	query := fmt.Sprintf("SELECT date, COUNT(*) as count FROM %s GROUP BY date", table)

	if err := r.conn.Query(&res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (r *RepositoryImpl) CountDailyByType() ([]DailyTypeCount, error) {
	res := []DailyTypeCount{}
	query := fmt.Sprintf("SELECT date, type, COUNT(*) as count FROM %s GROUP BY date, type", table)

	if err := r.conn.Query(&res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (r *RepositoryImpl) GetDailyMaxWithdraw() ([]DailyMaxWithdraw, error) {
	res := []DailyMaxWithdraw{}

	query := `SELECT t1.date, t1.atm_id, t1.amount FROM %s t1 JOIN 
		(SELECT date, MAX(amount) AS max_amount FROM %s WHERE type = 0 GROUP BY date) t2
		ON t1.date = t2.date WHERE t1.amount = t2.max_amount`

	query = fmt.Sprintf(query, table, table)

	if err := r.conn.Query(&res, query); err != nil {
		return nil, err
	}
	return res, nil
}
