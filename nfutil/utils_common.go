package nfutil

import (
	"bufio"
	"bytes"
	"fmt"
	"jfcsrv/nflog"
	"os"
	"time"
)

var jlog = nflog.Logger

func PrintHexArray(arr []byte) {
	writer := new(bytes.Buffer)
	writer.WriteString("[ ")
	for _, v := range arr {
		str := fmt.Sprintf("%02X ", v)
		writer.WriteString(str)
	}
	writer.WriteString(" ]")
	jlog.Debug(writer.String())
}

func WriteLocalFile(filename string, data []byte) (n int, err error) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0777)
	if err != nil {
		return 0, err
	}
	writer := bufio.NewWriter(file)
	nn, err1 := writer.Write(data)
	writer.Flush()
	defer file.Close()
	return nn, err1
}

func GetNow() (y int, mon int, d int, h int, min int, s int) {
	t := time.Now()
	return t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute(), t.Second()
}

func GetNowString() (str string) {
	t := time.Now()
	str = fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute(), t.Second())
	return
}

// 创建文件路径,如果不存在
func CreateLocalFolderIfNotExist(path string) error {
	err := os.MkdirAll(path, 0777)
	return err
}
