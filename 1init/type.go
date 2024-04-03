package main

import (
	"fmt"
	"math"
)

func Array() {
	var n [10]int /* n 是一个长度为 10 的数组 */
	var i, j int

	/* 为数组 n 初始化元素 */
	for i = 0; i < 10; i++ {
		n[i] = i + 100 /* 设置元素为 i + 100 */
	}

	/* 输出每个数组元素的值 */
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}

	//还可以使用初始化列表来初始化数组的元素：
	//var numbers = [5]int{1, 2, 3, 4, 5}
	//https://blog.csdn.net/m0_53328239/article/details/131493346
}

func Slice() {
	// 切片是一个拥有相同类型元素的可变长度的序列。切片的声明方式如下：

	a := make([]int, 3)

	a[0] = 1
	a[1] = 2
	a[2] = 3
	fmt.Println(a[0], a[1], a[2])

	// 切片还可以在其元素集合内连续地选取一段区域作为新的切片，就像其名字"切片"一样，切出一块区域，形成新的切片。
	str := "hello6world" // 可以把字符串也理解为切片
	fmt.Print(str[6:])   // 打印world，即切片开始即下标

}

func Bool() {
	//在Go语言中吗，布尔型只有true和false两值

	// Go语言不允许将整形强制转换为布尔型
}

func TypeChange() {
	//类型转换
	var a int32 = 1047483647
	fmt.Printf("int32: 0x%x %d\n", a, a)
	b := int16(a)
	fmt.Printf("int16: 0x%x %d\n", b, b)
	var c float32 = math.Pi
	fmt.Println(int(c)) // 转换为int
}

func String() {
	// 计算字符串长度
	str := "hello world"
	fmt.Println(len(str)) // 输出：11
}
func main() {

	Array()
	Slice()
	TypeChange()
	String()
}
