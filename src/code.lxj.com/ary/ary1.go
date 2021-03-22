package main

import (
	"fmt"
	// "sort"
)

func main() {
	// s := []int{1, 3, 5}
	// for i := 0; i < len(s); i++ {
	// 	fmt.Println(i, s[i])
	// }

	// for index, value := range s {
	// 	fmt.Println(index, value)
	// }
	// var s []int
	// s = append(s, 1)
	// s = append(s, 2, 3, 4)
	// s2 := []int{5, 6, 7}
	// s = append(s, s2...)
	// fmt.Println(s)
	// var numSlice []int
	// for i := 0; i < 10; i++ {
	// 	numSlice = append(numSlice, i)
	// 	fmt.Printf("%v  len:%d  cap:%d  ptr:%p\n", numSlice, len(numSlice), cap(numSlice), numSlice)
	// }
	// var citySlice []string
	// // 追加一个元素
	// citySlice = append(citySlice, "北京")
	// // 追加多个元素
	// citySlice = append(citySlice, "上海", "广州", "深圳")
	// // 追加切片
	// a := []string{"成都", "重庆"}
	// citySlice = append(citySlice, a...)
	// fmt.Println(citySlice) //[北京 上海 广州 深圳 成都 重庆]
	// a := []int{30, 31, 32, 33, 34, 35, 36, 37}
	// a = append(a[:2], a[3:]...)
	// fmt.Println(a)
	// var a = make([]string, 5, 10)
	// for i := 0; i < 10; i++ {
	// 	a = append(a, fmt.Sprintf("%v", i))
	// }
	// fmt.Println(a)
	// var a = [...]int{3, 7, 8, 9, 1}
	// b := a[:]
	// // sort.Ints(b)
	// sort.Sort(sort.Reverse(sort.IntSlice(b)))
	// fmt.Println(a)
	// var a = []int{3, 7, 8, 9, 1}
	// sort.Ints(a)
	// fmt.Println(a)
	var a = [3]int{1, 2, 3}
	a1 := [3]int{1, 2, 3}
	fmt.Println(a)
	fmt.Println(a1)
}
