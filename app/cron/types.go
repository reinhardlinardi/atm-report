package cron

import (
	"encoding/json"
	"encoding/xml"

	"github.com/gocarina/gocsv"
	"gopkg.in/yaml.v3"
)

type parseFunc = func([]byte, any) error

var parser = map[string]parseFunc{
	"csv":  gocsv.UnmarshalBytes,
	"json": json.Unmarshal,
	"yaml": yaml.Unmarshal,
	"xml":  xml.Unmarshal,
}
