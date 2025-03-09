package model

type Transaction struct {
	Id              int64  `db:"id"`
	AtmId           string `db:"atm_id"`
	TransactionId   string `db:"transaction_id"`
	TransactionDate string `db:"transaction_date"`
	TransactionType int    `db:"transaction_type"`
	Amount          int    `db:"amount"`
	CardNum         string `db:"card_number"`
	DestCardNum     string `db:"destination_card_number"`
}
