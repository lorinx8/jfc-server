package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"runtime"
)

const (
	REQUEST_PARAM    byte = 10
	PARAM_TYPE_ANGLE byte = 1
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
			printHexArray(buffer[0:n])
			connChan <- buffer[0:n]
		}
	}
}

// data detail
//  header  type   len                     data                     cs     end
// 0    1    2    3   4                                            n-2  n-1   n
// F5  A6   0A   00  0D   73 7A 31 32 33 34 35 36 37 38 39 30 01   00   BE   EF
func write_tcp_conn(tcpConn *net.TCPConn, connChan chan []byte) {
	for {
		msg := <-connChan
		// there to handle the data
		// parse the command type
		cmd, data, err := parse_orignal_data(msg)
		if err != nil {
			fmt.Println("error:", err.Error())
			continue
		}
		retData, _ := handle_recive_data(cmd, data)
		fmt.Print("Send: ")
		printHexArray(retData)
		tcpConn.Write(retData)
	}
}

func parse_orignal_data(msg []byte) (cmd byte, data []byte, err error) {
	if msg[0] != 0xF5 || msg[1] != 0xA6 {
		err = errors.New("data invalid, header not correct")
		return
	}
	cmd = msg[2]
	data = msg[5 : len(msg)-3]
	return
}

func handle_recive_data(cmd byte, data []byte) (retData []byte, err error) {
	fmt.Println("Command:", cmd)
	var datab []byte
	switch cmd {
	case REQUEST_PARAM:
		datab, _ = business_handle_request_param(data)
	}
	retData, _ = packageRetData(datab, cmd)
	return
}

func business_handle_request_param(data []byte) (retData []byte, err error) {
	serial := string(data[0 : len(data)-1])
	paramType := data[len(data)-1]
	fmt.Println("Handle request param, Serail:", serial, ", ParamType:", paramType)

	var datab []byte
	switch paramType {
	case PARAM_TYPE_ANGLE:
		datab, _ = business_get_angles_param(serial)
	}

	retData = data
	retData = append(retData, datab...)
	return
}

//  uint8    uint8    uint8   uint8    uint8    uint8   uint16   uint16  uint16  uint16
// 角度2个数 角度2编号 角度2数值 角度1个数 角度1编号 角度1数值 裁剪X坐标 裁剪Y坐标 裁剪宽度 裁剪高度
func business_get_angles_param(serial string) (data []byte, err error) {
	// temp data
	b_buf := bytes.NewBuffer([]byte{})
	binary.Write(b_buf, binary.BigEndian, []uint8{2, 1, 20, 2})
	binary.Write(b_buf, binary.BigEndian, []uint8{1, 40})
	binary.Write(b_buf, binary.BigEndian, []uint16{320, 320, 740, 360})
	binary.Write(b_buf, binary.BigEndian, []uint8{2, 80})
	binary.Write(b_buf, binary.BigEndian, []uint16{320, 320, 740, 360})
	binary.Write(b_buf, binary.BigEndian, []uint8{2, 70, 1})
	binary.Write(b_buf, binary.BigEndian, []uint8{1, 50})
	binary.Write(b_buf, binary.BigEndian, []uint16{300, 300, 740, 360})
	return b_buf.Bytes(), nil
}

func check_error(err error) {
	if err != nil {
		fmt.Printf("Fatal error : %s", err.Error())
	}
}

func packageRetData(data []byte, cmd byte) (retData []byte, err error) {
	retData = make([]byte, 5, 128)
	retData[0] = 0xF5
	retData[1] = 0xA6
	retData[2] = cmd
	datalen := len(data)
	retData[3] = (byte)(datalen >> 8 & 0xFF)
	retData[4] = (byte)(datalen & 0xFF)
	retData = append(retData, data...)
	retData = append(retData, 0, 0xBE, 0xEF)
	return
}

func printHexArray(arr []byte) {
	fmt.Print("[ ")
	for _, v := range arr {
		fmt.Printf("%02X ", v)
	}
	fmt.Println("]")
}
