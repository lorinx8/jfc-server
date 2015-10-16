package paramlogic

import (
	"jfcsrv/nfconst"
	"jfcsrv/nfutil"
	"log"
	"testing"
)

func TestHandleAngleParamRequest(t *testing.T) {
	nfconst.InitialConst()
	ret, err := handleAngleParamRequest("sz1234567890", nfconst.CMD_PARAM_TYPE_ANGLE)
	if err != nil {
		log.Println(err)
	} else {
		nfutil.PrintHexArrayToStdout(ret)
	}
}
