package main

import (
	"code.lxj.com/mylogger"
)

//调用接口
var log mylogger.Logger

func main() {
	//log = mylogger.NewConsoleLog("error")
	log := mylogger.NewFileLogger("info", "./", "lxj.log", 10*1024)

	for {
		log.Debug("这是一条debug日志")
		log.Info("这是一条Info日志")
		log.Warning("这是一条Warning日志")
		a := 100
		b := "理想"
		c := "北京大学"
		log.Error("这是一条Error日志，id:%d,name:%s,address:%s", a, b, c)
		log.Fatal("这是一条Fatal日志")
		//time.Sleep(time.Second * 2)
	}
}
