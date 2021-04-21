package main

import (
	"fmt"
	"time"
)

//jobs:只读;results:只写
func worker1(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker:%d start job:%d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job:%d\n", id, j)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	//启动3个goroutine
	for w := 1; w <= 3; w++ {
		go worker1(w, jobs, results)
	}
	//5个任务
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)
	//输出结果
	for a := 1; a <= 5; a++ {
		<-results
	}
}
