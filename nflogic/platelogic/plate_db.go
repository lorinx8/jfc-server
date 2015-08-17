package platelogic

import (
	"fmt"
	"jfcsrv/nfdb"
)

type PlateResultToDb struct {
	Serial       string
	Bid          int
	Nid          int
	ParkStatus   int
	ProvinceCode int
	ProvinceChar string
	CityCode     string
	PlateNo      string
	PlateLiteral string
	PlateImg     string
}

func addOrUpdataPlateResultToDb(r *PlateResultToDb) (err error) {
	db, err1 := nfdb.GetConn()
	if err1 != nil {
		return err1
	}
	fmt.Println(r)
	stmt, err2 := db.Prepare("select nf_save_plate_result ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)")
	if err2 != nil {
		fmt.Println(err2)
		return err2
	}
	_, err3 := stmt.Exec(r.Serial, r.Bid, r.Nid, r.ParkStatus, r.ProvinceCode, r.ProvinceChar, r.CityCode, r.PlateNo, r.PlateLiteral, r.PlateImg)
	if err3 != nil {
		return err3
	}
	defer db.Close()
	return nil
}
