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

type PlateNumberInfo struct {
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
	//msgChan := make(chan []byte)
	handlePlateReport(msg)
	//msgChan <- msg[:]

	return _cmd, nil, nil
}

// 处理车牌结果
func handlePlateReport(msg []byte) {
	// 解析角度结果数据包
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
		handlePlateInPool(serial, int(t.Bid), int(t.Nid), &PlateNumberInfo{})
		return
	}

	// 包含车牌
	orgString := string(t.PlateNumber[:])
	pinfo, _err := parseOrgPlateString(orgString)
	if _err != nil {
		// 车牌都解析不成功，不接续了
		fmt.Println("plate parse error, return", _err)
		return
	}
	pinfo.ImageByte = img_data // 好了，现在车牌信息全了

	// 开始处理和比较车牌
	handlePlateInPool(serial, int(t.Bid), int(t.Nid), &pinfo)
	// 。。。
}

// 处理车牌识别结果, 判断是否用本次的上报数据
func handlePlateInPool(serial string, bid int, nid int, pinfo *PlateNumberInfo) error {
	fmt.Println("enter handle plate in pool")

	// 从缓存中取得已有的车牌缓存结果
	plateCacheTemp, err := getPlateTempInCache(serial, bid, nid)
	if err != nil {
		fmt.Println("get plate in cache err", err, serial, bid, nid)
		return err // 不再处理这个了
	}

	// 说明这个设备的这个位置，还没有数据
	if plateCacheTemp == nil {
		fmt.Println("cond 2: no data in cache for", serial, bid, nid, ",accept it")
		// 处理第一个数据， 直接接受这个车牌
		acceptPlateNumber(serial, bid, nid, pinfo)
	} else {
		_last_plate_no := plateCacheTemp.Last_plate_No
		_using_plate_no := plateCacheTemp.Using_plate_No
		_new_plate_no := pinfo.PlateNo

		fmt.Println("cond 3: have data in cache for", serial, bid, nid, _new_plate_no, plateCacheTemp, ",do extra work")
		// cache中已经存在
		// 相似度比较，使用last中的值进行相似度比较

		// 新上报的车牌就是正在使用中的车牌
		if _using_plate_no == _new_plate_no {
			fmt.Println("cond 3.1: new reported plate is the same as the using one", _using_plate_no)
			// 只需要更新update时间即可
			updateCacheTime(serial, bid, nid, plateCacheTemp)
			return nil
		}

		similarty := calSimilarity(_new_plate_no, _last_plate_no)
		// 有三个字符都不一样
		if similarty < nfconst.PLATE_SIMI_THRESHOLD {
			fmt.Println("cond 3.2: similarty below threshold, save it temporary")
			setCachePlateTempLikeCount(plateCacheTemp, 0)
			// 暂存新的
			saveInCacheTemporary(serial, bid, nid, pinfo, plateCacheTemp)
		} else {
			// 三个字符都一样,并且大于了最多比较次数，则接受
			_count := getCachePlateTempLikeCount(plateCacheTemp)
			if _count > nfconst.PLATE_SIMI_MAX_COMPARE {
				fmt.Println("cond 3.3: similarty above threshold > MAX_COMPARE, accept it")
				acceptPlateNumber(serial, bid, nid, pinfo)
			} else {
				fmt.Println("cond 3.4: similarty above threshold, save it temporary")
				increCachePlateTempLikeCount(plateCacheTemp)
				saveInCacheTemporary(serial, bid, nid, pinfo, plateCacheTemp)
			}
		}
	}
	return nil
}

func setCachePlateTempLikeCount(p *PlateCacheTemp, count int) {
	str := strconv.Itoa(count)
	p.Like_count = str
}

func getCachePlateTempLikeCount(p *PlateCacheTemp) (count int) {
	i, err := strconv.Atoi(p.Like_count)
	if err != nil {
		return 0
	} else {
		return i
	}
}

func increCachePlateTempLikeCount(p *PlateCacheTemp) (count int) {
	i := getCachePlateTempLikeCount(p)
	count = i + 1
	setCachePlateTempLikeCount(p, count)
	return count
}

// 更新时间
func updateCacheTime(serial string, bid int, nid int, c *PlateCacheTemp) {
	c.Updatetime = nfutil.GetNowString()
	addOrUpdatePlateTemp(serial, bid, nid, c)
}

