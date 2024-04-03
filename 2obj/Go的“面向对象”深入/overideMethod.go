package main

import "fmt"

/*
	方法重写:在父子关系中,方法名称相同叫做重写,方法重写是基于继承的
*/

//定义父结构体

type Animal struct {
	name string
	age  int
}

//定义子结构体

type Dog struct {
	Animal
}

type Cat struct {
	Animal
}

//定义父类方法

func (animal Animal) eat() {
	fmt.Println(animal.name, "吃东西")
}

// 子类重写父类方法

func (dog Dog) eat() {
	fmt.Println(dog.name, "吃骨头")
}

func (cat Cat) eat() {
	fmt.Println(cat.name, "吃鱼")
}

// 子类除了重写父类方法还可以有自己的方法

func (cat Cat) sleep() {
	fmt.Println(cat.name, "睡觉")
}

func main() {

	// 创建一个动物对象
	animal := Animal{
		name: "动物",
		age:  2,
	}
	animal.eat()

	// 创建一个猫对象
	cat := Cat{
		Animal{
			"大花",
			3,
		},
	}
	cat.eat()
	cat.sleep()
}
