package main

import (
	"fmt"
	"io"
	"os"
)

/*
#Write和WriteString

func main() {
	file,err := os.OpenFile("write.txt",os.O_CREATE|os.O_TRUNC|os.O_WRONLY,0666)
	if err != nil{
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	str := "hello 土豆"
	file.Write([]byte(str))
	file.WriteString("hello 小王子")
}
*/
/*
#bufio.NewWriter
func main() {
	file, err := os.OpenFile("write.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("hello 土豆呼叫萝卜\n")
	}
	writer.Flush()
}
*/
/*
func main() {
	str:="hello lxj\n"
	err := ioutil.WriteFile("write.txt",[]byte(str),0666)
	if err != nil{
		fmt.Println("write file failed,err:",err)
		return
	}
}
*/
// CopyFile 拷贝文件函数
func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		fmt.Println("open file failed,err:", err)
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open %s failed,err:%v", dst, err)
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

func main() {
	_, err := CopyFile("dst.txt", "src.txt")
	if err != nil {
		fmt.Println("copy file failed, err:", err)
		return
	}
	fmt.Println("copy done")
}
