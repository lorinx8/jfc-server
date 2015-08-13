package nflogic

import (
	"errors"
	"fmt"
	"jfcsrv/nfconst"
	"jfcsrv/nfnet"
)

type LogicHandler interface {
	// 入参为负载字节切片
	OnLogicMessage(payload []byte) (cmd byte, ret []byte, err error)
}

// 消息统一到达此处进行处理, 根据cmd的不同, 进行消息分发
func OnMsessage(msg []byte) (ret []byte, err error) {
	pkg, err1 := nfnet.NewNFRequestPackage(msg)
	if err1 != nil {
		err = err1
		return nil, err
	}
	fmt.Printf("OnMessage: cmd = %d, length = %d\n", pkg.Cmd, pkg.Length)

	var handler LogicHandler
	switch pkg.Cmd {
	case nfconst.CMD_REQUEST_PARAM:
		handler = &ParamLogic{}
	case nfconst.CMD_REQUEST_ONE_ANGLE_RESULT:
		handler = &AngleResultLogic{}
	default:
		err = errors.New("no command handler")
		return nil, err
	}

	//返回的为数据负载字节
	//需要打包为可以发出去的东东
	cmd, ret_payload, err2 := handler.OnLogicMessage(pkg.Payload)

	if err2 != nil {
		return nil, err2
	}

	ret, err = nfnet.NewNFResponseBytes(cmd, ret_payload)
	return
}
