package platelogic

import "testing"

import "jfcsrv/nfutil"
import "jfcsrv/nfconst"

func TestPlateRedisAddOrUpdate(t *testing.T) {
	nfconst.InitialConst()
	plate := PlateCacheTemp{
		Last_plate_No:        "BL879",
		Last_plate_img:       "/pic/test.jpg",
		Using_plate_No:       "BL987",
		Using_plate_img:      "/pic/test.jpg",
		Using_plate_province: "7",
		Using_plate_city:     "B",
		Crop_img:             "/pic/test.crop.jpg",
		Like_count:           "1",
		Updatetime:           nfutil.GetNowString(),
	}
	ss, err := addOrUpdatePlateTemp("sz1234567890", 1, 1, &plate)
	t.Log(ss, err)
}

func TestPlateRedisGetPlateTemp(t *testing.T) {
	nfconst.InitialConst()
	s, err := getPlateTempInCache("sz1234567890", 1, 1)
	if s == nil {
		t.Log("ss nil")
	} else {
		t.Log(s, err)
	}
}
