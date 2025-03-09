package dataset

import "time"

type Transaction struct {
	Id          string    `json:"transactionId"`
	Date        time.Time `json:"transactionDate"`
	Type        int       `json:"transactionType"`
	Amount      int       `json:"amount"`
	CardNum     string    `json:"cardNumber"`
	DestCardNum string    `json:"destinationCardNumber"`
}
