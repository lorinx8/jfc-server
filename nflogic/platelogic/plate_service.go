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

type plateNumberInfo struct {
	ProvinceCode int
	ProvinceChar string
	CityCode     string
	PlateNo      string
	ImageByte    []byte
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
	n_img := len(msg) - n // 图片的长度
	img_data := msg[n:]

	// 业务处理
	fmt.Printf("Device Serial:%s, report angle result: bid=%d, nid=%d, count=%d, img_n=%d\n", t.Serial, t.Bid, t.Nid, t.Count, n_img)

	// 处理上报数据无车牌的情况
	if t.Count == 0 {
		handlePlateInPool(serial, int(t.Bid), int(t.Nid), nil)
		return
	}

	// 上报的车牌不等于
	orgString := string(t.PlateNumber[:])
	pinfo, _err := parseOrgPlateString(orgString)
	if _err != nil {
		// 车牌都解析不成功，不接续了
		return
	}
	pinfo.ImageByte = img_data // 好了，现在车牌信息全了
	handlePlateInPool(serial, int(t.Bid), int(t.Nid), &pinfo)
	// 。。。
}

// 处理车牌识别结果, 判断是否用本次的上报数据
func handlePlateInPool(serial string, bid int, nid int, pinfo *plateNumberInfo) error {
	// 从缓存中取得已有的车牌缓存结果
	plateTemp, err := getPlateTemp(serial, bid, nid)
	if err != nil {
		return err // 不再处理这个了
	}

	// 说明这个设备的这个位置，还没有数据
	if plateTemp == nil {
		// 处理第一个数据， 直接接受这个车牌

	}

	return nil
}

// 采用这个车牌，需要做这么几件事情
// 更新缓存数据，写入数据库，向云存储中上传数据
func acceptPlateNumber(serial string, bid byte, nid byte, pinfo *plateNumberInfo) {

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
func parseOrgPlateString(orgString string) (pinfo plateNumberInfo, err error) {
	ss := strings.Split(orgString, "_")
	if len(ss) != 3 {
		err = errors.New("org plate invalid")
		return
	}
	pinfo.ProvinceCode, err = strconv.Atoi(ss[0])
	if err != nil {
		return
	}
	pinfo.ProvinceChar = nfconst.PPCharMap[pinfo.ProvinceCode]
	pinfo.CityCode = ss[1]
	org := ss[2]
	pinfo.PlateNo = string(([]byte(org))[0:5])
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
	_, _err2 := nfutil.WriteLocalFile(name, imgdata)
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
	_err := nfutil.CreateLocalFolderIfNotExist(_path)
	if _err != nil {
		return "", _err
	}
	return _path + _name, nil
}
