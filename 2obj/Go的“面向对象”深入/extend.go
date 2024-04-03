package main

import "fmt"

// 继承： 可以直接通过子类对象名.父类字段调用父类属性

//定义父类结构体

type Person struct {
	name string
	age  int
}

//定义子类结构体

type Student struct {
	Person // 通过定义父类的匿名字段实现继承
	school string
}

func main() {

	//1.定义单独的父类对象
	person := Person{
		name: "张三",
		age:  23,
	}
	fmt.Println(person)

	//2.定义子类对象
	student := Student{
		Person{
			name: "李四",
			age:  23,
		},
		"山东大学",
	}
	fmt.Println(student)
	fmt.Println(student.name) //注意: 继承之后可以直接通过子类对象名.父类字段调用父类属性

}
