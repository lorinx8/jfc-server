package nfutil

import "fmt"

func PrintHexArray(arr []byte) {
	fmt.Print("[ ")
	for _, v := range arr {
		fmt.Printf("%02X ", v)
	}
	fmt.Println("]")
}
