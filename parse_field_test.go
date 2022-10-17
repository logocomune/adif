package adif

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Example_ParseFieldHeaders() {

	s := `ADIF Export
<adif_ver:5>3.1.1
<created_timestamp:15>20210902 162247
<programid:4>Test
<programversion:5>2.3.0
<eoh>
....`

	row := ParseField(s)
	fmt.Printf("%+v\n", row)
	//Output: map[adif_ver:{FieldName:adif_ver DataType: Length:5 Data:3.1.1} created_timestamp:{FieldName:created_timestamp DataType: Length:15 Data:20210902 162247} programid:{FieldName:programid DataType: Length:4 Data:Test} programversion:{FieldName:programversion DataType: Length:5 Data:2.3.0}]
}

func Example_ParseField() {
	s := `<call:6>AAAAAA<band:3>20M<mode:4>RTTY<qso_date:8>19960513<time_on:4>1305<eor>`

	row := ParseField(s)
	fmt.Printf("%+v\n", row)
	//Output: map[band:{FieldName:band DataType: Length:3 Data:20M} call:{FieldName:call DataType: Length:6 Data:AAAAAA} mode:{FieldName:mode DataType: Length:4 Data:RTTY} qso_date:{FieldName:qso_date DataType: Length:8 Data:19960513} time_on:{FieldName:time_on DataType: Length:4 Data:1305}]
}

func Example_ParseFieldMultiLine() {
	s := `<qso_date:8>19960513
<time_on:4>1305
<call:6>AAAAAA
<band:3>20M
<mode:4>RTTY
<eor>`
	row := ParseField(s)

	fmt.Printf("%+v\n", row)
	//Output: map[band:{FieldName:band DataType: Length:3 Data:20M} call:{FieldName:call DataType: Length:6 Data:AAAAAA} mode:{FieldName:mode DataType: Length:4 Data:RTTY} qso_date:{FieldName:qso_date DataType: Length:8 Data:19960513} time_on:{FieldName:time_on DataType: Length:4 Data:1305}]
}

func Test_substr(t *testing.T) {
	type args struct {
		s      string
		offset int
		length int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Empty",
			args: args{
				s:      "",
				offset: 1,
				length: 10,
			},
			want: "",
		},
		{
			name: "Length bigger",
			args: args{
				s:      "abcd",
				offset: 1,
				length: 10,
			},
			want: "bcd",
		},
		{
			name: "Substring",
			args: args{
				s:      "abcd",
				offset: 1,
				length: 3,
			},
			want: "bcd",
		},
		{
			name: "Offset Bigger",
			args: args{
				s:      "abcd",
				offset: 10,
				length: 3,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, substr(tt.args.s, tt.args.offset, tt.args.length), "substr(%v, %v, %v)", tt.args.s, tt.args.offset, tt.args.length)
		})
	}
}
