package main

import "fmt"

/*
type 接口类型名 interface {
    方法名1(参数列表1) 返回值列表1
    方法名2(参数列表2) 返回值列表2
    ...
}
*/

type Shape interface {
	area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func main() {
	var s Shape //s是一个接口！

	s = Rectangle{width: 10, height: 5}
	fmt.Printf("矩形面积: %f\n", s.area()) // 会自动调用对应的area函数

	s = Circle{radius: 3}
	fmt.Printf("圆形面积: %f\n", s.area())
}
