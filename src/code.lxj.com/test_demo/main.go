package main

import (
	"fmt"
	"os"
	"runtime"
	"syscall"
	"time"
)

func main() {
	a := GetFileCreateTime("./main.go")
	//if a.Year() == time.Now().Year() && a.Month() == time.Now().Month() && a.Day() == time.Now().Day() {
	fmt.Println(a.Hour())
	b := time.Now().Add(time.Hour)
	fmt.Println(time.Now().Format("20060102150405"))
	fmt.Println(b.Hour())
	//}
}

func GetFileCreateTime(path string) time.Time {
	osType := runtime.GOOS
	fileInfo, _ := os.Stat(path)
	if osType == "windows" {
		wFileSys := fileInfo.Sys().(*syscall.Win32FileAttributeData)
		tNanSeconds := wFileSys.CreationTime.Nanoseconds() /// 返回的是纳秒
		tSec := tNanSeconds / 1e9                          ///秒
		return time.Unix(tSec, 0)
	}
	return time.Now()
}
