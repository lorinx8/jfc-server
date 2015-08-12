package nfnet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"jfcsrv/nfconst"
)

type NFPackage struct {
	Header1 byte
	Header2 byte
	Cmd     byte
	Length  uint16
	Payload []byte
	Check   byte
	Ender1  byte
	Ender2  byte
}

func NewNFRequestPackage(data []byte) (pk NFPackage, err error) {
	pk = NFPackage{}
	// heaer
	pk.Header1 = data[0]
	pk.Header2 = data[1]
	if pk.Header1 != nfconst.SOCK_PACK_HEADER_L || pk.Header2 != nfconst.SOCK_PACK_HEADER_H {
		err = errors.New("nf package data header invalid")
		return pk, err
	}

	dl := len(data)
	// ender
	pk.Ender1 = data[dl-2]
	pk.Ender2 = data[dl-1]
	if pk.Ender1 != nfconst.SOCK_PACK_ENDER_L || pk.Ender2 != nfconst.SOCK_PACK_ENDER_H {
		err = errors.New("nf package data ender invalid")
		return pk, err
	}

	// command
	pk.Cmd = data[2]

	// length
	pk.Length = binary.BigEndian.Uint16(data[3:5])
	if int(pk.Length+8) != dl {
		err = errors.New("nf package data length incorrect")
		return pk, err
	}

	// payload
	pk.Payload = data[5 : 5+pk.Length]

	// cs
	pk.Check = data[dl-3]
	return pk, nil
}

func NewNFResponseBytes(cmd byte, payload []byte) (ret []byte, err error) {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, nfconst.SOCK_PACK_HEADER_L)
	binary.Write(buf, binary.BigEndian, nfconst.SOCK_PACK_HEADER_H)
	binary.Write(buf, binary.BigEndian, cmd)
	binary.Write(buf, binary.BigEndian, uint16(len(payload)))
	binary.Write(buf, binary.BigEndian, payload)
	binary.Write(buf, binary.BigEndian, byte(0))
	binary.Write(buf, binary.BigEndian, nfconst.SOCK_PACK_ENDER_L)
	binary.Write(buf, binary.BigEndian, nfconst.SOCK_PACK_ENDER_H)
	return buf.Bytes(), nil
}
