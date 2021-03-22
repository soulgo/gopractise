package main

import "fmt"

//学生信息
type student struct {
	id, age     int
	name, class string
}

//学生管理信息
type stuMar struct {
	allStudents map[int]*student
}

//newStudent是student的学生构造函数
func newStudent(id, age int, name, class string) *student {
	return &student{
		id:    id,
		age:   age,
		name:  name,
		class: class,
	}
}

//newStuMar是stuMar的构造函数
func newStuMar() *stuMar {
	return &stuMar{
		allStudents: make(map[int]*student, 100),
	}
}

//1.添加学生信息
func (s stuMar) addStudent(newStu *student) {
	_, error := s.allStudents[newStu.id]
	if error {
		fmt.Println("已存在，若想修改，请重新选择2：修改功能")
	} else {
		s.allStudents[newStu.id] = newStudent(newStu.id, newStu.age, newStu.name, newStu.class)
	}
}

//2、编辑学生信息
func (s stuMar) modifyStudent(newStu *student) {
	_, error := s.allStudents[newStu.id]
	if !error {
		fmt.Println("该学生信息不存在，若想添加，请重新选择1：添加学生信息功能")
	} else {
		s.allStudents[newStu.id] = newStudent(newStu.id, newStu.age, newStu.name, newStu.class)
	}
}

//3、删除学生信息
func (s stuMar) deleteStudent(id int) {
	_, error := s.allStudents[id]
	if !error {
		fmt.Println("该学生信息不存在，若想添加，请重新选择1：添加学生信息功能")
	} else {
		delete(s.allStudents, id)
	}
}

//4.显示所有学生信息
func (s stuMar) showStumar() {
	fmt.Printf("\t%s\t%s\t%s\t%s\n", "学号  ", "姓名", "  年龄", "班级")
	for _, v := range s.allStudents {
		fmt.Printf("\t%d\t%s\t\t%d\t%s\n", v.id, v.name, v.age, v.class)
		//fmt.Printf("学号：%d  姓名：%s  年龄：%d  班级：%s\n",v.id,v.name,v.age,v.class)
	}
}
