package nflogic

import "testing"

func T_estParamLogic(t *testing.T) {

	// 请求参数的数据结构
	src := []byte{0xF5, 0xA6, 0x0A, 0x00, 0x0D, 0x73, 0x7A, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x30, 0x01, 0x00, 0xBE, 0xEF}

	// 请求参数后需要发出去的数据
	// F5 A6 0A 00 32 73 7A 31 32 33 34 35 36 37 38 39 30 01 02 01 14 02 01 28 01 40 01 40 02 E4 01 68 02 50 01 40 01 40 02 E4 01 68 02 46 01 01 32 01 2C 01 2C 02 E4 01 68 00 BE EF
	OnMsessage(src)

}

func TestParamT(t *testing.T) {
	handleAngleParamQuest("sz1234567890")
}
