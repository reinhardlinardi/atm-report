package cron

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gocarina/gocsv"
	"github.com/reinhardlinardi/atm-report/internal/transaction"
	"gopkg.in/yaml.v3"
)

var parser = map[string]parseFunc{
	"csv":  gocsv.UnmarshalBytes,
	"json": json.Unmarshal,
	"yaml": yaml.Unmarshal,
	"xml":  xml.Unmarshal,
}

func (c *Cron) handleFile(path string) (bool, error) {
	// Assume filename format valid
	filename := filepath.Base(path)
	ext := filepath.Ext(path)[1:]
	name := strings.Split(filename, ".")[0]

	s := strings.Split(name, "_")

	atmId := s[0]
	date := s[1]
	seq, _ := strconv.Atoi(s[2])

	// File processing
	skip, err := c.skipFile(atmId, date, seq)
	if err != nil {
		return false, fmt.Errorf("err check skip: %s: %s", err.Error(), filename)
	}
	if skip {
		return false, nil
	}

	if err := c.loadFile(path, ext, atmId, date, seq); err != nil {
		return false, fmt.Errorf("err load file: %s: %s", err.Error(), filename)
	}
	return true, nil
}

func (c *Cron) skipFile(atmId, date string, seq int) (bool, error) {
	skip, err := c.history.Check(atmId, date, seq)
	if err != nil {
		return true, errors.New("err check history")
	}

	return skip, nil
}

func (c *Cron) loadFile(path, ext, atmId, date string, seq int) error {
	bytes, err := c.fileStorage.Get(path)
	if err != nil {
		return fmt.Errorf("err read file: %s", err.Error())
	}

	data, err := parseFile(bytes, atmId, ext)
	if err != nil {
		return fmt.Errorf("err parse file: %s", err.Error())
	}

	_, err = c.transaction.Load(data)
	if err != nil {
		return fmt.Errorf("err load data: %s", err.Error())
	}

	_, err = c.history.Append(atmId, date, seq)
	if err != nil {
		return fmt.Errorf("err append history: %s", err.Error())
	}

	return nil
}

func parseFile(bytes []byte, atmId, ext string) ([]transaction.Transaction, error) {
	var err error
	data := []transaction.Transaction{}

	if ext == "xml" {
		doc := transaction.Transactions{}
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
