package nfutil

import (
	"fmt"
	"testing"
)

func T_estWriteFile(t *testing.T) {

	name := "/Users/lorin/tt.bat"
	nn, err := WriteLocalFile(name, []byte{0x31, 0x32})
	if err != nil && nn != 2 {
		t.Error(err)
	}
}

func TestGetNow(t *testing.T) {
	y, mon, d, h, min, s := GetNow()
	fmt.Println(y, mon, d, h, min, s)
}

func TestTT(t *testing.T) {

	bangles := [2]int{20, 40}
	nangles := [2]int{30, 60}
	//INSERT INTO tbl_jfcp_angle_param (device_serial, ref_map_block_id, bid, bangle, nid, nangle, crop_x, crop_y, crop_w, crop_h, remark) VALUES ('sz1234567890', 'rB2187', 1, 30, 2, 50, 400, 300, 640, 320, '');
	for i := 1; i <= 500; i++ {
		serial := fmt.Sprintf("yl%010d", i)
		for bid := 0; bid < 2; bid++ {
			for nid := 0; nid < 2; nid++ {
				map_id := fmt.Sprintf("rB00%d%d", bid, nid)
				bangle := bangles[bid]
				nangle := nangles[nid]
				sql := fmt.Sprintf("INSERT INTO tbl_jfcp_angle_param (device_serial, ref_map_block_id, bid, bangle, nid, nangle, crop_x, crop_y, crop_w, crop_h, remark) VALUES ('%s', '%s', %d, %d, %d, %d, 400, 300, 640, 320, '');", serial, map_id, bid, bangle, nid, nangle)
				fmt.Println(sql)
			}
		}
	}
}

func TestTT2(t *testing.T) {
	for i := 1; i <= 500; i++ {
		serial := fmt.Sprintf("yl%010d", i)
		for bid := 0; bid < 2; bid++ {
			for nid := 0; nid < 2; nid++ {
				map_id := fmt.Sprintf("rB-%d-%d-%d", i, bid, nid)
				sql := fmt.Sprintf("UPDATE tbl_jfcp_angle_param SET ref_map_block_id='%s' WHERE device_serial='%s' and bid=%d and nid=%d;", map_id, serial, bid, nid)
				fmt.Println(sql)
			}
		}
	}
}
