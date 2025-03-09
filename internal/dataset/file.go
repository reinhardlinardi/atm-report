package dataset

import (
	"errors"
	"path/filepath"
	"strings"
	"time"
)

type FileInfo struct {
	Name  string
	AtmId string
	Date  time.Time
	Ext   string
}

func ParseFileInfo(file string) (*FileInfo, error) {
	ext := filepath.Ext(file)
	name := filepath.Base(file)

	parts := strings.Split(strings.TrimSuffix(name, ext), "_")
	ext = ext[1:]

	if !isValidExt(ext) {
		return nil, errors.New("invalid ext")
	}
	if len(parts) != 2 {
		return nil, errors.New("invalid name format")
	}

	atmId := parts[0]
	dateStr := parts[1]

	date, err := time.Parse(FILENAME_DATE, dateStr)
	if err != nil {
		return nil, errors.New("invalid date format")
	}

	info := &FileInfo{Name: name, AtmId: atmId, Date: date, Ext: ext}
	return info, nil
}

func isValidExt(ext string) bool {
	return ext == FILE_CSV || ext == FILE_JSON || ext == FILE_YAML || ext == FILE_XML
}
