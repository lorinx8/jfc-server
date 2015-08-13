package nfutil

import (
	"fmt"
	"testing"
)

func T_estWriteFile(t *testing.T) {

	name := "/Users/lorin/tt.bat"
	nn, err := WriteFile(name, []byte{0x31, 0x32})
	if err != nil && nn != 2 {
		t.Error(err)
	}
}

func TestGetNow(t *testing.T) {
	y, mon, d, h, min, s := getNow()
	fmt.Println(y, mon, d, h, min, s)
}
