package paramlogic

import (
	"bytes"
	"encoding/binary"

	"jfcsrv/nfconst"
	"jfcsrv/nflog"
)

// 参数请求消息,数据包负载结构体
type paramRequestPackage struct {
	Serial    [nfconst.LEN_DEVICE_SERIAL]byte
	ParamType byte
}

type neckAngle struct {
	nid    byte   // 角度编号
	angle  byte   // 角度
	crop_x uint16 // 剪裁区域参数x
	crop_y uint16 // 剪裁局域参数y
	crop_w uint16 // 剪裁区域参数w
	crop_h uint16 // 剪裁区域参数h
}

type baseAngle struct {
	bid   byte
	angle byte
	necks []neckAngle
}

type ParamLogic struct {
}

var jlog = nflog.Logger

func (logic *ParamLogic) OnLogicMessage(msg []byte) (cmd byte, ret []byte, err error) {

	var t paramRequestPackage
	reader := bytes.NewReader(msg)
	binary.Read(reader, binary.BigEndian, &t)
	var s string = string(t.Serial[0:len(t.Serial)])

	// 业务处理
	jlog.Info("Device Serial:", s, "request param with type", t.ParamType)

	switch t.ParamType {
	case nfconst.CMD_PARAM_TYPE_ANGLE:
		ret, err = handleAngleParamQuest(s, nfconst.CMD_PARAM_TYPE_ANGLE)
	}

	cmd = nfconst.CMD_REQUEST_PARAM_RESPONSE
	return
}

func handleAngleParamQuest(serial string, ptype byte) (ret []byte, err error) {

	retDb, err := queryAngleParamByDeviceSerial(serial)
	if err != nil {
		return nil, err
	}

	var retMap map[int]*baseAngle = make(map[int]*baseAngle)
	for _, v := range retDb {

		_bid := v.Bid
		var _nangle neckAngle
		_nangle = neckAngle{}
		_nangle.nid = byte(v.Nid)
		_nangle.angle = byte(v.Nangle)
		_nangle.crop_x = uint16(v.CropX)
		_nangle.crop_y = uint16(v.CropY)
		_nangle.crop_w = uint16(v.CropW)
		_nangle.crop_h = uint16(v.CropH)

		_bangle, ok := retMap[_bid]

		if ok {
			_bangle.necks = append(_bangle.necks, _nangle)
			// fmt.Println(_bangle)
		} else {
			var __bangle *baseAngle = &baseAngle{
				bid:   byte(v.Bid),
				angle: byte(v.Bangle),
				necks: make([]neckAngle, 0),
			}
			__bangle.necks = append(__bangle.necks, _nangle)
			retMap[_bid] = __bangle
		}
	}

	// convert map to bytes
	buf := new(bytes.Buffer)

	binary.Write(buf, binary.BigEndian, []byte(serial))
	binary.Write(buf, binary.BigEndian, ptype)

	// base angle count, uint8

	binary.Write(buf, binary.BigEndian, uint8(len(retMap))) // 角度2个数

	for k, v := range retMap {
		_bid := k // 角度2编号

		_bst := *v
		_bangle := _bst.angle      // 角度2数值
		_ncount := len(_bst.necks) // 角度1个数

		binary.Write(buf, binary.BigEndian, byte(_bid))
		binary.Write(buf, binary.BigEndian, byte(_bangle))
		binary.Write(buf, binary.BigEndian, byte(_ncount))

		for _, vv := range _bst.necks {
			binary.Write(buf, binary.BigEndian, vv.nid)
			binary.Write(buf, binary.BigEndian, vv.angle)
			binary.Write(buf, binary.BigEndian, vv.crop_x)
			binary.Write(buf, binary.BigEndian, vv.crop_y)
			binary.Write(buf, binary.BigEndian, vv.crop_w)
			binary.Write(buf, binary.BigEndian, vv.crop_h)
		}
	}
	return buf.Bytes(), nil
}
