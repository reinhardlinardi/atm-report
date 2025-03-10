package app

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/gocarina/gocsv"
	"github.com/reinhardlinardi/atm-report/internal/repository/transactionrepo"
	"gopkg.in/yaml.v3"
)

func (c *Cron) handleFile(path string) error {
	// Assume filename format valid
	filename := filepath.Base(path)
	ext := filepath.Ext(path)[1:]

	name := strings.Split(filename, ".")[0]
	parts := strings.Split(name, "_")
	atmId := parts[0]
	date := parts[1]

	skip, err := c.skipFile(atmId, date)
	if err != nil {
		return fmt.Errorf("err check skip file: %s: %s", err.Error(), filename)
	}

	if !skip {
		if err := c.loadFile(path, atmId, date, ext); err != nil {
			return fmt.Errorf("err load file: %s: %s", err.Error(), filename)
		}
	}
	return nil
}

func (c *Cron) skipFile(atmId, date string) (bool, error) {
	exist, err := c.atmRepo.IsExist(atmId)
	if err != nil {
		return true, errors.New("err check atm id")
	}
	if !exist {
		return true, errors.New("atm id not exist")
	}

	skip, err := c.historyRepo.IsExist(atmId, date)
	if err != nil {
		return true, errors.New("err check load history")
	}

	return skip, nil
}

func (c *Cron) loadFile(path, atmId, date, ext string) error {
	bytes, err := c.storage.Fetch(path)
	if err != nil {
		return fmt.Errorf("err fetch file: %s", err.Error())
	}

	data, err := parseFile(bytes, atmId, ext)
	if err != nil {
		return fmt.Errorf("err parse file: %s", err.Error())
	}

	_, err = c.transactionRepo.InsertRows(data)
	if err != nil {
		return fmt.Errorf("err insert data: %s", err.Error())
	}

	_, err = c.historyRepo.Insert(atmId, date)
	if err != nil {
		return fmt.Errorf("err insert load history: %s", err.Error())
	}

	return nil
}

func parseFile(bytes []byte, atmId, ext string) ([]transactionrepo.Transaction, error) {
	data := []Transaction{}
	res := []transactionrepo.Transaction{}

	var doc XmlRoot
	var err error

	switch ext {
	case "csv":
		err = gocsv.UnmarshalBytes(bytes, &data)
	case "json":
		err = json.Unmarshal(bytes, &data)
	case "yaml":
		err = yaml.Unmarshal(bytes, &data)
	case "xml":
		err = xml.Unmarshal(bytes, &doc)
	}

	if err != nil {
		return nil, err
	}
	if ext == "xml" {
		data = doc.Data
	}

	for _, item := range data {
		t := transactionrepo.Transaction{}

		t.AtmId = atmId
		t.TransactionId = item.Id
		t.Date = item.Date
		t.Type = item.Type
		t.Amount = item.Amount
		t.CardNum = item.CardNum
		t.DestCardNum = item.DestCardNum

		res = append(res, t)
	}

	return res, nil
}
