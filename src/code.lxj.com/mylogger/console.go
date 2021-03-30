package mylogger

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

//在终端写日志相关内容
type LogLevel uint16

//常量
const (
	Unknown LogLevel = iota
	DEBUG
	INFO
	WARNING
	ERROR
	FATAL
)

//日志开关
func parseLogLevel(s string) (LogLevel, error) {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("无效得日志级别")
		return Unknown, err
	}
}

//Logger日志结构体
type Logger struct {
	Level LogLevel
}

//NewLog构造函数
func NewLog(levelStr string) Logger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return Logger{
		Level: level,
	}
}

//判断日志级别
func (l Logger) enable(LogLevel LogLevel) bool {
	return l.Level <= LogLevel
}

//输出到日志
func (l Logger) Debug(msg string) {
	if l.enable(DEBUG) {
		now := time.Now()
		fmt.Printf("[%s] [DEBUG] %s\n", now.Format("2006-01-02 15:04:05"), msg)
	}
}

func (l Logger) Info(msg string) {
	if l.enable(INFO) {
		now := time.Now()
		fmt.Printf("[%s] [INFO] %s\n", now.Format("2006-01-02 15:04:05"), msg)
	}
}

func (l Logger) Warning(msg string) {
	if l.enable(WARNING) {
		now := time.Now()
		fmt.Printf("[%s] [WARNING] %s\n", now.Format("2006-01-02 15:04:05"), msg)
	}
}

func (l Logger) Error(msg string) {
	if l.enable(ERROR) {
		now := time.Now()
		fmt.Printf("[%s] [ERROR] %s\n", now.Format("2006-01-02 15:04:05"), msg)
	}
}

func (l Logger) Fatal(msg string) {
	if l.enable(FATAL) {
		now := time.Now()
		fmt.Printf("[%s] [FATAL] %s\n", now.Format("2006-01-02 15:04:05"), msg)
	}
}
