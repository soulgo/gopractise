package main

import (
	"fmt"
	"os"
)

/*
##切片模式
使用“面向对象”的思维方式编写一个学生信息管理系统。
学生有id、姓名、年龄、分数等信息
程序提供展示学生列表、添加学生、编辑学生信息、删除学生等功能
分析：1、增加学生信息
	 2、编辑学生信息
     3、删除学生信息
	 4、展示所有学生信息
*/

func showMenu() {
	fmt.Println("欢迎来到学生信息管理系统")
	fmt.Println("1.添加学生")
	fmt.Println("2.编辑学生信息")
	fmt.Println("3.删除学生信息")
	fmt.Println("4.展示所有学生信息")
}

//获取用户输入的学生信息
func getInput() *student {
	var (
		id, age     int
		name, class string
	)
	fmt.Println("请按要求输入学生信息：")
	fmt.Print("请输入学生的学号：")
	fmt.Scanf("%d\n", &id)
	fmt.Print("请输入学生的姓名：")
	fmt.Scanf("%s\n", &name)
	fmt.Print("请输入学生的年龄：")
	fmt.Scanf("%d\n", &age)
	fmt.Print("请输入学生的班级：")
	fmt.Scanf("%s\n", &class)
	//就能拿到用户输入的学生的所有信息
	stu := newStudent(id, age, name, class) //调用构造函数创造一个学生信息
	return stu
}

func main() {
	/*
		1、打开打印菜单
		2、等待用户选择要执行的选项
		3、执行用户选择的动作
	*/
	sm := newStudentMar()
	for {
		//1、打开打印菜单
		showMenu()
		//2、等待用户选择要执行的选项
		var input, id int
		fmt.Printf("请输入你要输入的序号：")
		fmt.Scanf("%d\n", &input)
		switch input {
		case 1:
			//添加学生信息
			stu := getInput()
			sm.addStudent(stu)
		case 2:
			//编辑学生信息
			stu := getInput()
			sm.modifyStudent(stu)
		case 3:
			//删除学生信息
			fmt.Print("请输入学生的学号：")
			fmt.Scanf("%d\n", &id)
			sm.delStudent(id)
		case 4:
			//展示所有学生信息
			sm.showAllStudent()
		case 5:
			//退出
			os.Exit(0)
		}
	}
}
