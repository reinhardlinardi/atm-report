package historyrepo

type History struct {
	Id    int64  `db:"id"`
	AtmId string `db:"atm_id"`
	Date  string `db:"date"`
}
