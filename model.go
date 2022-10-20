package adif

import (
	"errors"
	"time"
)


type Decoded struct {
	Headers       map[string]Data   `json:"headers"`
	HeadersFields []string          `json:"headersFields"`
	Rows          []map[string]Data `json:"rows"`
	RowsFields    []string          `json:"rowsFields"`
}

type Data struct {
	FieldName string `json:"fieldName"`
	DataType  string `json:"dataType"`
	Length    int    `json:"length"`
	Data      string `json:"data"`
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
