package nfredis

import (
	"jfcsrv/nfconst"
	"testing"
)

func TestExists(t *testing.T) {
	nfconst.InitialConst()
	ret, err := exists("test")
	t.Log(ret, err)
}

func TestHSet(t *testing.T) {
	nfconst.InitialConst()
	ret, err := hset("1233", "3434", "4444")
	t.Log(ret, err)
}

func TestHGet(t *testing.T) {
	nfconst.InitialConst()
	ret, err := hget("1233", "3434")
	t.Log(ret, err)
}

func TestHGetAll(t *testing.T) {
	nfconst.InitialConst()
	hgetall("plate_temp:serial:sz1234567890:bid:1:nid:1")
}
