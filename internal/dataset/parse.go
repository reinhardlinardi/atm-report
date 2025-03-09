package dataset

import (
	"encoding/json"
	"errors"

	"github.com/gocarina/gocsv"
	"github.com/reinhardlinardi/atm-report/internal/datestr"
)

var parserFunc = map[string]parser{
	"csv":  parseCsv,
	"json": parseJson,
	"yaml": parseYaml,
	"xml":  parseXml,
}

func Parse(raw []byte, ext string) ([]Transaction, error) {
	parser, valid := parserFunc[ext]
	if !valid {
		return nil, nil
	}
	return parser(raw)
}

func parseCsv(raw []byte) ([]Transaction, error) {
	list := []Transaction{}
	err := gocsv.UnmarshalBytes(raw, &list)

	if err != nil {
		return nil, err
	}
	for _, t := range list {
		if _, valid := datestr.Parse(t.Date); !valid {
			return nil, errors.New("invalid date")
		}
	}
	return list, nil
}

func parseJson(raw []byte) ([]Transaction, error) {
	list := []Transaction{}
	err := json.Unmarshal(raw, &list)

	if err != nil {
		return nil, err
	}
	for _, t := range list {
		if _, valid := datestr.Parse(t.Date); !valid {
			return nil, errors.New("invalid date")
		}
	}
	return list, nil
}

func parseYaml(raw []byte) ([]Transaction, error) {
	// arr := []Transaction{}
	// err := yaml.Unmarshal(raw, &arr)

	// if err != nil {
	// 	return nil, err
	// }
	// for _, t := range arr {
	// 	if _, valid := datestr.Parse(t.Date); !valid {
	// 		return nil, errors.New("invalid date")
	// 	}
	// }
	// return arr, nil
	return nil, nil
}

func parseXml(raw []byte) ([]Transaction, error) {
	return nil, nil
}
