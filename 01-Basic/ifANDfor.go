package main

import "fmt"

func repeate() {
	// 传统的 for 循环
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// 无限 for 循环，使用 break 语句来退出
	for {
		fmt.Println("这是无限循环中的一行代码。")
		break
	}

	//for key, value := range oldMap
	for _, v := range []int{1, 2, 3} {
		fmt.Println(v)
	}
}

func main() {
	/* 定义局部变量 */
	var a int = 10

	/* 使用 if 语句判断布尔表达式 */
	if a < 20 {
		/* 如果条件为 true 则执行以下语句 */
		fmt.Printf("a 小于 20\n")
	}
	fmt.Printf("a 的值为 : %d\n", a)

	repeate()
}
