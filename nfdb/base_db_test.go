package nfdb

import (
	"fmt"
	"testing"
	"time"
)

func T_estBaseDb(t *testing.T) {

	// park_number | map_block_id | bid | bangle | nid | nangle | crop_x | crop_y | crop_w | crop_h | remark

	for i := 0; i < 3; i++ {

		db, err := getDb()
		if err == nil {
			fmt.Println(db)
		} else {
			fmt.Println(err)
		}

		rows, err1 := db.Query("select * from tbl_jfcp_angle_param")
		if err1 != nil {
			fmt.Println(err1)
		} else {
			for rows.Next() {
				var ss angleParam
				err = rows.Scan(&ss.id, &ss.deviceSerial, &ss.parkNumber, &ss.mapBlockId, &ss.bid, &ss.bangle, &ss.nid, &ss.nangle, &ss.cropX, &ss.cropY, &ss.cropW, &ss.cropH, &ss.remark)
				fmt.Println(ss)
			}
		}
		time.Sleep(3000 * time.Millisecond)
	}
}
