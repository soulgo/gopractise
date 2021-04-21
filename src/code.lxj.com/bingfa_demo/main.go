package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg1 sync.WaitGroup

func hello(i int) {
	defer wg1.Done()
	time.Sleep(time.Second * time.Duration(3))
	fmt.Println("Hello----->", i)
}
func hello1(i int) {
	defer wg1.Done()
	time.Sleep(time.Second * time.Duration(3))
	fmt.Println("Hello1----->", i)
}
func main() {
	runtime.GOMAXPROCS(4)
	for i := 0; i < 100; i++ {
		wg1.Add(2)
		go hello(i)
		go hello1(i)
		time.Sleep(time.Second)
	}
	wg1.Wait()
}
