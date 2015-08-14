package nfdb

type angleParamDb struct {
	Id           int
	DeviceSerial string
	RefMapBlockId   string
	Bid          int
	Bangle       int
	Nid          int
	Nangle       int
	CropX        int
	CropY        int
	CropW        int
	CropH        int
	Remark       string
}

func QueryAngleParamByDeviceSerial(serial string) (params []angleParamDb, err error) {
	db, err1 := getConn()
	if err != nil {
		return nil, err1
	}

	rows, err2 := db.Query("select * from tbl_jfcp_angle_param order by bid, nid")
	if err2 != nil {
		return nil, err2
	}

	params = make([]angleParamDb, 0, 10)
	for rows.Next() {
		var ss angleParamDb
		err = rows.Scan(&ss.Id, &ss.DeviceSerial, &ss.RefMapBlockId, &ss.Bid, &ss.Bangle, &ss.Nid, &ss.Nangle, &ss.CropX, &ss.CropY, &ss.CropW, &ss.CropH, &ss.Remark)
		params = append(params, ss)
	}
	return params, nil
}
