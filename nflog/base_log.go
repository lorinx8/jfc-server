package nflog

import (
	"fmt"

	seelog "github.com/cihub/seelog"
)

var (
	//logger *log.Logger
	Logger seelog.LoggerInterface
)

func loadAppConfig() {
	appConfig := `
<seelog minlevel="info">
    <outputs formatid="common">
        <rollingfile type="size" filename="log/jcar.log" maxsize="500000000" maxrolls="5"/>	
		<filter levels="critical,error">
			<rollingfile type="size" filename="log/jcar.wrong.log" maxsize="50000000" maxrolls="5" formatid="critical"/>	
        </filter>
    </outputs>
    <formats>
        <format id="common" format="%Date %Time [%LEV] %Func:%Line : %Msg%n" />
        <format id="critical" format="%Date %Time [%LEV] %File %Func:%Line %Msg%n" />
		<format id="trace" format="%Date %Time %Msg%n" />
        <format id="criticalemail" format="Critical error on our server!\n    %Time %Date %RelFile %Func %Msg \nSent by Seelog"/>
    </formats>
</seelog>
`
	logger, err := seelog.LoggerFromConfigAsBytes([]byte(appConfig))
	if err != nil {
		fmt.Println(err)
		return
	}
	UseLogger(logger)
}
func init() {
	DisableLog()
	loadAppConfig()
}

// DisableLog disables all library log output
func DisableLog() {
	Logger = seelog.Disabled
}

// UseLogger uses a specified seelog.LoggerInterface to output library log.
// Use this func if you are using Seelog logging system in your app.
func UseLogger(newLogger seelog.LoggerInterface) {
	Logger = newLogger
}

/*
func GetLogger() (l *log.Logger) {
	if logger != nil {
		return logger
	} else {
		e, err := os.OpenFile("j2srv.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
		if err != nil {
			fmt.Println(err)
			logger = log.New(os.Stdout, "\r\n", log.Lshortfile|log.Ldate|log.Ltime)
		} else {
			logger = log.New(e, "\r\n", log.Lshortfile|log.Ldate|log.Ltime)
		}
	}
	return logger
}
*/
