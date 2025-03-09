package dataset

type parser = func([]byte) ([]Transaction, error)

type Transaction struct {
	Id          string `csv:"transactionId"`
	Date        string `csv:"transactionDate"`
	Type        int    `csv:"transactionType"`
	Amount      int    `csv:"amount"`
	CardNum     string `csv:"cardNumber"`
	DestCardNum string `csv:"destinationCardNumber"`
}
