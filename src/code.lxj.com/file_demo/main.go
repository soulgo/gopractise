package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file failed,err=", err)
		return
	}
	file.Close()
}

/*
func main() {
	/*
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file failed")
		return
	}
	defer file.Close()
	for {
		var tmp = make([]byte, 128)
		n, err := file.Read(tmp)
		if err == io.EOF {
			fmt.Print("文件读完了")
			return
		}
		if err != nil {
			fmt.Println("read file failed,err=", err)
			return
		}
		fmt.Printf("读取了%d字节数据\n", n)
		fmt.Printf(string(tmp[:n]))
	}
}
*/
