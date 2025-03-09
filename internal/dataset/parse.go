package dataset

import (
	"encoding/json"
	"encoding/xml"
	"errors"

	"github.com/gocarina/gocsv"
	"github.com/reinhardlinardi/atm-report/internal/datestr"
	"gopkg.in/yaml.v3"
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
	list := []Transaction{}
	err := yaml.Unmarshal(raw, &list)

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

func parseXml(raw []byte) ([]Transaction, error) {
	doc := XmlRoot{}
	err := xml.Unmarshal(raw, &doc)

	if err != nil {
		return nil, err
	}

	list := doc.Data
	for _, t := range list {
		if _, valid := datestr.Parse(t.Date); !valid {
			return nil, errors.New("invalid date")
		}
	}
	return list, nil
}
