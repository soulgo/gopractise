package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

//往文件里写日志相关内容，日志文件结构体
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

//判断日志级别，判断是否需要记录该日志
func (f *FileLogger) enable(LogLevel LogLevel) bool {
	return f.Level <= LogLevel
}

//判断文件大小,是否需要切割
func (f *FileLogger) checkSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed,err:%v\n", err)
		return false
	}
	//如果当前文件得大小大于最大尺寸就返回true
	return fileInfo.Size() >= f.maxFileSize
}

//根据大小切割日志
func (f *FileLogger) splitSize(file *os.File) (*os.File, error) {
	//需要切割日志文件
	nowStr := time.Now().Format("20060102150405000")
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed,err:%v\n", err)
		return nil, err
	}
	logNmae := path.Join(f.filePath, fileInfo.Name())
	newLogName := fmt.Sprintf("%s.%s", logNmae, nowStr)
	//1、关闭当前得日志文件
	file.Close()
	//2、备份旧文件并且根据时间重命名
	os.Rename(logNmae, newLogName)
	//3、打开一个新的日志文件
	fileObj, err := os.OpenFile(logNmae, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open new log file failed,err:%v\n", err)
		return nil, err
	}
	//4、将打开得新日志文件对象赋值给f.fileObj
	return fileObj, nil
}

//记录日志得方法，输出日志格式，error及以上得日志输出再输出到error日志里
func (f *FileLogger) log(lv LogLevel, format string, a ...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		//判断日志级别，如果是error级别以上得日志就单独记录
		if lv >= ERROR {
			if f.checkSize(f.errFileObj) {
				newFile, err := f.splitSize(f.errFileObj)
				if err != nil {
					return
				}
				f.errFileObj = newFile
			}
			fmt.Fprintf(f.errFileObj, "[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), funcName, fileName, lineNo, msg)
		} else {
			if f.checkSize(f.fileObj) {
				newFile, err := f.splitSize(f.fileObj)
				if err != nil {
					return
				}
				f.fileObj = newFile
			}
			fmt.Fprintf(f.fileObj, "[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), funcName, fileName, lineNo, msg)
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
