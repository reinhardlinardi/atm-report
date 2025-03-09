package atmrepo

type Atm struct {
	Id    int64  `db:"id"`
	AtmId string `db:"atm_id"`
}
