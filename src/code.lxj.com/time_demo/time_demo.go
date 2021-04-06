package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		//定时器
		ticker := time.Tick(time.Second)
		for i:= range ticker{
			fmt.Println(i)
		}
	*/
	/*
		now := time.Now()
		fmt.Println(time.Now())
		fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
		fmt.Println(now.Format("2006/01/02 15:04 PM"))
		fmt.Println(now.Format("15:04 2006/01/02"))
	*/
	/*
		now := time.Now()
		loc,err := time.LoadLocation("Asia/Shanghai")
		if err != nil{
			fmt.Println(err)
			return
		}

		timeObj,err := time.ParseInLocation("2006/01/02 15:04:05", "2019/08/04 14:15:20", loc)
		if err != nil{
			fmt.Println(err)
			return
		}
		fmt.Println(timeObj)
		fmt.Println(timeObj.Sub(now))
	*/
	/*
		now := time.Now()
		timenow :=now.Format("2006/01/02 15:04:05")
		fmt.Println(timenow)
	*/
	/*
		start := time.Now()
		now := time.Now().Format("2006/01/02 15:04:05")
		fmt.Println(now)
		end := time.Now()
		last := end.Sub(start)
		fmt.Println(start.Add(time.Hour*2))
		fmt.Println(last)
	*/
	now := time.Now()
	m2 := now.Format("2006-01-02 15:04:00")
	m, _ := time.ParseDuration("1h")
	fmt.Println(m)
	fmt.Println(m2)
	m1 := now.Add(m)
	fmt.Println(m1)
	//m1 := now.Add(m)
	//m3 := m1.Format("2006-01-02 15:04:00")
	//t1, _ := time.ParseInLocation("2006-01-02 15:04:05", m2, time.Local)
	//t2, _ := time.ParseInLocation("2006-01-02 15:04:05", m3, time.Local)
	/*
		fmt.Println(t1)
		if t1.Equal(t2){
			fmt.Println("相等")
		}else {
			fmt.Println("不相等")
		}
	*/
	//fmt.Println(t1.After(t2))

}
