package main

import (
	"fmt"
	"github.com/logocomune/adif"
	"log"
)

func main() {

	parsed, err := adif.ParseFile("example.adi")
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	fmt.Printf("%+v\n", parsed)

	//Output: {Headers:map[programid:{FieldName:programid DataType: Length:7 Data:MonoLog} userdef1:{FieldName:USERDEF1 DataType:N Length:8 Data:QRP_ARCI} userdef2:{FieldName:USERDEF2 DataType:E Length:19 Data:SweaterSize,{S,M,L}} userdef3:{FieldName:USERDEF3 DataType: Length:15 Data:ShoeSize,{5:20}}] HeadersFields:[programid userdef1 userdef2 userdef3] Rows:[map[band:{FieldName:band DataType: Length:3 Data:17M} call:{FieldName:call DataType: Length:6 Data:AAAAAA} mode:{FieldName:mode DataType: Length:4 Data:RTTY} qso_date:{FieldName:qso_date DataType: Length:8 Data:19960513} time_on:{FieldName:time_on DataType: Length:4 Data:1305}] map[band:{FieldName:band DataType: Length:3 Data:20M} call:{FieldName:call DataType: Length:6 Data:BBBBBB} mode:{FieldName:mode DataType: Length:3 Data:FT8} qso_date:{FieldName:qso_date DataType: Length:8 Data:19960513} time_on:{FieldName:time_on DataType: Length:4 Data:1305}]] RowsFields:[band mode qso_date time_on call]}
}
