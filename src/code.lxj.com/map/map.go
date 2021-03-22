package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	/*
		scoreMap := make(map[string]int, 8)
		scoreMap["张三"] = 90
		scoreMap["小明"] = 100
		fmt.Println(scoreMap)
		fmt.Println(scoreMap["小明"])
		fmt.Printf("type of a:%T\n", scoreMap)
	*/
	/*
		userInfo := map[string]string{
			"username": "沙河小王子",
			"password": "123456",
		}
		fmt.Println(userInfo)
	*/
	/*
		scoreMap := make(map[string]int, 8)
		scoreMap["张三"] = 90
		scoreMap["小明"] = 100
		v, ok := scoreMap["张三"]
		if ok {
			fmt.Println(v)
		} else {
			fmt.Println("查无此人")
		}
	*/
	/*
		scoreMap := make(map[string]int, 8)
		scoreMap["张三"] = 90
		scoreMap["小明"] = 100
		scoreMap["娜扎"] = 60
		for _, v := range scoreMap {
			fmt.Println(v)
		}
	*/
	/*
		scoreMap := make(map[string]int, 8)
		scoreMap["张三"] = 90
		scoreMap["小明"] = 100
		scoreMap["娜扎"] = 60
		delete(scoreMap, "小明")
		for k, v := range scoreMap {
			fmt.Println(k, v)
		}
	*/
	/*
		rand.Seed(time.Now().UnixNano())
		var scoreMap = make(map[string]int, 200)
		for i := 0; i < 100; i++ {
			key := fmt.Sprintf("stu%02d", i)
			value := rand.Intn(100)
			scoreMap[key] = value
		}
		var keys = make([]string, 0, 200)
		for key := range scoreMap {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		for _, key := range keys {
			fmt.Println(key, scoreMap[key])
		}
	*/
	/*
		var mapSlice = make([]map[string]string,3)
		for index,value := range mapSlice{
			fmt.Printf("index:%d value:%v\n", index, value)
		}
		fmt.Println("after init")
		mapSlice[0] = make(map[string]string,10)
		mapSlice[0]["name"] = "小王子"
		mapSlice[0]["password"] = "123456"
		mapSlice[0]["address"] = "沙河"
		for index,value := range mapSlice{
			fmt.Printf("index:%d value:%v\n", index, value)
		}
	*/
	/*
		var sliceMap = make(map[string][]string,3)
		fmt1.Println(sliceMap)
		fmt1.Println("after init")
		key := "中国"
		value,ok := sliceMap[key]
		fmt1.Println(value,ok)
		if !ok {
			value = make([]string,0,2)
		}
		value = append(value,"北京","上海")
		sliceMap[key] = value
		fmt1.Println(sliceMap)
	*/
	/*
		str := "how do you do think you about"
		strSlice := strings.Split(str," ")
		fmt1.Println(strSlice)

		coutMap := make(map[string]int,10)
		for _,key := range strSlice{
			_,isReal := coutMap[key]
			if !isReal{
				coutMap[key] = 1
			}else {
				coutMap[key] +=1
			}
		}
		fmt1.Println(coutMap)
	*/
	/*
		type Map map[string][]int
		m := make(Map)
		s := []int{1, 2}
		s = append(s, 3)
		fmt.Printf("%+v\n", s)
		m["q1mi"] = s
		s = append(s[:1], s[2:]...)
		fmt.Printf("%+v\n", s)
		fmt.Printf("%+v\n", m["q1mi"])
		}
	*/
	/*
		var ary1 = [...]int{1, 3, 5, 7, 8}
		for i :=0;i<len(ary1);i++{
			for j := i+1;j < len(ary1);j++{
				if ary1[i] + ary1[j] == 8{
					fmt1.Println(i,j)
				}
			}
		}
	*/
	var scoreMap = make(map[string]int, 200)
	for i := 0; i < 20; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		scoreMap[key] = value
	}

	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}
