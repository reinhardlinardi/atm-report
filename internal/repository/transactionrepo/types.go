package transactionrepo

type Transaction struct {
	Id            int64  `db:"id"`
	AtmId         string `db:"atm_id"`
	TransactionId string `db:"transaction_id"`
	Date          string `db:"date"`
	Type          int    `db:"type"`
	Amount        int    `db:"amount"`
	CardNum       string `db:"card_num"`
	DestCardNum   string `db:"dest_card_num"`
}
