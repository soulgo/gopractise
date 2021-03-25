package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func readByBufio() {
	file, err := os.Open("main.go")
	if err != nil {
		fmt.Printf("open file failed ,err=\n", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			if len(line) != 0 {
				fmt.Println(line)
			}
			fmt.Println("文件读完了")
			return
		}
		if err != nil {
			fmt.Print("read file by bufio failed,err:%v\n", err)
			return
		}
		fmt.Printf(line)
	}

}
func main() {
	readByBufio()
}
