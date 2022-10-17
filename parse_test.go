package adif

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const exHeaeder = `<adif_ver:5>3.0.5
<programid:7>MonoLog
<USERDEF1:8:N>QRP_ARCI
<USERDEF2:19:E>SweaterSize,{S,M,L}

<USERDEF3:15>ShoeSize,{5:20}
<EOH>
`

const exRows = `
<call:6>AAAAAA<band:3>17M<mode:4>RTTY<qso_date:8>19960513<time_on:4>1305<eor>

<call:6>BBBBBB

<band:3>20M
<mode:3>FT8
<qso_date:8>19960513
<time_on:4>1305<eor>`

func TestParseStr(t *testing.T) {
	//Empty string
	parsed, err := ParseStr(exHeaeder)
	assert.Nil(t, err)
	assert.NotEmpty(t, parsed.Headers)
	assert.Empty(t, parsed.Rows)

	//Only header
	parsed, err = ParseStr(exHeaeder)
	assert.Nil(t, err)
	assert.NotEmpty(t, parsed.Headers)
	assert.EqualValues(t, len(parsed.Headers), 5)
	assert.Empty(t, parsed.Rows)

	//Only body
	parsed, err = ParseStr(exRows)
	assert.Nil(t, err)
	assert.Empty(t, parsed.Headers)
	assert.NotEmpty(t, parsed.Rows)
	assert.EqualValues(t, len(parsed.Rows), 2)
	assert.EqualValues(t, len(parsed.Rows[0]), 5)
	assert.EqualValues(t, len(parsed.Rows[1]), 5)
	assert.EqualValues(t, parsed.Rows[0]["call"], Data{
		FieldName: "call",
		DataType:  "",
		Length:    6,
		Data:      "AAAAAA",
	})
	assert.EqualValues(t, parsed.Rows[0]["band"], Data{
		FieldName: "band",
		DataType:  "",
		Length:    3,
		Data:      "17M",
	})
	assert.EqualValues(t, parsed.Rows[0]["mode"], Data{
		FieldName: "mode",
		DataType:  "",
		Length:    4,
		Data:      "RTTY",
	})

	assert.EqualValues(t, parsed.Rows[0]["qso_date"], Data{
		FieldName: "qso_date",
		DataType:  "",
		Length:    8,
		Data:      "19960513",
	})
	assert.EqualValues(t, parsed.Rows[0]["time_on"], Data{
		FieldName: "time_on",
		DataType:  "",
		Length:    4,
		Data:      "1305",
	})
	assert.EqualValues(t, parsed.Rows[1]["call"], Data{
		FieldName: "call",
		DataType:  "",
		Length:    6,
		Data:      "BBBBBB",
	})
	assert.EqualValues(t, parsed.Rows[1]["band"], Data{
		FieldName: "band",
		DataType:  "",
		Length:    3,
		Data:      "20M",
	})
	assert.EqualValues(t, parsed.Rows[1]["mode"], Data{
		FieldName: "mode",
		DataType:  "",
		Length:    3,
		Data:      "FT8",
	})

	assert.EqualValues(t, parsed.Rows[1]["qso_date"], Data{
		FieldName: "qso_date",
		DataType:  "",
		Length:    8,
		Data:      "19960513",
	})
	assert.EqualValues(t, parsed.Rows[1]["time_on"], Data{
		FieldName: "time_on",
		DataType:  "",
		Length:    4,
		Data:      "1305",
	})

	//Headers and Rows

	//Only body
	parsed, err = ParseStr(exHeaeder + exRows)
	assert.Nil(t, err)
	assert.NotEmpty(t, parsed.Headers)
	assert.NotEmpty(t, parsed.Rows)
	assert.EqualValues(t, len(parsed.Rows), 2)
	assert.EqualValues(t, len(parsed.Rows[0]), 5)
	assert.EqualValues(t, len(parsed.Rows[1]), 5)
	assert.EqualValues(t, parsed.Rows[0]["call"], Data{
		FieldName: "call",
		DataType:  "",
		Length:    6,
		Data:      "AAAAAA",
	})
	assert.EqualValues(t, parsed.Rows[0]["band"], Data{
		FieldName: "band",
		DataType:  "",
		Length:    3,
		Data:      "17M",
	})
	assert.EqualValues(t, parsed.Rows[0]["mode"], Data{
		FieldName: "mode",
		DataType:  "",
		Length:    4,
		Data:      "RTTY",
	})

	assert.EqualValues(t, parsed.Rows[0]["qso_date"], Data{
		FieldName: "qso_date",
		DataType:  "",
		Length:    8,
		Data:      "19960513",
	})
	assert.EqualValues(t, parsed.Rows[0]["time_on"], Data{
		FieldName: "time_on",
		DataType:  "",
		Length:    4,
		Data:      "1305",
	})
	assert.EqualValues(t, parsed.Rows[1]["call"], Data{
		FieldName: "call",
		DataType:  "",
		Length:    6,
		Data:      "BBBBBB",
	})
	assert.EqualValues(t, parsed.Rows[1]["band"], Data{
		FieldName: "band",
		DataType:  "",
		Length:    3,
		Data:      "20M",
	})
	assert.EqualValues(t, parsed.Rows[1]["mode"], Data{
		FieldName: "mode",
		DataType:  "",
		Length:    3,
		Data:      "FT8",
	})

	assert.EqualValues(t, parsed.Rows[1]["qso_date"], Data{
		FieldName: "qso_date",
		DataType:  "",
		Length:    8,
		Data:      "19960513",
	})
	assert.EqualValues(t, parsed.Rows[1]["time_on"], Data{
		FieldName: "time_on",
		DataType:  "",
		Length:    4,
		Data:      "1305",
	})
}
