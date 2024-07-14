package main

import (
	"container/list"
	"fmt"
	"sync"
)

func Slice_() {
	// 初始化
	// 方式1：var team = [3]string{"hammer","soldier","mum"}
	// 方式2
	var team = [...]string{"hammer", "soldier", "mum"}
	fmt.Println(team)

	// 进行切片
	// slice [开始位置:结束位置]
	var highRiseBuilding [30]int

	for i := 0; i < 30; i++ {
		highRiseBuilding[i] = i + 1
	}
	fmt.Println(highRiseBuilding[10:15])

	fmt.Println(highRiseBuilding[20:])

	fmt.Println(highRiseBuilding[:2])
	/**输出——
		[11 12 13 14 15]
	[21 22 23 24 25 26 27 28 29 30]
	[1 2]
	*/

	// 这表示生成原有切片
	a := []int{1, 2, 3}
	fmt.Println(a[:])

	// 生成的切片将变空
	a2 := []int{1, 2, 3}
	fmt.Println(a2[0:0])

	// 使用make()函数构造切片
	// 如果需要动态地创建一个切片，可以使用make()内建函数:make([]T,size,cap)
	a3 := make([]int, 2)
	b3 := make([]int, 2, 10)

	fmt.Println(a3, b3)
	fmt.Println(len(a3), len(b3))

	// 输出：
	// [0,0][0,0]
	// 2 2
	//a和b均是预分配2个元素的切片，只是b的内部存储空间已经分配了10个，但实际使用了2个元素。

	//Apend 自动扩容
	// Go语言的内建函数append()可以为切片动态添加元素。每个切片会指一片内存空间 ，这篇内存空间能容纳一定数量的元素。当空间不能容纳足够多的元素时，切片就会进行"扩容"。"扩容"操作往往发生在append()函数调用时。

	var numbers []int
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
		fmt.Printf("len：%d  cap: %d  pointer:  %p\n", len(numbers), cap(numbers), numbers)
	}
	/**
	  输出：
	  len：1  cap: 1  pointer:  0xc00001c0a8
	  len：2  cap: 2  pointer:  0xc00001c0f0
	  len：3  cap: 4  pointer:  0xc000014180
	  len：4  cap: 4  pointer:  0xc000014180
	  len：5  cap: 8  pointer:  0xc000018440
	  len：6  cap: 8  pointer:  0xc000018440
	  len：7  cap: 8  pointer:  0xc000018440
	  len：8  cap: 8  pointer:  0xc000018440
	  len：9  cap: 16  pointer:  0xc000020180
	  len：10  cap: 16  pointer:  0xc000020180
	  **/

	// 删除元素
	//Go语言中切片删除元素的本质是：以被删除元素为分界点，将前后两个部分的内存重新连接起来。
	seq := []string{"a", "b", "c", "d", "e"}

	index := 2

	fmt.Println(seq[:index], seq[index+1:])

	seq = append(seq[:index], seq[index:]...)
	// apend可以用类似于js的...

	fmt.Println(seq)

}

func Map_() {
	scene := make(map[string]int)

	scene["route"] = 66

	fmt.Println(scene["route"])

	v := scene["route2"]

	fmt.Println(v) // 输出0，因为默认值为0，找不到。

	//某些情况下，需要明确知道查询中某个键是否在map中存在，可以使用一种特殊的写法来实现，如下代码：
	v, ok := scene["route3"] //如果键不存在，ok的值为false，v的值为该类型的零值
	fmt.Println(v, ok)

	// 遍历键值对
	scene = make(map[string]int)
	scene["route"] = 66
	scene["brazil"] = 4
	scene["china"] = 960
	for k, v := range scene {
		fmt.Println(k, v)
	}

	// 删除键
	scene = make(map[string]int)
	scene["route"] = 66
	scene["brazil"] = 4
	scene["china"] = 960

	delete(scene, "brazil")

	for k := range scene {
		fmt.Println(k)
	}

	SyncMap()
}
func SyncMap() {
	/**
		Go语言中的map在并发情况下，只读是线程安全的，同时读写线程不安全

	需要并发读写时，一般的做法是加锁，但这样性能并不高。Go语言在1.9版本中提供了一种效率较高的并发安全的sync.Map。sync.Map和map不同，不是以语言原生形态提供而是在sync包下的特殊结构。

	sync.Map有一下特性：

	无须初始化，直接声明即可。
	sync.Map不能使用map的方式进行取值和设置等操作，而是使用sync.Map的方法进行调用。
	Store表示存储，Load表示获取，Delete表示删除。
	使用Range配合一个回调函数进行遍历操作，通过回调函数返回内部遍历出来的值。Range参数中的回调函数的返回值功能是：需要继续迭代遍历时，返回true；终止迭代遍历时，返回false
	*/
	var scene sync.Map // 注意导入包

	scene.Store("freece", 97)
	scene.Store("london", 100)
	scene.Store("egypt", 200)

	fmt.Println(scene.Load("london")) // 100 true,即有两个返回值。第一个为值，第二个为是否存在

	scene.Delete("london")

	scene.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		return true
	})
	// /输出

	// iterate: freece 97
	// iterate: egypt 200
}

func List_() {
	l := list.New() // 定义列表，注意导入包

	l.PushBack("fist") //这俩也有返回值，没有接收罢了
	l.PushFront(67)
	element := l.PushFront("first") // 返回为该插入元素的值
	l.InsertAfter("high", element)
	l.InsertBefore("noon", element)
	l.Remove(element)

	// 遍历
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

}

func main() {
	fmt.Println("VECTOR")
	Slice_()
	Map_()
	List_()
}
