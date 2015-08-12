package main

import (
	"fmt"
	"jfcsrv/nflogic"
	"jfcsrv/nfutil"
	"net"
	"runtime"
)

func main() {

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":9630")
	check_error(err)
	tcpListener, err := net.ListenTCP("tcp", tcpAddr)
	check_error(err)
	fmt.Println("JFC Server Start...")
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
			go write_tcp_conn(tcpConn, connChan)
			go read_tcp_conn(tcpConn, connChan)
		}
	}
}

func read_tcp_conn(tcpConn *net.TCPConn, connChan chan []byte) {
	buffer := make([]byte, 2048)
	tcpConn.SetReadBuffer(2048)
	for {
		n, err := tcpConn.Read(buffer[0:])
		if err != nil {
			fmt.Println("one tcp connection read function failed!")
			fmt.Println("one tcp connection close now!")
			tcpConn.Close()
			runtime.Goexit()
		} else {
			fmt.Print("Recieve: ")
			nfutil.PrintHexArray(buffer[0:n])
			connChan <- buffer[0:n]
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
