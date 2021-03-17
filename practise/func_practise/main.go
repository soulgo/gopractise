package main

import (
	fmt1 "fmt"
)

/*
func intSum(x int, y int) int {
	return  x + y
}
 */
/*
func intSum2(x ...int) int {
	fmt1.Println(x)
	sum := 0
	for _,v := range x{
		sum += v
	}
	return sum
}

func main() {
	ret1 := intSum2()
	ret2 := intSum2(10)
	ret3 := intSum2(10,20)
	fmt1.Println(ret1,ret2,ret3)
}
 */
//全局变量
/*
var num int64 = 10
func testGlobalVar(){
	fmt.Printf("num=%d\n",num)
}
func main(){
	testGlobalVar()
}
 */
/*
type calculation func(int,int) int

func add(x, y int) int {
	return x + y
}

func main() {
	var c calculation
	c = add
	fmt1.Println(c(1,2))
}
 */
/*
func add(x, y int) int {
	return x + y
}
func calc(x, y int, ab func(int, int)int) int {
	return ab(x,y)
}
func main() {
	ret2 := calc(10,20,add)
	fmt1.Println(ret2)
}
 */
/*
func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	fmt1.Println("=======")
	return x - y
}

func do(s string) (func(x,y int) int,error) {
	switch s {
	case "+":
		return add,nil
	case "-":
		return sub,nil
	default:
		err := errors.New("无法识别的操作符")
		return nil, err
	}
}

func main() {
	fmt1.Println(do("-"))
}
 */

/*
func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}
func main() {
	var f = adder()
	fmt1.Println(f(10))
	fmt1.Println(f(20))
	fmt1.Println(f(30))

	f1 := adder()
	fmt1.Println(f1(40))
	fmt1.Println(f1(50))
}
 */
/*
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		//判断字符串suffix是否以name结尾
		if !strings.HasSuffix(name,suffix){
			return name + suffix
		}
		return name
	}
}
func main() {
	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("test"))
	fmt.Println(txtFunc("test"))
}
 */
/*
func calc(base int) (func(int) int,func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}
	sub := func(i int) int {
		base -= i
		return base
	}
	return add,sub
}

func main() {
	f1,f2 := calc(10)
	fmt.Println(f1(1), f2(2))
}
 */
/*
func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}
func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}
 */
/*
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20
}
 */

func dispatchCoin() (left int) {
	var (
		coins = 50
		users = []string{
			"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
		}
		distribution = make(map[string]int, len(users))
	)
	for _,name := range users{
		getCoin := 0
		//这是不需要split
		for _,v := range name{
			switch v {
			case 'e', 'E': getCoin += 1
			case 'i', 'I': getCoin += 2
			case 'o', 'O': getCoin += 3
			case 'u', 'U': getCoin += 4
			}
		}
		//第二中字符串分拆的
		/*
		for _,v := range strings.Split(name,""){
			switch {
			case v == "e" || v == "E":
				getCoin++
			case v == "i" || v == "I":
				getCoin += 2
			case v == "o" || v == "O":
				getCoin += 3
			case v == "u" || v == "U":
				getCoin += 4
			}
		}
		 */
		distribution[name] = getCoin
		coins -= getCoin
	}
	fmt1.Println(distribution)
	return coins
}
func main() {
	left := dispatchCoin()
	fmt1.Println("剩下：", left)
}