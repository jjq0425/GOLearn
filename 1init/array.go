package main

import "fmt"

func main() {

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
