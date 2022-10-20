package adif

import (
	"log"
	"strconv"
	"strings"
)

// ParseField parse a field (all header or a qso row)
func ParseField(s string) map[string]Data {
	items := strings.Split(s, "<")
	ret := make(map[string]Data)

	for _, item := range items {
		if !strings.Contains(item, ":") {
			continue
		}
		index := strings.Index(item, ">")
		if index == -1 {
			log.Printf("No end of > for data specifier: '%s'\n", item)
			continue
		}
		dataSpecifierRow := substr(item, 0, index)

		dataSpecifierSplitted := strings.Split(dataSpecifierRow, ":")
		if len(dataSpecifierSplitted) < 2 {
			log.Printf("Too low element in date specified: %s\n", dataSpecifierRow)
			continue
		}
		fDataSpecified := dataSpecifierSplitted[0]

		lDataSpecified := 0
		v := dataSpecifierSplitted[1]
		if atoi, err := strconv.Atoi(v); err == nil {
			lDataSpecified = atoi
		}
		tDataSpecified := ""
		if len(dataSpecifierSplitted) >= 3 {
			tDataSpecified = dataSpecifierSplitted[2]
		}

		payload := substr(item, index+1, lDataSpecified)

		ret[strings.ToLower(fDataSpecified)] = Data{
			FieldName: fDataSpecified,
			DataType:  tDataSpecified,
			Length:    lDataSpecified,
			Data:      payload,
		}

	}

	return ret
}
func substr(s string, offset, length int) string {
	l := len(s)
	if l <= offset {
		return ""
	}
	if l < (offset + length) {
		return s[offset:]
	}
	return s[offset : offset+length]

}
func substrUTF8(s string, start, end int) string {
	counter, startIdx := 0, 0
	for i := range s {
		if counter == start {
			startIdx = i
		}
		if counter == end {
			return s[startIdx:i]
		}
		counter++
	}
	return s[startIdx:]
}
