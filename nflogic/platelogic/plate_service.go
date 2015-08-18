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
	_ret := handlePlateReport(msg)
	ret = make([]byte, 1)
	ret[0] = _ret
	return _cmd, ret, nil
}

// 处理车牌结果
func handlePlateReport(msg []byte) (ret byte) {
	// 解析角度结果数据包
	ret = nfconst.CMD_ONE_ANGLE_RESULT_NO_NEED_CROP
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
	jlog.Infof("Device Serial:%s, report angle result: bid=%d, nid=%d, count=%d, img_n=%d\n", t.Serial, t.Bid, t.Nid, t.Count, n_img)
	// 处理上报数据无车牌的情况
	if t.Count == 0 {
		handlePlateInPool(serial, int(t.Bid), int(t.Nid), &PlateNumberInfo{})
		return ret
	}

	// 包含车牌
	orgString := string(t.PlateNumber[:])
	pinfo, _err := parseOrgPlateString(orgString)
	if _err != nil {
		// 车牌都解析不成功，不接续了
		jlog.Error("plate parse error, return", _err)
		return ret
	}
	pinfo.ImageByte = img_data // 好了，现在车牌信息全了

	// 开始处理和比较车牌
	accepe, _ := handlePlateInPool(serial, int(t.Bid), int(t.Nid), &pinfo)
	if accepe {
		ret = nfconst.CMD_ONE_ANGLE_RESULT_NEED_CROP
	}
	return ret
}

// 处理车牌识别结果, 判断是否用本次的上报数据
func handlePlateInPool(serial string, bid int, nid int, pinfo *PlateNumberInfo) (accepe bool, err error) {
	jlog.Debug("enter handle plate in pool")
	// 从缓存中取得已有的车牌缓存结果
	plateCacheTemp, err1 := getPlateTempInCache(serial, bid, nid)
	if err1 != nil {
		jlog.Error("get plate in cache error ", err, serial, bid, nid)
		return false, err1 // 不再处理这个了
	}

	// 说明这个设备的这个位置，还没有数据
	if plateCacheTemp == nil {
		fmt.Println("cond 1: no data in cache for", serial, bid, nid, ", accept it")
		// 处理第一个数据， 直接接受这个车牌
		acceptPlateNumber(serial, bid, nid, pinfo, nil)
		accepe = true
	} else {
		_last_plate_no := plateCacheTemp.Last_plate_No
		_new_plate_no := pinfo.PlateNo
		jlog.Debug("cond 2: have data in cache for", serial, bid, nid, "new palte:", _new_plate_no, "old plate:", plateCacheTemp.Last_plate_No, ", do extra work")

		similarty := calSimilarity(_new_plate_no, _last_plate_no)

		// 有三个字符都不一样
		if similarty < nfconst.PLATE_SIMI_THRESHOLD {
			jlog.Debug("cond 2.1: similarty below threshold, save it temporary")
			setCachePlateTempLikeCount(plateCacheTemp, 0)
			saveInCacheTemporary(serial, bid, nid, pinfo, plateCacheTemp)
		} else {
			// 三个字符都一样,并且大于了最多比较次数，则接受
			_count := getCachePlateTempLikeCount(plateCacheTemp)
			if _count >= nfconst.PLATE_SIMI_MAX_COMPARE {
				jlog.Debug("cond 2.2: similarty above threshold > MAX_COMPARE, accept it")
				acceptPlateNumber(serial, bid, nid, pinfo, plateCacheTemp)
				accepe = true
			} else {
				jlog.Debug("cond 2.3: similarty above threshold, save it temporary")
				increCachePlateTempLikeCount(plateCacheTemp)
				saveInCacheTemporary(serial, bid, nid, pinfo, plateCacheTemp)
			}
		}
	}

	return accepe, nil
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
	addOrUpdatePlateTempCache(serial, bid, nid, c)
}

// 未到接受车牌号条件的时候的处理， 更新last_plate字段和更新时间， 此处不更新存储中的车牌文件
func saveInCacheTemporary(serial string, bid int, nid int, pinfo *PlateNumberInfo, cache *PlateCacheTemp) {
	cache.Last_plate_No = pinfo.PlateNo
	/*
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
	*/
	// 更新时间
	updateCacheTime(serial, bid, nid, cache)
}

