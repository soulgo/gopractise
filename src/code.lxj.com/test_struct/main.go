package main

import (
	"fmt"
	"sort"
)

//结构体
/*
func main() {
	//匿名接口体
	var user struct{Name string;Age int}
	user.Name = "小王子"
	user.Age = 18
	fmt.Printf("%#v\n", user)
}
*/
/*
type person struct {
	name,city string
	age int8
}

func main() {
	//var p2 = new(person)
	//fmt.Printf("%T\n",p2)
	//fmt.Printf("p2=%#v\n",p2)
	var p2 = new(person)
	p2.name = "小王子"
	p2.age = 19
	p2.city = "兰州"
	fmt.Printf("p2=%#v\n",p2)
}
*/
/*
type student struct {
	name string
	age  int
}

func main() {
	m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}
	//原因：for range每次产生的 k , v 都是一个值拷贝，不是 stus对应值得引用，故而会出现这种结果。
	for _, stu := range stus {
		stu := stu
		m[stu.name] = &stu
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
}
*/
//结构体
/*
type person struct {
	name,city string
	age int8
}
//newPerson 构造函数
func newPerson(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:age,
	}
}
//Dream person做梦的方法
func (p person) Dream() {
	fmt.Printf("%s的梦想是学好Go语言！\n", p.name)
}
func main() {
	p1 :=  newPerson("张三","沙河",90)
	p1.Dream()
}
*/
/*
//Address 地址结构体
type Address struct {
	Province string
	City     string
}
//User 用户结构体
type User struct {
	Name    string
	Gender  string
	Address Address
}

func main() {
	user1 := User{
		Name: "小王子",
		Gender: "男",
		Address: Address{
			Province: "山东",
			City: "威海",
		},
	}
	fmt.Printf("user1=%#v\n",user1)
}
*/
/*
//Address 地址结构体
type Address struct {
	Province   string
	City       string
	AddressCreateTime string
}

//Email 邮箱结构体
type Email struct {
	Account    string
	EmailCreateTime string
}

//User 用户结构体
type User struct {
	Name   string
	Gender string
	Address
	Email
}

func main() {
	var user3 User
	user3.Name = "沙河娜扎"
	user3.Gender = "男"
	// user3.CreateTime = "2019" //ambiguous selector user3.CreateTime
	user3.Address.AddressCreateTime = "2000" //指定Address结构体中的CreateTime
	user3.Email.EmailCreateTime = "2000"   //指定Email结构体中的CreateTime
	fmt.Println(user3)
}
*/
/*
type Animal struct {
	name string
}

func (a *Animal) move(){
	fmt.Println("%s会动！\n",a.name)
}

type Dog struct {
	Feet int8
	*Animal
}

func (d *Dog) wang(){
	//fmt.Printf("%s会汪汪\n",d.name)
	fmt.Printf("%s会汪汪汪~\n", d.name)
}
func main() {
	d1 := &Dog{
		Feet: 4,
		Animal: &Animal{
			name: "乐乐",
		},
	}
	d1.wang()
	d1.move()
}
*/
/*
//Student 学生
type Student struct {
	ID     int
	Gender string
	Name   string
}

//Class 班级
type Class struct {
	Title    string
	Students []*Student
}

func main() {
	c :=&Class{
		Title: "101",
		Students: make([]*Student,0,200),
	}
	for i := 0;i< 10;i++{
		stu := &Student{
			Name: fmt.Sprintf("stu%02d",i),
			Gender: "男",
		}
		c.Students = append(c.Students,stu)
	}
	data,err := json.Marshal(c)
	if err != nil{
		fmt.Println("json marshal failed")
		return
	}
	fmt.Printf("json:%s\n",data)
	str := `{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}`
	c1 := &Class{}
	err = json.Unmarshal([]byte(str),c1)
	if err != nil{
		fmt.Println("json unmarshal failed")
		return
	}
	fmt.Printf("%#v\n",c1)
}
*/
/*
//Student 学生
type Student struct {
	ID     int    `json:"id"` //通过指定tag实现json序列化该字段时的key
	Gender string //json序列化是默认使用字段名作为key
	name   string //私有不能被json包访问
}

func main() {
	s1 := Student{
		ID:     1,
		Gender: "男",
		name:   "沙河娜扎",
	}
	data, err := json.Marshal(s1)
	fmt.Println(data)
	if err != nil {
		fmt.Println("json marshal failed!")
		return
	}
	fmt.Printf("json str:%s\n", data) //json str:{"id":1,"Gender":"男"}
}
*/
/*
type Person struct {
	name   string
	age    int8
	dreams []string
}

func (p *Person) SetDreams(dreams []string) {
	p.dreams = make([]string,len(dreams))
	copy(p.dreams,dreams)
}

func main() {
	p1 := Person{name: "小王子", age: 18}
	data := []string{"吃饭", "睡觉", "打豆豆"}
	p1.SetDreams(data)

	// 你真的想要修改 p1.dreams 吗？
	data[1] = "不睡觉"
	fmt.Println(p1.dreams)  // ?
}
*/
/*
使用“面向对象”的思维方式编写一个学生信息管理系统。
学生有id、姓名、年龄、分数等信息
程序提供展示学生列表、添加学生、编辑学生信息、删除学生等功能
*/
type Student struct {
	Id, Age, Score int
	Name           string
}

