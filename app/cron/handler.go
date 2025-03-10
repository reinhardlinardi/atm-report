package cron

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	transactiondb "github.com/reinhardlinardi/atm-report/internal/transaction"
)

func (c *Cron) handleFile(path string) error {
	// Assume filename format valid
	filename := filepath.Base(path)
	ext := filepath.Ext(path)[1:]

	name := strings.Split(filename, ".")[0]

	parts := strings.Split(name, "_")
	atmId := parts[0]
	date := parts[1]

	// File processing
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
	skip, err := c.historyDB.IsExist(atmId, date)
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

	_, err = c.transactionDB.Load(data)
	if err != nil {
		return fmt.Errorf("err insert data: %s", err.Error())
	}

	_, err = c.historyDB.Insert(atmId, date)
	if err != nil {
		return fmt.Errorf("err insert load history: %s", err.Error())
	}

	return nil
}

func parseFile(bytes []byte, atmId, ext string) ([]transactiondb.Transaction, error) {
	var err error
	data := []transactiondb.Transaction{}

	if ext == "xml" {
		doc := transactiondb.Transactions{}
		err = parser[ext](bytes, &doc)
		data = doc.List
	} else {
		err = parser[ext](bytes, &data)
	}

	if err != nil {
		return nil, err
	}
	for idx := range data {
		data[idx].AtmId = atmId
	}
	return data, nil
}
