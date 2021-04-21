package mylogger_time

import (
	"fmt"
	"os"
	"path"
	"strings"
	"time"
)

/*
	根据时间切割日志
*/

// FileLogger 定义结构体
type FileLogger struct {
	level    logLevel
	isFile   bool
	filePath string
	fileName string
	fileObj  *os.File
}

// NewFileLogger 构造函数
func NewFileLogger(level string) *FileLogger {
	level = strings.ToUpper(level)
	Level, err := parseLogLevel(level)
	if err != nil {
		fmt.Printf("解析级别出现错误，err:%v\n", err)
	}
	return &FileLogger{
		level:  Level,
		isFile: false,
	}
}

// FileInit 如用户选择输出至文件，初始化文件
func (c *FileLogger) FileInit(isFile bool, filePath, fileName string) error {
	c.isFile = isFile
	c.filePath = filePath
	c.fileName = fileName
	//fmt.Println(c.isFile,isFile,filePath) //为毛不打印
	logPath := path.Join(filePath, fileName)
	fileObj, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open logpath failed,err:%v \n", err)
		return err
	}
	c.fileObj = fileObj
	return nil
}

//日志开关，确定当前日志级别是否可以打印，返回true则可以打印
func (c *FileLogger) isTrue(cLevel logLevel) bool {
	return cLevel >= c.level
}

//定义日志格式，打印日志到文件或屏幕
func (c *FileLogger) messageFormat(lv logLevel, message string) {
	now := time.Now()
	funcName, fileName, lineNo := getInfo(3)
	//定义日志格式
	msg := fmt.Sprintf("[%s] [%v] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"),
		getLogString(lv), fileName, funcName, lineNo, message)
	if c.isTrue(lv) {
		//fmt.Printf(msg) //输出至屏幕
		if c.isFile {
			//检查时间，如果时间符合要求，那就切割
			if c.checkTime() {
				c.MoveFile()
			}
		}
		_, err := fmt.Fprintf(c.fileObj, msg)
		if err != nil {
			fmt.Println("err:", err)
		}
	}
}

// checkTime 检查当前时间是否为每分钟的00秒
func (c *FileLogger) checkTime() bool {
	/*
		timef := time.Now().Format("200601021504") + "00"
		checkTime, err := strconv.Atoi(timef)
		nowTime, err := strconv.Atoi(time.Now().Format("20060102150405"))
		if err != nil {
		 	fmt.Printf("转换错误：%s", err)
		 }
		 //fmt.Printf("时间 %v %T,标志时间:%v %T \n",nowTime,nowTime,checkTime,checkTime)
		 return checkTime == nowTime
	*/
	//当秒数为00的时候切割
	return time.Now().Format("05") == "00"

}

// MoveFile 切割日志文件
func (c *FileLogger) MoveFile() {
	c.fileObj.Close()
	nowStr := time.Now().Format("20060102150405")
	logPath := path.Join(c.filePath, c.fileName)
	newLogPath := logPath + nowStr
	os.Rename(logPath, newLogPath)
	fileObj, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("打开新的日志文件错误,%s", err)
		return
	}
	c.fileObj = fileObj
}

// Debug ...
func (c *FileLogger) Debug(message string) {
	c.messageFormat(DEBUG, message)
}

// Info ...
func (c *FileLogger) Info(message string) {
	c.messageFormat(INFO, message)
}
