package nfnet

import (
	"encoding/binary"
	"errors"

	"jfcsrv/nfconst"
	"jfcsrv/nflog"
)

type OrgStreamHandler struct {
	n_data_copyed  int
	n_pack_data    int
	n_data_remain  int
	is_new_package bool
	buffer         []byte
}

var jlog = nflog.Logger

func NewOrgStreamHandler() *OrgStreamHandler {
	return &OrgStreamHandler{
		buffer: make([]byte, 1024*100),
	}
}

func (h *OrgStreamHandler) AddStream(src []byte) (done bool, buf []byte, err error) {

	// 本次的包长度
	n_cur := len(src)

	if n_cur >= nfconst.LEN_MIN_PACKAGE {
		// 已经是一个完整的包, 不需要再处理了
		if src[0] == nfconst.SOCK_PACK_HEADER_L && src[1] == nfconst.SOCK_PACK_HEADER_H && src[n_cur-2] == nfconst.SOCK_PACK_ENDER_L && src[n_cur-1] == nfconst.SOCK_PACK_ENDER_H {
			jlog.Trace("a complete package recieved, n =", len(src))
			return true, src, nil
		}
	}

	// 是不是一个新来的包
	if n_cur >= 2 && src[0] == nfconst.SOCK_PACK_HEADER_L && src[1] == nfconst.SOCK_PACK_HEADER_H {
		h.is_new_package = true
		h.n_data_copyed = 0
		h.n_pack_data = 0
		h.n_data_remain = 0
	} else {
		h.is_new_package = false
	}

	// 如果不是一个新包, 但此时n_data_copyed为零值,理论上不应该出现这样的情况
	if h.is_new_package == false && h.n_data_copyed == 0 {
		return false, nil, errors.New("not a new package but with a zero values copyed data")
	}

	// 是一个新包
	if h.is_new_package == true {
		// 数据负载长度
		var n_pay_load uint16 = binary.BigEndian.Uint16(src[3:5])
		// 数据长度
		h.n_pack_data = int(n_pay_load + 8)
		// 接受该包剩余的接受长度
		h.n_data_remain = h.n_pack_data

		jlog.Trace("a new package, n_payload=", n_pay_load, ", n_pack_data=", h.n_pack_data, ", n_data_ramain=", h.n_data_remain)
		copy(h.buffer[0:], src[0:n_cur])
		h.n_data_copyed = h.n_data_copyed + n_cur
		h.n_data_remain = h.n_data_remain - n_cur
		jlog.Trace(n_cur, "bytes copyed to package buffer, total ", h.n_data_copyed, ", bytes copyed,", h.n_data_remain, " bytes remained")

	} else { // 不是一个新包
		if n_cur <= h.n_data_remain {
			copy(h.buffer[h.n_data_copyed:], src[0:n_cur])
			h.n_data_copyed = h.n_data_copyed + n_cur
			h.n_data_remain = h.n_data_remain - n_cur
			jlog.Trace("a remained package, ", n_cur, " recieved this time, ", h.n_data_copyed, " bytes copyed, ", h.n_data_remain, " bytes remained")
		} else {
			// 如果是这种情况, 说明上一个包后面紧跟着下一个包, 此种情况完全有可能,需要处理
		}
	}

	if h.n_data_remain == 0 {
		jlog.Trace("n_data_remain = 0, data copyed ", h.n_data_copyed)
		_n_pack_data := h.n_pack_data
		h.n_pack_data = 0
		h.n_data_copyed = 0
		h.n_data_remain = 0
		return true, h.buffer[0:_n_pack_data], nil
	} else {
		return false, h.buffer[0:h.n_data_copyed], nil
	}
}
