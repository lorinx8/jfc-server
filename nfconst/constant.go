package nfconst

const (
	// sock command and socket related
	CMD_REQUEST_PARAM            byte = 10
	CMD_REQUEST_PARAM_RESPONSE   byte = 11
	CMD_PARAM_TYPE_ANGLE         byte = 1
	CMD_REQUEST_ONE_ANGLE_RESULT byte = 12

	SOCK_PACK_HEADER_L byte = 0xF5
	SOCK_PACK_HEADER_H byte = 0xA6
	SOCK_PACK_ENDER_L  byte = 0xBE
	SOCK_PACK_ENDER_H  byte = 0xEF

	// length
	LEN_PACKAGE_EXTRA_DATE int = 8
	LEN_DEVICE_SERIAL      int = 12
)
