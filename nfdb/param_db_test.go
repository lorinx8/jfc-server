package nfdb

import (
	"fmt"
	"testing"
)

func TestParamDb(t *testing.T) {

	params, err := QueryAngleParamByDeviceSerial("sz1234567890")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(params)
	}

}
