package adif

import (
	"fmt"
	"testing"
)

func TestDate(t *testing.T) {
	dd := "20221223"
	tt := "1200"
	date, err := Date(dd, tt)
	fmt.Println(err)
	fmt.Println(date)
}
