package paramlogic

import (
	"fmt"
	"testing"
)

func TestParamDb(t *testing.T) {

	params, err := queryAngleParamByDeviceSerial("sz1234567890")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(params)
	}
}
