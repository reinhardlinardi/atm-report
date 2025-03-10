package transaction

import "encoding/xml"

const table = "transaction"

type Transactions struct {
	XMLName xml.Name      `xml:"transactions"`
	List    []Transaction `xml:"transaction"`
}

type Transaction struct {
	Id            int64  `db:"id" json:"id"`
	AtmId         string `db:"atm_id" json:"atm_id"`
	TransactionId string `db:"transaction_id" csv:"transactionId" json:"transactionId" yaml:"transactionId" xml:"transactionId,attr"`
	Date          string `db:"date" csv:"transactionDate" json:"transactionDate" yaml:"transactionDate" xml:"transactionDate"`
	Type          int    `db:"type" csv:"transactionType" json:"transactionType" yaml:"transactionType" xml:"transactionType"`
	Amount        int    `db:"amount" csv:"amount" json:"amount" yaml:"amount" xml:"amount"`
	CardNum       string `db:"card_num" csv:"cardNumber" json:"cardNumber" yaml:"cardNumber" xml:"cardNumber"`
	DestCardNum   string `db:"dest_card_num" csv:"destinationCardNumber" json:"destinationCardNumber" yaml:"destinationCardNumber" xml:"destinationCardNumber"`
}

type DailyCount struct {
	Date  string `db:"date" json:"date"`
	Count int    `db:"count" json:"count"`
}

type DailyTypeCount struct {
	Date  string `db:"date" json:"date"`
	Type  int    `db:"type" json:"type"`
	Count int    `db:"count" json:"count"`
}

type DailyMaxWithdraw struct {
	Date   string `db:"date" json:"date"`
	AtmId  string `db:"atm_id" json:"atm_id"`
	Amount int    `db:"amount" json:"amount"`
}
