package platelogic

import (
	"fmt"
	"jfcsrv/nfconst"
	"jfcsrv/nfutil"
)

type CropImageLogic struct {
}

func (logic *CropImageLogic) OnLogicMessage(msg []byte) (cmd byte, ret []byte, err error) {
	cmd = nfconst.CMD_REPORT_CROP_IMAGE_RESPONSE
	ret = []byte{0x0}
	err = nil
	return
}

// 生成截图文件云端存储的key，包含日期序列，每个角度包含日期序列
// picp/{年-月-日}/{设备号}/{bid}-{nid}_{时-分-秒}_crop.jpg
func (logic *CropImageLogic) generateCloudCropImageKeyHistory(serial string, bid int, nid int) (key string) {
	y, mon, d, h, min, s := nfutil.GetNow()
	key = fmt.Sprintf("picp_history/%04d-%02d-%02d/%s/%02d-%02d_%02d-%02d-%02d_crop%s", y, mon, d, serial, bid, nid, h, min, s, nfconst.FILENAME_IMG_EXTENT)
	return
}

// 生成截图文件云端存储的key, 不包含日期序列，每个角度唯一一个
func (logic *CropImageLogic) generateCloudCropImageKeyUnique(serial string, bid int, nid int) (key string) {
	key = fmt.Sprintf("picp_u/%s/%02d-%02d_crop%s", serial, bid, nid, nfconst.FILENAME_IMG_EXTENT)
	return
}