type Class struct {
	Map map[int]*Student
}

//添加学生信息
func (c *Class) addStudent() {
	var id, age, score int
	var name string
	fmt.Print("输入id: ")
	_, err := fmt.Scan(&id)
	fmt.Print("输入姓名: ")
	_, err = fmt.Scan(&name)
	fmt.Print("输入年龄: ")
	_, err = fmt.Scan(&age)
	fmt.Print("输入分数: ")
	_, err = fmt.Scan(&score)
	if err != nil {
		fmt.Println("输入格式有问题，请检查输入信息！")
	}
	_, isSave := c.Map[id]
	if isSave {
		fmt.Println("学生ID已存在")
		return
	}
	student := &Student{
		Id:    id,
		Name:  name,
		Age:   age,
		Score: score,
	}
	c.Map[id] = student
	fmt.Println("保存成功")
}

//查看学生列表
func (c *Class) showStuden() {
	fmt.Printf("\t%s\t%s\t%s\t%s\n", "ID", "姓名", "年龄", "分数")
	sortId := make([]int, 0)
	for k := range c.Map {
		sortId = append(sortId, k)
	}
	sort.Ints(sortId)
	for _, k := range sortId {
		s := c.Map[k]
		fmt.Printf("\t%d\t%s\t%d\t%d\n", s.Id, s.Name, s.Age, s.Score)
	}
}

//删除学生
func (c *Class) deleteStudent() {
	fmt.Print("输入要删除学生的Id: ")
	var id int
	_, err := fmt.Scan(&id)
	if err != nil {
		fmt.Println("请检查ID输入格式")
	}
	_, isSave := c.Map[id]
	if !isSave {
		fmt.Println("要删除的Id不存在！")
		return
	}
	delete(c.Map, id)
	fmt.Println("删除成功！！！")
}

//修改学生信息
func (c *Class) modiyStudent() {
	fmt.Print("输入要修改学生的Id： ")
	var id, age, score int
	var name string
	_, err := fmt.Scan(&id)
	if err != nil {
		fmt.Println("请检查输入ID的格式！")
	}
	_, isSave := c.Map[id]
	if !isSave {
		fmt.Println("要修改的ID不存在！")
		return
	}
	fmt.Print("输入姓名: ")
	_, err = fmt.Scan(&name)
	fmt.Print("输入年龄: ")
	_, err = fmt.Scan(&age)
	fmt.Print("输入分数: ")
	_, err = fmt.Scan(&score)
	if err != nil {
		fmt.Println("保存出错！")
	}
	student := &Student{
		Id:    id,
		Name:  name,
		Age:   age,
		Score: score,
	}
	c.Map[id] = student
	fmt.Println("修改成功")

}
func main() {
	c := &Class{}
	c.Map = make(map[int]*Student, 50)
Loop:
	for {
		fmt.Println("要执行的操作：")
		fmt.Print("1. 添加  2.查看  3.删除  4.修改 5.退出\n")
		var do int8
		_, err := fmt.Scan(&do)
		if err != nil {
			fmt.Println("输入有误！")
		}
		switch do {
		case 1:
			c.addStudent()
		case 2:
			c.showStuden()
		case 3:
			c.deleteStudent()
		case 4:
			c.modiyStudent()
		case 5:
			break Loop
		default:
			fmt.Println("输入有误！")
		}
	}
}
