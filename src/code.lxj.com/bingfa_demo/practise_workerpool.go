package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
1使用goroutine和channel实现一个计算 int64 随机数各位数和的程序。
2开启一个goroutine循环生成 int64 类型的随机数，发送到jobChan
3开启 24 个goroutine从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
4主goroutine从resultChan取出结果并打印到终端输出
*/

type job struct {
	value int64
}

type result struct {
	job *job
	sum int64
}

//初始化
var jobChan = make(chan *job, 100)
var resultChan = make(chan *result, 100)
var wg3 sync.WaitGroup

//生成一个int64位的随机数
// rn:只写
func randNumber(rn chan<- *job) {
	wg3.Done()
	//循环生成int64类型的随机数,发送到jobChan
	for {
		//获取一个int64位的随机值
		x := rand.Int63()
		newJob := &job{
			value: x,
		}
		//rn读取newJob的值
		rn <- newJob
		time.Sleep(time.Millisecond * 500)
	}
}

//rn:只读,resultChan:只读
func worker(rn <-chan *job, resultChan chan<- *result) {
	wg3.Done()
	//从jobChan中取出随机数计算各位数的和,将结果发送到resultChan
	for {
		//rn的值发送到job里
		job := <-rn
		sum := int64(0)
		n := job.value
		for n > 0 {
			sum += n % 10
			n = n / 10
		}
		newResult := &result{
			job: job,
			sum: sum,
		}
		resultChan <- newResult
	}
}
func main() {
	wg3.Add(1)
	go randNumber(jobChan)
	//开启24个goroutine执行worker
	wg3.Add(24)
	for i := 0; i < 24; i++ {
		go worker(jobChan, resultChan)
	}
	//主goroutine从resultChan取出结果并打印到终端输出
	for result := range resultChan {
		fmt.Printf("value:%d  sum:%d\n", result.job.value, result.sum)
	}
	wg3.Wait()
}
