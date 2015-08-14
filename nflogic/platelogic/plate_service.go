package platelogic

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"jfcsrv/nfconst"
	"jfcsrv/nfutil"
	"strconv"
	"strings"
)

// 结果上报,数据负载结构体
type plateResultRequestPackage struct {
	Serial      [nfconst.LEN_DEVICE_SERIAL]byte
	Bid         byte
	Nid         byte
	Count       byte
	PlateNumber [nfconst.LEN_MAX_PLATE_NUMBER]byte
}

type PlateResultLogic struct {
}

func (logic *PlateResultLogic) OnLogicMessage(msg []byte) (cmd byte, ret []byte, err error) {
	_cmd := nfconst.CMD_REQUEST_ONE_ANGLE_RESULT_RESPONSE
	msgChan := make(chan []byte)
	go handlePlateReport(msgChan)
	msgChan <- msg[:]

	return _cmd, nil, nil
}

// 处理车牌结果
func handlePlateReport(msgChan chan []byte) {
	// 解析角度结果数据包
	msg := <-msgChan
	var t plateResultRequestPackage
	n := binary.Size(t)
	buffer := bytes.NewBuffer(msg[0:n])
	binary.Read(buffer, binary.BigEndian, &t)

	// 设备编号
	serial := string(t.Serial[0:len(t.Serial)])

	// 图片数据
	n_img := len(msg) - n
	img_data := msg[n:]

	// 业务处理
	fmt.Printf("Device Serial:%s, report angle result: bid=%d, nid=%d, count=%d, img_n=%d\n", t.Serial, t.Bid, t.Nid, t.Count, n_img)

	// 处理上报数据无车牌的情况
	if t.Count == 0 {
		handlePlateInPool(serial, t.Bid, t.Nid, "")
		return
	}

	// 如果上报的车牌数量不为1
	if t.Count != 0 {
		orgNumber := string(t.PlateNumber[:])
		pcode, pchar, cchar, pnumber, _err := parsePlateNumber(orgNumber)
		if _err == nil {
			// 能够得到像样的车牌号码，进行下一步处理
			fmt.Println(pcode, pchar, cchar, pnumber)
			handlePlateInPool(serial, t.Bid, t.Nid, pnumber)
			// 保存图片
			savePlateImage(serial, t.Bid, t.Nid, img_data)

		} else {
			// 这次上报的数据不要了
			fmt.Println(_err)
			return
		}
	}
}

// 处理车牌识别结果, 判断是否用本次的上报数据
func handlePlateInPool(serial string, bid byte, nid byte, platenumber string) {
}

// 比较新旧车牌号码的相似度
func calSimilarity(plate_new string, plate_old string) (ret int) {
	if plate_old == "" && plate_new != "" {
		return 0
	}
	if plate_new == "" && plate_old != "" {
		return 0
	}

	oldByte := []byte(plate_old)
	newByte := []byte(plate_new)

	var max int
	if len(oldByte) > len(newByte) {
		max = len(oldByte)
	} else {
		max = len(newByte)
	}
	step := 20
	for i := 0; i < max; i++ {
		if oldByte[i] == newByte[i] {
			ret = ret + step
		}
	}
	return
}

// 从原始的车牌号码中解析数据
// 从JFC过来的格式为  7_B_12345 [粤B12345]
// 第一个7为省份号码，B为城市编号，后五位车牌号码
// pcode 省份号码
// pchar 城市字面值，粤
// cchar 城市编号AB
// pnumber 五位号码
func parsePlateNumber(orgNumber string) (pcode int, pchar string, cchar string, pnumber string, err error) {
	ss := strings.Split(orgNumber, "_")
	if len(ss) != 3 {
		err = errors.New("org plate invalid")
		return
	}
	pcode, err = strconv.Atoi(ss[0])
	if err != nil {
		return
	}
	pchar = nfconst.PPCharMap[pcode]
	cchar = ss[1]
	org := ss[2]
	pnumber = string(([]byte(org))[0:5])
	return
}

// 处理车牌图片
func savePlateImage(serial string, bid byte, nid byte, imgdata []byte) (err error) {
	// 生成图片路径
	name, _err := generateOrgFilePath(serial, bid, nid)
	if _err != nil {
		return _err
	}
	// 写文件
	_, _err2 := nfutil.WriteFile(name, imgdata)
	if _err2 != nil {
		return _err2
	}
	return nil
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
