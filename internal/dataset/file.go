package dataset

import (
	"errors"
	"path/filepath"
	"strings"
	"time"
)

type FileInfo struct {
	AtmId string
	Date  time.Time
	Ext   string
}

func ParseFileInfo(file string) (*FileInfo, error) {
	ext := filepath.Ext(file)
	name := strings.TrimSuffix(filepath.Base(file), ext)

	ext = ext[1:]
	parts := strings.Split(name, "_")

	if len(parts) != 2 {
		return nil, errors.New("invalid name format")
	}

	atmId := parts[0]
	dateStr := parts[1]

	date, err := time.Parse(FILENAME_DATE, dateStr)
	if err != nil {
		return nil, errors.New("invalid date format")
	}

	return &FileInfo{AtmId: atmId, Date: date, Ext: ext}, nil
}
