package datestr

import "time"

func Parse(str string) (time.Time, bool) {
	date, err := time.Parse(compactLayout, str)
	if err == nil {
		return date, true
	}
	date, err = time.Parse(dashedLayout, str)
	if err == nil {
		return date, true
	}
	return time.Time{}, false
}

func Format(date time.Time) string {
	return date.Format(compactLayout)
}
