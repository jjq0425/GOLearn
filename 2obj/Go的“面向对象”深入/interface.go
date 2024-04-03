/*
	结构体用于定义属性,接口用于定义方法;接口只定义方法,不做方法的实现,结构体实现接口中的所有方法就叫实现了接口;
*/

// 定义接口
package main

import "fmt"

type Usb interface {
	Input()
	Output()
}

//定义结构体

type Mouse struct {
	brand string
}

type Keyboard struct {
	brand string
}

// 结构体重写接口中的所有方法来实现接口

func (mouse Mouse) Input() {
	fmt.Println(mouse.brand, "鼠标输入")
}

func (mouse Mouse) Output() {
	fmt.Println(mouse.brand, "鼠标输出")
}

func (keyboard Keyboard) Input() {
	fmt.Println(keyboard.brand, "键盘输入")
}

func (keyboard Keyboard) Output() {
	fmt.Println(keyboard.brand, "键盘输出")
}

// 定义一个测试函数,传递usb的实现类,自动去判断类型调用相应的方法

func test(usb Usb) {
	usb.Input()
	usb.Output()
}

func main() {

	//创建一个鼠标对象
	mouse := Mouse{"罗技"}
	test(mouse)

	//创建一个键盘对象
	keyboard := Keyboard{"樱桃"}
	test(keyboard)

}
