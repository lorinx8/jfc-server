package nfredis

import (
	"fmt"
	"reflect"

	"github.com/garyburd/redigo/redis"
)

/*
key:     plate_temp:serial:{设备编号}:bid:{bid序号}:nid:{nid序号}
value:   value中采取hash map进行保存
	 field: last_plate  			value: 上一个上报的车牌号
	 field: last_plate_img  		value: 上一个上报的车牌区域图存储路径
	 field: using_plate     		value: 正在使用的车牌号
	 field: using_plate_img 		value: 正在使用的车牌区域图存储路径
	 field: using_plate_province 	value: 正在使用的车牌号的省份代码
	 field: using_plate_city		value: 正在使用的车牌号的城市字母
	 field: crop_img   				value: 上一个上报的大截图存储路径
	 field: like_count				value: 计数器，前后两次上报车牌相似的次数
	 field: updatetime				value: longlong时间
*/

type PlateTemp struct {
	Last_plate           string
	Last_plate_img       string
	Using_plate          string
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
func AddOrUpdate(serial string, bid int, nid int, data *PlateTemp) (ret string, err error) {
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
	ret, err = hmset(args...)
	return
}

func GetPlateTemp(serial string, bid int, nid int) (ret *PlateTemp, err error) {
	key := generateKey(serial, bid, nid)
	ret = &PlateTemp{}
	s := reflect.ValueOf(ret).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		fieldname := typeOfT.Field(i).Name
		v, err1 := hget(key, fieldname)
		if err1 != nil {
			err = err1
			break
		}
		s.FieldByName(fieldname).SetString(v)
	}
	return
}

func GetPlateTemp2(serial string, bid int, nid int) (ret *PlateTemp, err error) {
	key := generateKey(serial, bid, nid)
	ret = &PlateTemp{}

	values, err1 := hgetall(key)
	if err1 != nil {
		return nil, err1
	}
	if err2 := redis.ScanStruct(values, ret); err2 != nil {
		return nil, err2
	}
	return
}
