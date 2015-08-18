package nfconst

import (
	"github.com/Unknwon/goconfig"
)

const (
	CONFIG_FILENAME string = "cfg/j2.cfg"

	// sock command and socket related
	CMD_REQUEST_PARAM          byte = 10
	CMD_REQUEST_PARAM_RESPONSE byte = 11
	CMD_PARAM_TYPE_ANGLE       byte = 1

	CMD_REQUEST_ONE_ANGLE_RESULT          byte = 12
	CMD_REQUEST_ONE_ANGLE_RESULT_RESPONSE byte = 13
	CMD_ONE_ANGLE_RESULT_NO_NEED_CROP     byte = 0
	CMD_ONE_ANGLE_RESULT_NEED_CROP        byte = 1

	CMD_BAD_RESPONSE byte = 255

	SOCK_PACK_HEADER_L byte = 0xF5
	SOCK_PACK_HEADER_H byte = 0xA6
	SOCK_PACK_ENDER_L  byte = 0xBE
	SOCK_PACK_ENDER_H  byte = 0xEF

	// length
	LEN_PACKAGE_EXTRA_DATE int = 8
	LEN_DEVICE_SERIAL      int = 12
	LEN_MAX_PLATE_NUMBER   int = 12        // 车牌最大字符长度
	LEN_MAX_PLATE_IMG_SIZE int = 1024 * 15 // 车牌图片最多的字节数
	LEN_TCP_BUFFER         int = 2048
	LEN_MIN_PACKAGE        int = 8 // 一个数据包的最小字节

	FILENAME_IMG_EXTENT string = ".jpg"

	// 相似度临界值
	PLATE_SIMI_THRESHOLD int = 60
	// 连续多少次，车牌相似度都小于临界值，则接受该车牌
	PLATE_SIMI_MAX_COMPARE int = 1

	PP_CHUAN int = 0  /* "zh_cuan" 川 */
	PP_E     int = 1  /* "zh_e" 鄂 */
	PP_GAN   int = 2  /* "zh_gan" 赣*/
	PP_GAN1  int = 3  /* "zh_gan1" 甘*/
	PP_GUI   int = 4  /* "zh_gui" 贵 */
	PP_GUI1  int = 5  /* "zh_gui1" 桂 */
	PP_HEI   int = 6  /* "zh_hei" 黑 */
	PP_HU    int = 7  /* "zh_hu" 沪 */
	PP_JI    int = 8  /* "zh_ji" 冀 */
	PP_JIN   int = 9  /* "zh_jin" 津 */
	PP_JING  int = 10 /* "zh_jing" 京 */
	PP_JL    int = 11 /* "zh_jl" 吉 */
	PP_LIAO  int = 12 /* "zh_liao" 辽 */
	PP_LU    int = 13 /* "zh_lu" 鲁 */
	PP_MENG  int = 14 /* "zh_meng" 蒙 */
	PP_MIN   int = 15 /* "zh_min" 闽 */
	PP_NING  int = 16 /* "zh_ning" 宁 */
	PP_QING  int = 17 /* "zh_qing" 青 */
	PP_QIONG int = 18 /* "zh_qiong" 琼 */
	PP_SHAN  int = 19 /* "zh_shan" 陕 */
	PP_SU    int = 20 /* "zh_su" 苏 */
	PP_SX    int = 21 /* "zh_sx" 晋 */
	PP_WAN   int = 22 /* "zh_wan" 皖 */
	PP_XIANG int = 23 /* "zh_xiang" 湘 */
	PP_XIN   int = 24 /* "zh_xin" 新 */
	PP_YU    int = 25 /* "zh_yu" 豫 */
	PP_YU1   int = 26 /* "zh_yu1" 渝 */
	PP_YUE   int = 27 /* "zh_yue" 粤 */
	PP_YUN   int = 28 /* "zh_yun" 云 */
	PP_ZANG  int = 29 /* "zh_zang" 藏 */
	PP_ZHE   int = 30 /* "zh_zhe" 浙 */

	PP_CHUAN_CHAR string = "川"
	PP_E_CHAR     string = "鄂"
	PP_GAN_CHAR   string = "赣"
	PP_GAN1_CHAR  string = "甘"
	PP_GUI_CHAR   string = "贵"
	PP_GUI1_CHAR  string = "桂"
	PP_HEI_CHAR   string = "黑"
	PP_HU_CHAR    string = "沪"
	PP_JI_CHAR    string = "冀"
	PP_JIN_CHAR   string = "津"
	PP_JING_CHAR  string = "京"
	PP_JL_CHAR    string = "吉"
	PP_LIAO_CHAR  string = "辽"
	PP_LU_CHAR    string = "鲁"
	PP_MENG_CHAR  string = "蒙"
	PP_MIN_CHAR   string = "闽"
	PP_NING_CHAR  string = "宁"
	PP_QING_CHAR  string = "青"
	PP_QIONG_CHAR string = "琼"
	PP_SHAN_CHAR  string = "陕"
	PP_SU_CHAR    string = "苏"
	PP_SX_CHAR    string = "晋"
	PP_WAN_CHAR   string = "皖"
	PP_XIANG_CHAR string = "湘"
	PP_XIN_CHAR   string = "新"
	PP_YU_CHAR    string = "豫"
	PP_YU1_CHAR   string = "渝"
	PP_YUE_CHAR   string = "粤"
	PP_YUN_CHAR   string = "云"
	PP_ZANG_CHAR  string = "藏"
	PP_ZHE_CHAR   string = "浙"
)

