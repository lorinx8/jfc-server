package main

import (
	"fmt"
	"jfcsrv/nfconst"
	"jfcsrv/nflogic"
	"jfcsrv/nfnet"
	"jfcsrv/nfutil"
	"net"
	"runtime"
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
			connChan := make(chan []byte)
			go read_tcp_conn(tcpConn, connChan)
			go write_tcp_conn(tcpConn, connChan)

		}
	}
}

func read_tcp_conn(tcpConn *net.TCPConn, connChan chan []byte) {
	buffer := make([]byte, 2048)
	tcpConn.SetReadBuffer(2048)
	var handle *nfnet.OrgStreamHandler = nfnet.NewOrgStreamHandler()
	for {
		n, err := tcpConn.Read(buffer[0:])
		if err != nil {
			fmt.Println(err.Error())
			tcpConn.Close()
			runtime.Goexit()
		} else {
			fmt.Printf("Recieve: %d bytes", n)
			done, buf, err1 := handle.AddStream(buffer[0:n])
			if err1 != nil {
				fmt.Println("read tcp add stream error", err1)
			} else {
				if done {
					connChan <- buf[0:len(buf)]
				}
			}
		}
	}
}

func write_tcp_conn(tcpConn *net.TCPConn, connChan chan []byte) {
	for {
		msg := <-connChan

		// handle data
		resp, err := nflogic.OnMsessage(msg)
		if err != nil {
			fmt.Println("error:", err.Error())
			continue
		} else {
			fmt.Print("Send: ")
			nfutil.PrintHexArray(resp)
			tcpConn.Write(resp)
		}
	}
}

func check_error(err error) {
	if err != nil {
		fmt.Printf("Fatal error : %s", err.Error())
	}
}
