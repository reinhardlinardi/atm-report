package atm

const table = "atm"

type Atm struct {
	Id    int64  `db:"id" json:"id"`
	AtmId string `db:"atm_id" json:"atm_id"`
}
