package main

import (
	"jfcsrv/nfconst"
	"jfcsrv/nflog"
	"jfcsrv/nflogic"
	"jfcsrv/nfnet"
	"jfcsrv/nfutil"
	"net"
	"runtime"
	"runtime/pprof"
)

var jlog = nflog.Logger

func main() {
	jlog.Info("JFC Server For Papakaka...")
	jlog.Info("Config loading...")
	nfconst.InitialConst()
	jlog.Info(nfconst.JCfg)
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":"+nfconst.JCfg.JFCPort)
	check_error(err)
	tcpListener, err := net.ListenTCP("tcp", tcpAddr)
	check_error(err)
	jlog.Info("JFC Server listening on port ", nfconst.JCfg.JFCPort)
	for i := 0; i < 100; i++ {
		go handle_tcp_accept(tcpListener)
	}
	select {}
}

func handle_tcp_accept(tcpListener *net.TCPListener) {
	for {
		clientConn, err := tcpListener.AcceptTCP()
		if err != nil {
			jlog.Error("tcp accept failed!")
			continue
		} else {
			go handle_logic(clientConn)
			/*
				inDataChan := make(chan []byte)
				outDataChan := make(chan []byte)
				go read_tcp_conn(tcpConn, inDataChan)
				go business_handler(inDataChan, outDataChan)
				go write_tcp_conn(tcpConn, outDataChan)
			*/
			p := pprof.Lookup("goroutine")
			jlog.Infof("\n\n==================== one tcp connected >> %s, goroutine count = %d", clientConn.RemoteAddr().String(), p.Count())
		}
	}
}

// 每个连接在此函数中完成所有事情，不再启动多个协程
func handle_logic(clientConn *net.TCPConn) {
	tcpBuffer := make([]byte, nfconst.LEN_TCP_BUFFER)
	clientConn.SetReadBuffer(nfconst.LEN_TCP_BUFFER)
	var handle *nfnet.OrgStreamHandler = nfnet.NewOrgStreamHandler()
	nullmsg := make([]byte, 0)
	var msg []byte
	for {
		n, err := clientConn.Read(tcpBuffer[0:])
		if err != nil {
			if err.Error() == "EOF" {
				// 跳出循环， 并处理已经接收到的数据
				jlog.Infof("%s close, tcp read end, break read loop", clientConn.RemoteAddr().String())
			} else {
				jlog.Error("tcp read error, ", err.Error(), ", break read loop")
			}
			break
		}
		done, buf, err1 := handle.AddStream(tcpBuffer[0:n])
		msg = buf
		if err1 != nil {
			jlog.Error("data add stream error, ", err1.Error())
		} else if done {
			// 处理消息
			msgHandler(clientConn, msg)
			msg = nullmsg
		}
	}

	// 如果执行到了这里， 那么就是 read的时候发生错误到了此处
	// 一个完整的包具备的最小长度为 头2字节，cmd 1字节， 负载长度2字节，校验1字节，尾2字节 = 8字节
	n_msg := len(msg)
	if n_msg >= nfconst.LEN_MIN_PACKAGE && msg[0] == nfconst.SOCK_PACK_HEADER_L && msg[1] == nfconst.SOCK_PACK_HEADER_H &&
		msg[n_msg-2] == nfconst.SOCK_PACK_ENDER_L && msg[n_msg-1] == nfconst.SOCK_PACK_ENDER_H {
		msgHandler(clientConn, msg)
	}
	clientConn.Close()
	jlog.Infof("==================== goroutine exit\n")
	runtime.Goexit()
}

func msgHandler(clientConn *net.TCPConn, msg []byte) {
	// 处理消息
	resp, err2 := nflogic.OnMsessage(msg)
	if err2 != nil {
		jlog.Error("nflogic OnMsessage error, ", err2.Error())
		resp = nflogic.ReturnBadMessage()
	}
	jlog.Debug("Data ready to send...")
	nfutil.PrintHexArray(resp)
	_, err3 := clientConn.Write(resp)
	if err3 != nil {
		jlog.Error("tcp write error, ", err3.Error())
	}
}

/*
func read_tcp_conn(tcpConn *net.TCPConn, inDataChan chan []byte) {
	tcpBuffer := make([]byte, 2048)
	tcpConn.SetReadBuffer(2048)
	var handle *nfnet.OrgStreamHandler = nfnet.NewOrgStreamHandler()
	for {
		n, err := tcpConn.Read(tcpBuffer[0:])
		if err != nil {
			fmt.Println("tcp read error --", err.Error(), ", close tcp conn and in data channel, goroutine exit")
			time.Sleep(1 * time.Second)
			tcpConn.Close()
			close(inDataChan)
			runtime.Goexit()
		} else {
			fmt.Printf("Recieve: %d bytes", n)
			done, buf, err1 := handle.AddStream(tcpBuffer[0:n])
			if err1 != nil {
				fmt.Println("read tcp add stream error", err1)
			} else {
				if done {
					inDataChan <- buf[0:len(buf)]
				}
			}
		}
	}
}

func business_handler(inDataChan chan []byte, outDataChan chan []byte) {
	var resp []byte
	var err error
	for {
		msg, ok := <-inDataChan
		if ok {
			// handle data
			resp, err = nflogic.OnMsessage(msg)
			if err != nil {
				fmt.Println("error:", err.Error())
				resp = nflogic.ReturnBadMessage()
			}
			outDataChan <- resp[0:len(resp)]
		} else {
			fmt.Println("business_handler inDataChan false, close out data channel, goroutine exit")
			close(outDataChan)
			runtime.Goexit()
		}
	}
}

func write_tcp_conn(tcpConn *net.TCPConn, outDataChan chan []byte) {
	for {
		msg, ok := <-outDataChan
		if ok {
			fmt.Print("Send: ")
			nfutil.PrintHexArray(msg)
			tcpConn.Write(msg)
		} else {
			fmt.Println("write_tcp_conn outDataChan false, goroutine exit")
			runtime.Goexit()
		}
	}
}
*/
func check_error(err error) {
	if err != nil {
		jlog.Errorf("Fatal error : %s", err.Error())
	}
}
