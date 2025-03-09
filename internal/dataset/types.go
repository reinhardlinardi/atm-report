package dataset

import "time"

type Transaction struct {
	Id          string    `csv:"transactionId"`
	Date        time.Time `csv:"transactionDate"`
	Type        int       `csv:"transactionType"`
	Amount      int       `csv:"amount"`
	CardNum     string    `csv:"cardNumber"`
	DestCardNum string    `csv:"destinationCardNumber"`
}
