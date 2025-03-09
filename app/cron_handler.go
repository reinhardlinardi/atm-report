package app

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/reinhardlinardi/atm-report/internal/dataset"
	"github.com/reinhardlinardi/atm-report/internal/datestr"
)

func (app *App) handleFile(path string) error {
	filename := filepath.Base(path)
	ext := filepath.Ext(path)[1:]

	name := strings.Split(filename, ".")[0]
	parts := strings.Split(name, "_")

	if len(parts) != 2 {
		return fmt.Errorf("invalid name format: %s", filename)
	}
	atmId := parts[0]

	date, valid := datestr.Parse(parts[1])
	if !valid {
		return fmt.Errorf("invalid date format: %s", filename)
	}
	if !isExtValid(ext) {
		return fmt.Errorf("invalid ext: %s", filename)
	}

	skip, err := app.checkSkipFile(atmId, date)
	if err != nil {
		return fmt.Errorf("err check skip file: %s: %s", err.Error(), filename)
	}
	if skip {
		return nil
	}

	if err := app.loadFile(path, atmId, ext, date); err != nil {
		return fmt.Errorf("err load file: %s: %s", err.Error(), filename)
	}
	return nil
}

func (app *App) checkSkipFile(atmId string, date time.Time) (bool, error) {
	exist, err := app.atmRepo.IsExist(atmId)
	if err != nil {
		return true, errors.New("err check atm id")
	}
	if !exist {
		return true, errors.New("atm id not exist")
	}

	skip, err := app.fileLoadRepo.IsExist(atmId, datestr.Format(date))
	if err != nil {
		return true, errors.New("err check file load history")
	}

	return skip, nil
}

func (app *App) loadFile(path, atmId, ext string, date time.Time) error {
	raw, err := app.storage.Fetch(path)
	if err != nil {
		return fmt.Errorf("err fetch file: %s", err.Error())
	}

	data, err := dataset.Parse(raw, ext)
	if err != nil {
		return fmt.Errorf("err parse file: %s", err.Error())
	}

	fmt.Println(data)
	return nil
}

func isExtValid(ext string) bool {
	return ext == "csv" || ext == "json" || ext == "yaml" || ext == "xml"
}
