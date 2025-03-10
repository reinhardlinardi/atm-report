package history

const table = "history"

type History struct {
	Id    int64  `db:"id" json:"id"`
	AtmId string `db:"atm_id" json:"atm_id"`
	Date  string `db:"date" json:"date"`
}
