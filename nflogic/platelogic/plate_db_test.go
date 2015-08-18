package platelogic

import (
	"testing"
)

func TestAddOrUpdataPlateResult(t *testing.T) {
	var r *PlateResultToDb = &PlateResultToDb{
		Serial:       "sz1234567890",
		Bid:          1,
		Nid:          1,
		ParkStatus:   0,
		ProvinceCode: 7,
		ProvinceChar: "粤",
		CityCode:     "B",
		PlateNo:      "8LB56",
		PlateLiteral: "粤B 8LB56",
		PlateImg:     "http://7xl4c2.com1.z0.glb.clouddn.com/picp/sz1234567890/2015-08-17/01-01_15-33-17.jpg",
	}

	err := addOrUpdataPlateResultToDb(r)
	t.Log(err)

}
