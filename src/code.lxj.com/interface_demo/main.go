package main

import "fmt"

/*
type Mover interface {
	move()
}
type dog struct {}

func (d *dog) move() {
	fmt.Println("狗会动")
}

func main() {
	var x Mover
	//var wangcai = dog{}
	//x = wangcai
	var fugui = &dog{}
	x = fugui
	x.move()
}
*/
/*
type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}

func main() {
	var peo People = &Student{}
	think := "sb"
	fmt.Println(peo.Speak(think))
}
/*
func main() {
	/*
	var m = make(map[string]interface{},16)
	m["name"] = "娜扎"
	m["age"] = 18
	m["hobby"] = []string{"篮球","足球","双色球"}
	fmt.Println(m)
	//var x interface{}
	//x = "hello"
	//x = 100
	//x = false
	//switch v := x.(type) {
	//case string:
	//	fmt.Printf("是字符串类型，value:%v\n",v)
	//case int:
	//	fmt.Printf("是整型类型，value:%v\n",v)
	//case bool:
	//	fmt.Printf("是布尔类型，value:%v\n",v)
	//}
}
*/
func justifyType(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string，value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupport type！")
	}
}
func main() {
	var x interface{}
	x = "hello"
	justifyType(x)
}
