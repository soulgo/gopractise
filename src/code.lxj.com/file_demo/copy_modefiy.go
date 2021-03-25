package main

import (
	"fmt"
	"io"
	"os"
)

/*
   在文件中插入相关内容
*/
func copy_modeify_demo() {
	//打开要操作的文件：src.txt
	fileObj, err := os.OpenFile("src.txt", os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("open file failed,err:%v\n", err)
		return
	}
	//因为没有办法直接在文件中插入内容，所以要借助一个临时文件
	tmpFile, err := os.OpenFile("sb.tmp", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("create tmp file failed,err:%v\n", err)
		return
	}
	//读取源文件写入临时文件
	var ret [1]byte
	n, err := fileObj.Read(ret[:])
	if err != nil {
		fmt.Printf("read from failed,err:%v", err)
		return
	}
	//写入临时文件
	tmpFile.Write(ret[:n])
	//写入要插入的内容
	var s []byte
	s = []byte{'c'}
	tmpFile.Write(s)
	//紧接着把源文件后续的内容写入临时文件
	var x [1024]byte
	for {
		n, err := fileObj.Read(x[:])
		if err == io.EOF {
			tmpFile.Write(x[:n])
			break
		}
		if err != nil {
			fmt.Printf("read from file failed,err:%v\n", err)
			return
		}
		//源文件后续的也写入临时文件中
		tmpFile.Write(x[:n])
	}
	fileObj.Close()
	tmpFile.Close()
	//删除临时文件（更改名称覆盖）
	os.Rename("sb.tmp", "src.txt")
}

func main() {
	copy_modeify_demo()
}
