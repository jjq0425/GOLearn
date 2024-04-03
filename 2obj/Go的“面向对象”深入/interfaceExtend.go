package main

import "fmt"

// 定义接口

type A interface {
	testA()
}

type B interface {
	testB()
}

type C interface {
	A
	B
	testC()
}

//定义结构体

type Str struct {
}

//结构体实现接口

func (str Str) testA() {
	fmt.Println("testA")
}
func (str Str) testB() {
	fmt.Println("testB")
}
func (str Str) testC() {
	fmt.Println("testC")
}
func main() {

	//子类可以调用实现的所有方法
	str := Str{}
	str.testA()
	str.testB()
	str.testC()

	//多态:只能调用父接口有的方法
	var a A = Str{} //这句话的意思是：将 Str 类型的实例赋值给 A 类型的变量 a。因为 Str 实现了 A 接口
	a.testA()

	var b B = Str{}
	b.testB()

	var c C = Str{}
	c.testA()
	c.testB()
	c.testC()
}
