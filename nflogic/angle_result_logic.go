package nflogic

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"jfcsrv/nfconst"
	"jfcsrv/nfutil"
)

// 结果上报,数据负载结构体
type angleResultRequestPackage struct {
	Serial      [nfconst.LEN_DEVICE_SERIAL]byte
	Bid         byte
	Nid         byte
	Count       byte
	PlateNumber [nfconst.LEN_MAX_PLATE_NUMBER]byte
}

type AngleResultLogic struct {
}

func (logic *AngleResultLogic) OnLogicMessage(msg []byte) (cmd byte, ret []byte, err error) {
	_cmd := nfconst.CMD_REQUEST_ONE_ANGLE_RESULT_RESPONSE
	var t angleResultRequestPackage
	n := binary.Size(t)
	buffer := bytes.NewBuffer(msg[0:n])
	binary.Read(buffer, binary.BigEndian, &t)
	n_img := len(msg) - n
	// 业务处理
	fmt.Printf("Device Serial:%s, report angle result: bid=%d, nid=%d, count=%d, img_n=%d\n", t.Serial, t.Bid, t.Nid, t.Count, n_img)

	if t.Count != 0 {
		name, _err := generateOrgFilePath(string(t.Serial[0:len(t.Serial)]), t.Bid, t.Nid)
		if _err != nil {
			return _cmd, nil, _err
		}
		fmt.Println("write file to", name)
		nn, _err2 := nfutil.WriteFile(name, msg[n:])
		if _err2 != nil {
			return _cmd, nil, _err2
		} else {
			fmt.Println("Plate number:", string(t.PlateNumber[0:len(t.PlateNumber)]), ",", nn, "bytes writed")
		}
	}

	return _cmd, nil, nil
}

// 保存的路径
// 目前暂时为相对于应用程序当前目录  picp/{设备序号}/{日期}/{文件名}
// 按照接受时间, 生成最原始的文件名, 时分秒均为2字符
// bid(2字符)_nid(2字符)_时_分_秒.jpg
func generateOrgFilePath(serial string, bid byte, nid byte) (path string, err error) {
	y, mon, d, h, min, s := nfutil.GetNow()

	_path := fmt.Sprintf("picp/%s/%04d-%02d-%02d/", serial, y, mon, d)
	_name := fmt.Sprintf("%02d-%02d_%02d-%02d-%02d%s", bid, nid, h, min, s, nfconst.FILENAME_IMG_EXTENT)
	_err := nfutil.CreateFolderIfNotExist(_path)
	if _err != nil {
		return "", _err
	}
	return _path + _name, nil
}
