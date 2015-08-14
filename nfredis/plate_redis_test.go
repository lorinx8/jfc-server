package nfredis

import "testing"

import "jfcsrv/nfutil"

func TestPlateRedisAddOrUpdate(t *testing.T) {
	plate := PlateTemp{
		Last_plate:           "BL879",
		Last_plate_img:       "/pic/test.jpg",
		Using_plate:          "BL987",
		Using_plate_img:      "/pic/test.jpg",
		Using_plate_province: "7",
		Using_plate_city:     "B",
		Crop_img:             "/pic/test.crop.jpg",
		Like_count:           "1",
		Updatetime:           nfutil.GetNowString(),
	}
	ss, err := AddOrUpdate("sz1234567890", 1, 1, &plate)
	t.Log(ss, err)
}

func TestPlateRedisGetPlateTemp(t *testing.T) {
	s, err := GetPlateTemp2("sz12345678rr", 1, 1)
	t.Log(s, err)
}
