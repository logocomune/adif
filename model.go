package adif

import (
	"errors"
	"time"
)

// http://www.adif.org.uk/313/ADIF_313.htm

type Decoded struct {
	Headers       map[string]Data
	HeadersFields []string
	Rows          []map[string]Data
	RowsFields    []string
}

type Data struct {
	FieldName string
	DataType  string
	Length    int
	Data      string
}

func Date(d, t string) (time.Time, error) {
	const dateFormat = "20060102"
	const time4Digit = "1504"
	const time6Digit = time4Digit + "05"

	if len(d) != 8 {
		return time.Time{}, errors.New("bad date format")
	}

	parseLayout := dateFormat

	switch len(t) {
	case 4:
		parseLayout += " " + time4Digit
	case 6:
		parseLayout += " " + time6Digit
	default:
		return time.Time{}, errors.New("bad time format")

	}
	parseLayout += " Z"
	return time.Parse(parseLayout, d+" "+t+" Z")
}
