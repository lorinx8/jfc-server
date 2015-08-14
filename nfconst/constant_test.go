package nfconst

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	LoadConfig("D:\\Devo\\go\\src\\jfcsrv\\nfconst\\cfg\\j2.cfg")
	fmt.Println(JCfg.JFCPort)
	fmt.Println(JCfg.DbConnString)
	fmt.Println(JCfg.RedisServer)
	fmt.Println(JCfg.RedisPort)
	fmt.Println(JCfg.RedisDbIndex)

}
