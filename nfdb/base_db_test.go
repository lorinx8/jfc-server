package nfdb

import (
	"fmt"
	"testing"
	"time"
)

func T_estBaseDb(t *testing.T) {

	// park_number | map_block_id | bid | bangle | nid | nangle | crop_x | crop_y | crop_w | crop_h | remark

	for i := 0; i < 3; i++ {

		db, err := getConn()
		if err == nil {
			fmt.Println(db)
		} else {
			fmt.Println(err)
		}

		_, err1 := db.Query("select 1")
		if err1 != nil {
			fmt.Println(err1)
		}
		time.Sleep(3000 * time.Millisecond)
	}
}