type JFCConfig struct {
	JFCPort      string
	DbConnString string
	RedisServer  string
	RedisPort    string
	RedisDbIndex int
}

var (
	JCfg      *JFCConfig
	PPCharMap map[int]string
)

func initPP() {
	PPCharMap = make(map[int]string, 33)
	PPCharMap[PP_CHUAN] = PP_CHUAN_CHAR
	PPCharMap[PP_E] = PP_E_CHAR
	PPCharMap[PP_GAN] = PP_GAN_CHAR
	PPCharMap[PP_GAN1] = PP_GAN1_CHAR
	PPCharMap[PP_GUI] = PP_GUI_CHAR
	PPCharMap[PP_GUI1] = PP_GUI1_CHAR
	PPCharMap[PP_HEI] = PP_HEI_CHAR
	PPCharMap[PP_HU] = PP_HU_CHAR
	PPCharMap[PP_JI] = PP_JI_CHAR
	PPCharMap[PP_JIN] = PP_JIN_CHAR
	PPCharMap[PP_JING] = PP_JING_CHAR
	PPCharMap[PP_JL] = PP_JL_CHAR
	PPCharMap[PP_LIAO] = PP_LIAO_CHAR
	PPCharMap[PP_LU] = PP_LU_CHAR
	PPCharMap[PP_MENG] = PP_MENG_CHAR
	PPCharMap[PP_MIN] = PP_MIN_CHAR
	PPCharMap[PP_NING] = PP_NING_CHAR
	PPCharMap[PP_QING] = PP_QING_CHAR
	PPCharMap[PP_QIONG] = PP_QIONG_CHAR
	PPCharMap[PP_SHAN] = PP_SHAN_CHAR
	PPCharMap[PP_SU] = PP_SU_CHAR
	PPCharMap[PP_SX] = PP_SX_CHAR
	PPCharMap[PP_WAN] = PP_WAN_CHAR
	PPCharMap[PP_XIANG] = PP_XIANG_CHAR
	PPCharMap[PP_XIN] = PP_XIN_CHAR
	PPCharMap[PP_YU] = PP_YU_CHAR
	PPCharMap[PP_YU1] = PP_YU1_CHAR
	PPCharMap[PP_YUE] = PP_YUE_CHAR
	PPCharMap[PP_YUN] = PP_YUN_CHAR
	PPCharMap[PP_ZANG] = PP_ZANG_CHAR
	PPCharMap[PP_ZHE] = PP_ZHE_CHAR
}

func InitialConst() error {
	initPP()
	err := loadConfig(CONFIG_FILENAME)
	return err
}

func loadConfig(path string) error {

	JCfg = &JFCConfig{}

	c, err := goconfig.LoadConfigFile(path)
	if err != nil {
		return err
	}

	JCfg.JFCPort, err = c.GetValue("Server", "Port")
	if err != nil {
		return err
	}

	JCfg.DbConnString, err = c.GetValue("Database", "DbConnString")
	if err != nil {
		return err
	}
	JCfg.RedisServer, err = c.GetValue("Redis", "Server")
	if err != nil {
		return err
	}
	JCfg.RedisPort, err = c.GetValue("Redis", "Port")
	if err != nil {
		return err
	}
	JCfg.RedisDbIndex, err = c.Int("Redis", "DbIndex")
	if err != nil {
		return err
	}
	return nil
}
