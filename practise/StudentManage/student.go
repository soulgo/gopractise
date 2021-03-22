package main

import (
	"fmt"
)

//学生信息
type student struct {
	id, age     int //学号是唯一的
	name, class string
}

//学生管理类型
type studentMar struct {
	//*student是因为结构体有多个字段
	//切片
	allStudents []*student
}

//newStudentMar是studentMar的构造函数
func newStudentMar() *studentMar {
	return &studentMar{
		allStudents: make([]*student, 0, 100),
	}
}

//newStudent 是student的构造函数
func newStudent(id, age int, name, class string) *student {
	return &student{
		id:    id,
		age:   age,
		name:  name,
		class: class,
	}
}

//1、添加学生
func (s *studentMar) addStudent(newStu *student) {
	s.allStudents = append(s.allStudents, newStu)
}

//2、编辑学生
func (s *studentMar) modifyStudent(newStu *student) {
	for i, v := range s.allStudents {
		if newStu.id == v.id { //当学号相同时就说明找到该学生
			s.allStudents[i] = newStu //根据切片的索引直接把新学生赋值进来
			return
		}
	}
	//没有找到输入的学号对应的学生信息
	fmt.Println("输入的学生信息有误，请检查重新输入！")
}

//3、删除学生
func (s *studentMar) delStudent(id int) {
	for i, v := range s.allStudents {
		if id == v.id {
			s.allStudents = append(s.allStudents[:i], s.allStudents[i+1:]...) //切片的删除方式，从切片中删除index的元素
		}
	}
}

//4、展示所有学生
func (s *studentMar) showAllStudent() {
	fmt.Printf("\t%s\t%s\t%s\t%s\n", "学号  ", "姓名", "  年龄", "班级")
	for _, v := range s.allStudents {
		fmt.Printf("\t%d\t%s\t\t%d\t%s\n", v.id, v.name, v.age, v.class)
		//fmt.Printf("学号：%d  姓名：%s  年龄：%d  班级：%s\n",v.id,v.name,v.age,v.class)
	}
}
