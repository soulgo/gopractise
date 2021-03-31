package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

//往文件里写日志相关内容

type FileLogger struct {
	Level       LogLevel
	filePath    string //日志文件保存得路径
	fileName    string //日志文件保存得文件名
	fileObj     *os.File
	errFileObj  *os.File
	maxFileSize int64
}

//NewFileLogger 构造函数
func NewFileLogger(levelStr, fp, fn string, maxSize int64) *FileLogger {
	logLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	f1 := &FileLogger{
		Level:       logLevel,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
	}
	err = f1.initFile() //按照文件路径和文件名称将文件打开
	if err != nil {
		panic(err)
	}
	return f1
}

//初始化日志文件
func (f *FileLogger) initFile() error {
	fullName := path.Join(f.filePath, f.fileName)
	//正常日志
	fileObj, err := os.OpenFile(fullName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed,err:%v\n", err)
		return err
	}
	//错误日志
	errFileObj, err := os.OpenFile(fullName+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed,err:%v\n", err)
		return err
	}
	//日志文件都已经打开了
	f.fileObj = fileObj
	f.errFileObj = errFileObj
	return nil
}

//判断日志级别
func (f *FileLogger) enable(LogLevel LogLevel) bool {
	return f.Level <= LogLevel
}

//输出日志格式，error及以上得日志输出再输出到error日志里
func (f *FileLogger) log(lv LogLevel, format string, a ...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		fmt.Fprintf(f.fileObj, "[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), funcName, fileName, lineNo, msg)
		if lv >= ERROR {
			//如果要记录得日志大于等于ERROR级别，我还要在err日志中在记录一遍
			fmt.Fprintf(f.errFileObj, "[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), funcName, fileName, lineNo, msg)
		}
	}
}

//输出到日志
func (f *FileLogger) Debug(format string, a ...interface{}) {
	f.log(DEBUG, format, a...)
}

func (f *FileLogger) Info(format string, a ...interface{}) {
	f.log(INFO, format, a...)
}

func (f *FileLogger) Warning(format string, a ...interface{}) {
	f.log(WARNING, format, a...)
}

func (f *FileLogger) Error(format string, a ...interface{}) {
	f.log(ERROR, format, a...)
}

func (f *FileLogger) Fatal(format string, a ...interface{}) {
	f.log(FATAL, format, a...)
}

//文件关闭
func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}
