package main

import (
	"code.lxj.com/mylogger_time"
	"fmt"
	"time"
)

//调用接口：声明一个全局的接口变量
//var log mylogger.Logger
/*
func main() {
	//终端日志实例
	//log = mylogger.NewConsoleLog("error")
	//文件日志实例
	log := mylogger.NewFileLogger("error", "./", "lxj.log", 10*1024)
	for {
		log.Debug("这是一条debug日志")
		log.Info("这是一条Info日志")
		log.Warning("这是一条Warning日志")
		a := 100
		b := "理想"
		c := "北京大学"
		log.Error("这是一条Error日志，id:%d,name:%s,address:%s", a, b, c)
		log.Fatal("这是一条Fatal日志")
		time.Sleep(time.Second * 2)
	}

}
*/
func main() {
	log := mylogger_time.NewFileLogger("info")
	err := log.FileInit(true, "./", "detail.log")
	if err != nil {
		fmt.Printf("日志文件初始化失败,%s", err)
		return
	}

	for {
		log.Debug("这是一条debug日志")
		log.Info("这是一条Info日志")
		//time.Sleep(time.Millisecond * 300)
		time.Sleep(time.Second)
	}
}