// 采用这个车牌，需要做这么几件事情
// 更新缓存数据，写入数据库，向云存储中上传数据
// 如果接受了一个车牌，那么就需要向客户端请求大的截图
func acceptPlateNumber(serial string, bid int, nid int, pinfo *PlateNumberInfo, pcache *PlateCacheTemp) (ret *PlateCacheTemp) {
	// 对于相同的键，直接写入既可，redis自己覆盖掉
	var platestatus int
	var url_unique, url_history string
	var err error

	if pcache == nil {
		pcache = &PlateCacheTemp{}
	}

	if pinfo.PlateNo == "" {
		// 无车牌
		clearPlateCacheTempStruct(pcache)
		platestatus = 0
	} else {
		transferPlateInfoToPlateCacheStruct(pinfo, pcache)
		platestatus = 1
	}
	// 如果新的车牌跟之前已经接受的车牌一致，那么就不用上传图片了, 也不需要更新数据库中的记录了
	// 需要做两件事情， 一个是上传到该角度需要使用的文件，然后再拷贝一份到历史数据中去

	if pinfo.PlateNo != "" && pinfo.PlateNo != pcache.Using_plate_No {
		url_unique, url_history, err = uploadAcceptPlateFile(serial, bid, nid, pinfo.ImageByte, int64(len(pinfo.ImageByte)))
		if err != nil {
			pcache.Last_plate_img = url_history
			pcache.Using_plate_img = url_unique
		}
	}

	addOrUpdatePlateTempCache(serial, bid, nid, pcache)
	var r *PlateResultToDb = &PlateResultToDb{
		Serial:          serial,
		Bid:             bid,
		Nid:             nid,
		ParkStatus:      platestatus,
		ProvinceCode:    pinfo.ProvinceCode,
		ProvinceChar:    nfconst.PPCharMap[pinfo.ProvinceCode],
		CityCode:        pinfo.CityCode,
		PlateNo:         pinfo.PlateNo,
		PlateLiteral:    nfconst.PPCharMap[pinfo.ProvinceCode] + pinfo.CityCode + " " + pinfo.PlateNo,
		PlateImgUnique:  url_unique,
		PlateImgHistory: url_history,
	}
	addOrUpdataPlateResultToDb(r)
	return pcache
}

func clearPlateCacheTempStruct(pcache *PlateCacheTemp) (ret *PlateCacheTemp) {
	pcache.Last_plate_No = ""
	pcache.Last_plate_img = ""
	pcache.Using_plate_No = ""
	pcache.Using_plate_province = ""
	pcache.Using_plate_city = ""
	pcache.Using_plate_img = ""
	pcache.Crop_img = ""
	pcache.Like_count = "0"
	pcache.Updatetime = nfutil.GetNowString()
	return pcache
}

func transferPlateInfoToPlateCacheStruct(pinfo *PlateNumberInfo, pcache *PlateCacheTemp) {
	pcache.Last_plate_No = pinfo.PlateNo
	pcache.Using_plate_No = pinfo.PlateNo
	pcache.Using_plate_province = strconv.Itoa(pinfo.ProvinceCode)
	pcache.Using_plate_city = pinfo.CityCode
	pcache.Like_count = "0"
	pcache.Updatetime = nfutil.GetNowString()
}

// 接受一个新车牌时，向云存储中上传数据，并作对应处理
// 对于一个串号-bid-nid组合，云存储中有一个唯一文件，该文件是最新的车牌 picp_u/{设备串号}/{bid}-{nid}.jpg
// 还有一个是历史的，每接受一个车牌，就会多一个历史的 picp_history/{设备串号}/{yyyy-MM-dd格式的年月日}/{bid}-{nid}_{HH-mm-ss形式的时分秒}.jpg
func uploadAcceptPlateFile(serial string, bid int, nid int, b []byte, size int64) (url_unique string, url_history string, err error) {
	// 先覆盖唯一的文件，然后再拷贝一个历史的
	cloud_key_unique := generateCloudPlateImageKeyUnique(serial, bid, nid)
	cloud_key_history := generateCloudPlateImageKeyHistory(serial, bid, nid)
	// 覆盖唯一的文件
	url_unique, err = nfutil.PutLocalToCloud(b, size, cloud_key_unique)
	if err != nil {
		// 打印日志
		jlog.Error("PutLocalToCloud cloud_key_unique error: ", err)
	}

	// 然后进行拷贝
	err = nfutil.CopyCloudFile(cloud_key_unique, cloud_key_history)
	if err != nil {
		// 打印日志
		jlog.Error("CopyCloudFile error: ", err)
	}
	return
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
// 从JFC过来的格式为  7_B12345 [粤B12345]
// 第一个7为省份号码，B为城市编号，后五位车牌号码
// pcode 省份号码
// pchar 城市字面值，粤
// cchar 城市编号AB
// pnumber 五位号码
func parseOrgPlateString(orgString string) (pinfo PlateNumberInfo, err error) {
	ss := strings.Split(orgString, "_")
	if len(ss) != 2 {
		err = errors.New("org plate invalid")
		return
	}
	pinfo.ProvinceCode, err = strconv.Atoi(ss[0])
	if err != nil {
		return
	}
	pinfo.ProvinceChar = nfconst.PPCharMap[pinfo.ProvinceCode]

	_strb := []byte(ss[1])
	pinfo.CityCode = string(_strb[0])
	pinfo.PlateNo = string(_strb[1:6])
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

// 生成云端存储的key，包含日期序列，每个角度包含日期序列
// picp/{设备号}/{年-月-日}/{bid}-{nid}_{时分秒}.jpg
func generateCloudPlateImageKeyHistory(serial string, bid int, nid int) (key string) {
	y, mon, d, h, min, s := nfutil.GetNow()
	key = fmt.Sprintf("picp_history/%s/%04d-%02d-%02d/%02d-%02d_%02d-%02d-%02d%s", serial, y, mon, d, bid, nid, h, min, s, nfconst.FILENAME_IMG_EXTENT)
	return
}

// 生成云端存储的key, 不包含日期序列，每个角度唯一一个
func generateCloudPlateImageKeyUnique(serial string, bid int, nid int) (key string) {
	key = fmt.Sprintf("picp_u/%s/%02d-%02d%s", serial, bid, nid, nfconst.FILENAME_IMG_EXTENT)
	return
}
