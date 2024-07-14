package main

import (
	"bytes"
	"fmt"
)

/*
	func 函数名(参数列表) (返回参数列表) {
	    函数体
	}
*/
func add(a, b int) (int, bool) {
	return a + b, true
	//Go语言也支持多返回值，多返回值能方便地获得函数执行后的多个返回参数
}

func shadeFunc() {
	// 匿名函数
	f := func(data int) {
		fmt.Println("hello", data)
	}
	f(100) // 执行
	//输出
	// hello，100

	shaeFuncAsCallback([]int{1, 2, 3, 4}, func(v int) { fmt.Println(v) })

}

func BiBao() {
	//引用了外部变量的匿名函数

	// 闭包是引用了自由变量的函数，被引用的自由变量和函数一同存在，即使已经离开了自由变量的环境也不会被释放或者删除，在闭包中可以继续使用这个自由变量。
	accumulate := Accumulate(1)
	fmt.Println("ACUU", accumulate())
	fmt.Println("ACUU", accumulate())
	accumulator2 := Accumulate(10)
	fmt.Println("ACUU", accumulator2())
}

func ChangeAbleFunc() {
	/*
		func 函数名(固定参数列表,v...T)(返回参数列表) {
		    函数体
		}
	*/
	// 可变参数一般被放置在函数列表的末尾，前面时固定参数列表，当没有固定参数时，所有变量就将是可变参数
	// T为可变参数的类型，当T为interface{}时，传入的可以时任意类型。

	//输入3个字符串，将它们连成一个字符串
	fmt.Println(joinStrings("pig", "and", "rat"))
	fmt.Println(joinStrings("hammer", "mom", "and", "hawk"))
	//将不同类型的变量通过printTypeValue()
	fmt.Println(printTypeValue(100, "str", true))

	// 多个可变参数函数中嵌套传递参数
	print(1, 2, 3)
}

// 定义一个函数，参数数量为0-n，类型约束为字符串
func joinStrings(slist ...string) string {
	//定义一个字节缓冲，快速地连接字符串
	var b bytes.Buffer
	//遍历可变参数列表slist，类型为[]string
	for _, s := range slist {
		//将遍历出的字符串连续写入字节数组
		b.WriteString(s)
	}
	//将连接好的字节数组转换为字符串并输出
	return b.String()

}

// 当可变参数为interface{}类型时，可以传入任何类型的值。此时，如果需要获得变量的类型，可以通过switch类型分支获得变量的类型。
func printTypeValue(slist ...interface{}) string {
	//字节缓存作为快速字符串连接
	var b bytes.Buffer
	//遍历参数
	for _, s := range slist {
		//将interface{}类型格式化为字符串
		str := fmt.Sprintf("%v", s)
		//类型的字符串描述
		var typeString string
		//对s进行类型断言
		switch s.(type) {
		case bool:
			typeString = "bool"
		case string:
			typeString = "string"
		case int:
			typeString = "int"
		}
		//写值字符串前缀
		b.WriteString("value: ")
		//写入值
		b.WriteString(str)
		//写类型前缀
		b.WriteString(" type: ")
		//写类型字符串
		b.WriteString(typeString)
		//写入换行符
		b.WriteString("\n")
	}
	//将连接好的字节数组转换为字符串输出
	return b.String()
}

func rawPrint(rawList ...interface{}) {
	//遍历可变参数切片
	for _, a := range rawList {
		//打印参数
		fmt.Println(a)
	}
}

// 打印函数封装
func print(slist ...interface{}) {
	//如果要在多个可变参数中传递参数，可以在传递时在可变参数变量中默认添加"…",将切片中的元素进行传递，而不是传递可变参数变量本身。
	rawPrint(slist...)
}

func Accumulate(value int) func() int {
	//被捕获到闭包中的变量让闭包本身拥有了[记忆效应]，闭包中的逻辑可以修改闭包捕获的变量，变量会跟随闭包生命期一直存在，闭包本省就如同变量一样拥有了记忆效应。
	return func() int {
		value++
		return value
	}
}

func shaeFuncAsCallback(list []int, f func(int2 int)) {
	for _, v := range list {
		f(v)
	}

}

func main() {
	add(1, 2)
	shadeFunc()
	BiBao()

}
