package adif

import (
	"bufio"
	"errors"
	"os"
	"sort"
	"strings"
)

func ParseFile(filename string) (Decoded, error) {
	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		return Decoded{}, err
	}

	f, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return Decoded{}, err
	}
	defer f.Close()

	sc := bufio.NewScanner(f)

	return parseScanner(sc)
}

func ParseStr(s string) (Decoded, error) {

	scanner := bufio.NewScanner(strings.NewReader(s))

	return parseScanner(scanner)
}

func parseScanner(scanner *bufio.Scanner) (Decoded, error) {
	d := Decoded{
		Headers:       make(map[string]Data),
		HeadersFields: make([]string, 0),
		Rows:          make([]map[string]Data, 0),
		RowsFields:    make([]string, 0),
	}

	isHeader := false

	canProcess := false
	row := ""

	headerFields := make(map[string]struct{})
	rowsFields := make(map[string]struct{})

	for scanner.Scan() {
		t := scanner.Text()
		tLower := strings.ToLower(t)
		row += t
		if strings.Contains(tLower, "<eoh>") || strings.Contains(tLower, "<eor>") {
			canProcess = true

		}
		if strings.Contains(tLower, "<eoh>") {
			isHeader = true
		}

		if canProcess {
			qsoRow := ParseField(row)
			if isHeader {
				d.Headers = qsoRow
				for k := range qsoRow {
					headerFields[k] = struct{}{}
				}
			} else {
				d.Rows = append(d.Rows, qsoRow)
				for k := range qsoRow {
					rowsFields[k] = struct{}{}
				}
			}

			isHeader = false
			canProcess = false
			row = ""
		}
	}

	for k := range headerFields {
		d.HeadersFields = append(d.HeadersFields, k)
	}
	if len(d.HeadersFields) > 1 {
		sort.Strings(d.HeadersFields)
	}
	for k := range rowsFields {
		d.RowsFields = append(d.RowsFields, k)
	}
	if len(d.RowsFields) > 1 {
		sort.Strings(d.RowsFields)
	}
	err := scanner.Err()

	return d, err
}
