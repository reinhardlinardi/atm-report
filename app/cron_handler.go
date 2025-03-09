package app

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"
	"time"
)

func (app *App) handleFile(path string) error {
	filename := getFilename(path)

	ext := filepath.Ext(path)[1:]
	name := strings.Split(filename, ".")[0]
	parts := strings.Split(name, "_")

	if len(parts) != 2 {
		return fmt.Errorf("invalid name format: %s", filename)
	}

	atmId := parts[0]
	dateStr := parts[1]

	date, err := time.Parse(dateFmt, dateStr)
	if err != nil {
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

	if err := app.loadFile(path); err != nil {
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

	skip, err := app.fileLoadRepo.IsExist(atmId, date.Format(dateFmt))
	if err != nil {
		return true, errors.New("err check file load history")
	}

	return skip, nil
}

func (app *App) loadFile(path string) error {
	raw, err := app.storage.Fetch(path)
	if err != nil {
		return fmt.Errorf("err fetch file: %s", err.Error())
	}

	fmt.Println(string(raw))

	// transcation, err := dataset.Parse()
	// filename := getFilename(path)
	return nil
}

func getFilename(path string) string {
	return filepath.Base(path)
}

func isExtValid(ext string) bool {
	return ext == "csv" || ext == "json" || ext == "yaml" || ext == "xml"
}
