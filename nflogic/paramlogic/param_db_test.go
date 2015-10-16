package paramlogic

import (
	"fmt"
	"jfcsrv/nfconst"
	"testing"
)

func TestParamDb(t *testing.T) {
	nfconst.InitialConst()
	params, err := queryAngleParamByDeviceSerial("sz1234567890")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(params)
	}
}
