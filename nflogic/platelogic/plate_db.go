package platelogic

import (
	"jfcsrv/nfdb"
	"jfcsrv/nflog"
)

type PlateResultToDb struct {
	Serial          string
	Bid             int
	Nid             int
	ParkStatus      int
	ProvinceCode    int
	ProvinceChar    string
	CityCode        string
	PlateNo         string
	PlateLiteral    string
	PlateImgUnique  string
	PlateImgHistory string
}

var jlog = nflog.Logger

func addOrUpdataPlateResultToDb(r *PlateResultToDb) (err error) {
	db, err1 := nfdb.GetConn()
	if err1 != nil {
		return err1
	}
	stmt, err2 := db.Prepare("select nf_save_plate_result ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)")
	if err2 != nil {
		jlog.Error("db prepare stmt error: ", err2)
		return err2
	}
	_, err3 := stmt.Exec(r.Serial, r.Bid, r.Nid, r.ParkStatus, r.ProvinceCode, r.ProvinceChar, r.CityCode, r.PlateNo, r.PlateLiteral, r.PlateImgUnique, r.PlateImgHistory)
	if err3 != nil {
		jlog.Error("db exec stmt error: ", err3)
		return err3
	}
	defer db.Close()
	return nil
}
