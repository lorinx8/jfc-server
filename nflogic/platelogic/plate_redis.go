package platelogic

import (
	"fmt"
	"reflect"

	"jfcsrv/nfredis"

	"github.com/garyburd/redigo/redis"
)

/*
key:     plate_temp:serial:{设备编号}:bid:{bid序号}:nid:{nid序号}
value:   value中采取hash map进行保存
	 field: Last_plate_No  			value: 上一个上报的车牌号
	 field: Last_plate_img  		value: 上一个上报的车牌区域图存储路径
	 field: Using_plate_No    		value: 正在使用的车牌号
	 field: Using_plate_img 		value: 正在使用的车牌区域图存储路径
	 field: Using_plate_province 	value: 正在使用的车牌号的省份代码
	 field: Using_plate_city		value: 正在使用的车牌号的城市字母
	 field: Crop_img   				value: 上一个上报的大截图存储路径
	 field: Like_count				value: 计数器，前后两次上报车牌相似的次数
	 field: Updatetime				value: longlong时间
*/

type PlateCacheTemp struct {
	Last_plate_No        string
	Last_plate_img       string
	Using_plate_No       string
	Using_plate_img      string
	Using_plate_province string
	Using_plate_city     string
	Crop_img             string
	Like_count           string
	Updatetime           string
}

// 生成key
func generateKey(serial string, bid int, nid int) (key string) {
	key = fmt.Sprintf("plate_temp:serial:%s:bid:%d:nid:%d", serial, bid, nid)
	return key
}

// 新增或更新
func addOrUpdatePlateTemp(serial string, bid int, nid int, data *PlateCacheTemp) (ret string, err error) {
	key := generateKey(serial, bid, nid)
	s := reflect.ValueOf(data).Elem()
	typeOfT := s.Type()
	var args []interface{} = make([]interface{}, 0, 32)
	args = append(args, key)
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		// fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
		args = append(args, typeOfT.Field(i).Name, f.Interface())
	}
	ret, err = nfredis.Hmset(args...)
	return
}

func getPlateTempInCache(serial string, bid int, nid int) (ret *PlateCacheTemp, err error) {
	key := generateKey(serial, bid, nid)
	values, err1 := nfredis.Hgetall(key)
	if err1 != nil {
		return nil, err1
	}
	if len(values) == 0 {
		return nil, nil
	}
	ret = new(PlateCacheTemp)
	err2 := redis.ScanStruct(values, ret)
	if err2 != nil {
		return nil, err2
	}
	return
}
