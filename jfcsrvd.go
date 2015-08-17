package main

import (
	"fmt"
	"jfcsrv/nfconst"
	"jfcsrv/nflogic"
	"jfcsrv/nfnet"
	"jfcsrv/nfutil"
	"net"
	"runtime"
	"runtime/pprof"
)

func main() {
	fmt.Println("JFC Server For Papakaka...")
	fmt.Println("Config loading...")
	nfconst.InitialConst()
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":"+nfconst.JCfg.JFCPort)
	check_error(err)
	tcpListener, err := net.ListenTCP("tcp", tcpAddr)
	check_error(err)
	fmt.Println("JFC Server listening on port", nfconst.JCfg.JFCPort)
	for i := 0; i < 100; i++ {
		go handle_tcp_accept(tcpListener)
	}
	select {}
}

func handle_tcp_accept(tcpListener *net.TCPListener) {
	for {
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			fmt.Println("tcp accept failed!")
			continue
		} else {
			fmt.Println("one tcp connected!")

			inDataChan := make(chan []byte)
			outDataChan := make(chan []byte)

			go read_tcp_conn(tcpConn, inDataChan)
			go business_handler(inDataChan, outDataChan)
			go write_tcp_conn(tcpConn, outDataChan)

			p := pprof.Lookup("goroutine")
			fmt.Println(p.Count())
		}
	}
}

func read_tcp_conn(tcpConn *net.TCPConn, inDataChan chan []byte) {
	tcpBuffer := make([]byte, 2048)
	tcpConn.SetReadBuffer(2048)
	var handle *nfnet.OrgStreamHandler = nfnet.NewOrgStreamHandler()
	for {
		n, err := tcpConn.Read(tcpBuffer[0:])
		if err != nil {
			fmt.Println("tcp read error --", err.Error(), ", close tcp conn and in data channel, goroutine exit")
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

func check_error(err error) {
	if err != nil {
		fmt.Printf("Fatal error : %s", err.Error())
	}
}
