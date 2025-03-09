package dataset

import "encoding/xml"

type parser = func([]byte) ([]Transaction, error)

type Transaction struct {
	Id          string `csv:"transactionId" json:"transactionId" yaml:"transactionId" xml:"transactionId,attr"`
	Date        string `csv:"transactionDate" json:"transactionDate" yaml:"transactionDate" xml:"transactionDate"`
	Type        int    `csv:"transactionType" json:"transactionType" yaml:"transactionType" xml:"transactionType"`
	Amount      int    `csv:"amount" json:"amount" yaml:"amount" xml:"amount"`
	CardNum     string `csv:"cardNumber" json:"cardNumber" yaml:"cardNumber" xml:"cardNumber"`
	DestCardNum string `csv:"destinationCardNumber" json:"destinationCardNumber" yaml:"destinationCardNumber" xml:"destinationCardNumber"`
}

type XmlDoc struct {
	XMLName xml.Name      `xml:"transactions"`
	Data    []Transaction `xml:"transaction"`
}
