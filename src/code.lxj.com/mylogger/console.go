package mylogger

import (
	"fmt"
	"time"
)

//Logger日志结构体
type consoleLogger struct {
	Level LogLevel
}

//NewLog构造函数
func NewLog(levelStr string) consoleLogger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return consoleLogger{
		Level: level,
	}
}

//判断日志级别
func (c consoleLogger) enable(LogLevel LogLevel) bool {
	return c.Level <= LogLevel
}

//提取出公共得部分,增加一个可写可不写得接口，这个参数可以有，也可以没有
func (c consoleLogger) log(lv LogLevel, format string, a ...interface{}) {
	if c.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), funcName, fileName, lineNo, msg)
	}
}

//输出到日志
func (c consoleLogger) Debug(format string, a ...interface{}) {
	c.log(DEBUG, format, a...)
}

func (c consoleLogger) Info(format string, a ...interface{}) {
	c.log(INFO, format, a...)
}

func (c consoleLogger) Warning(format string, a ...interface{}) {
	c.log(WARNING, format, a...)
}

func (c consoleLogger) Error(format string, a ...interface{}) {
	c.log(ERROR, format, a...)
}

func (c consoleLogger) Fatal(format string, a ...interface{}) {
	c.log(FATAL, format, a...)
}
