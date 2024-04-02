package main

import "fmt"

// 这里将返回两个int
func test() (int, int) {
	return 0, 1
	// 这个函数可以返回两个值
}
func VAR() {
	/**变量**/
	var name string = "hello world"
	// :=就是auto
	name2 := "hello world2"
	fmt.Println(name, name2)

	const (
		a1 = 1
		a2 // 这里将会自动沿用上面的值,即a2还是1

	)
	fmt.Println(a1, a2)

	/**iota**/
	// iota是一个内置的自增长变量，每次定义会进行一次自增
	const (
		a = iota // 0
		b = iota // 1
		c = iota // 2
		d        // 自动沿用上面的值，上面的值是iota，所以自增变成3
		e = "hah"
		f        // 自动沿用上面的值，为hah
		g = iota // 再次自增，为6! 注意每次定义变量即新增
	)
	fmt.Println(a, b, c, d, e, f, g)

	const (
		another_a = iota // 新的一组开始新的计数，所以为0
	)
	fmt.Println(another_a)
}

func main() {
	// 打印
	fmt.Printf("hello, world\n")

	// 打印并换行
	fmt.Println("hello, world")

	VAR()

	/**函数**/
	test_1, _ := test()
	// 这里的 _ 是个匿名变量，后面不会使用，所以用_直接代替即可。_可多次反复使用
	fmt.Println(test_1)

}