func saveInCacheTemporary(serial string, bid int, nid int, pinfo *PlateNumberInfo, cache *PlateCacheTemp) {
	cache.Last_plate_No = pinfo.PlateNo

	// 上传车牌数据
	if pinfo.PlateNo != "" {
		cloud_key := generateCloudFileKey(serial, bid, nid)
		url, err1 := nfutil.PutLocalToCloud(pinfo.ImageByte, int64(len(pinfo.ImageByte)), cloud_key)
		if err1 != nil {
			// 打印日志
			fmt.Printf("PutLocalToCloud err", err1)
		}
		cache.Last_plate_img = url
	}
	// 更新时间
	updateCacheTime(serial, bid, nid, cache)
}

// 采用这个车牌，需要做这么几件事情
// 更新缓存数据，写入数据库，向云存储中上传数据
func acceptPlateNumber(serial string, bid int, nid int, pinfo *PlateNumberInfo) (ret *PlateCacheTemp) {
	// 对于相同的键，直接写入既可，redis自己覆盖掉
	var pcache *PlateCacheTemp
	var platestatus int
	var url string
	if pinfo.PlateNo == "" {
		// 无车牌
		pcache = &PlateCacheTemp{
			Last_plate_No:        "",
			Last_plate_img:       "",
			Using_plate_No:       "",
			Using_plate_province: "",
			Using_plate_city:     "",
			Using_plate_img:      "",
			Crop_img:             "",
			Like_count:           "0",
			Updatetime:           nfutil.GetNowString(),
		}
		platestatus = 0
	} else {
		// 上传车牌数据
		if pinfo.PlateNo != "" {
			cloud_key := generateCloudFileKey(serial, bid, nid)
			_url, err1 := nfutil.PutLocalToCloud(pinfo.ImageByte, int64(len(pinfo.ImageByte)), cloud_key)
			if err1 != nil {
				// 打印日志
				fmt.Printf("PutLocalToCloud err", err1)
				url = ""
			} else {
				url = _url
			}
		}

		// 图片未上传成功也需要继续
		// 有车牌
		pcache = &PlateCacheTemp{
			Last_plate_No:        pinfo.PlateNo,
			Last_plate_img:       url,
			Using_plate_No:       pinfo.PlateNo,
			Using_plate_province: strconv.Itoa(pinfo.ProvinceCode),
			Using_plate_city:     pinfo.CityCode,
			Using_plate_img:      url,
			Crop_img:             "",
			Like_count:           "0",
			Updatetime:           nfutil.GetNowString(),
		}
		platestatus = 1
	}
	addOrUpdatePlateTemp(serial, bid, nid, pcache)

	// 然后保存到数据库里面去
	var r *PlateResultToDb = &PlateResultToDb{
		Serial:       serial,
		Bid:          bid,
		Nid:          nid,
		ParkStatus:   platestatus,
		ProvinceCode: pinfo.ProvinceCode,
		ProvinceChar: nfconst.PPCharMap[pinfo.ProvinceCode],
		CityCode:     pinfo.CityCode,
		PlateNo:      pinfo.PlateNo,
		PlateLiteral: nfconst.PPCharMap[pinfo.ProvinceCode] + pinfo.CityCode + " " + pinfo.PlateNo,
		PlateImg:     url,
	}
	addOrUpdataPlateResultToDb(r)
	return pcache
}

// 比较新旧车牌号码的相似度
func calSimilarity(plate_new string, plate_old string) (ret int) {
	if plate_old == "" && plate_new != "" {
		return 0
	}
	if plate_new == "" && plate_old != "" {
		return 0
	}
	if plate_new == "" && plate_old == "" {
		return 100
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
func parseOrgPlateString(orgString string) (pinfo PlateNumberInfo, err error) {
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
func savePlateImageLocal(serial string, bid int, nid int, imgdata []byte) (err error) {
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
func generateOrgFilePath(serial string, bid int, nid int) (path string, err error) {
	y, mon, d, h, min, s := nfutil.GetNow()
	_path := fmt.Sprintf("picp/%s/%04d-%02d-%02d/", serial, y, mon, d)
	_name := fmt.Sprintf("%02d-%02d_%02d-%02d-%02d%s", bid, nid, h, min, s, nfconst.FILENAME_IMG_EXTENT)
	_err := nfutil.CreateLocalFolderIfNotExist(_path)
	if _err != nil {
		return "", _err
	}
	return _path + _name, nil
}

// 生成
func generateCloudFileKey(serial string, bid int, nid int) (key string) {
	y, mon, d, h, min, s := nfutil.GetNow()
	key = fmt.Sprintf("picp/%s/%04d-%02d-%02d/%02d-%02d_%02d-%02d-%02d%s", serial, y, mon, d, bid, nid, h, min, s, nfconst.FILENAME_IMG_EXTENT)
	return
}
