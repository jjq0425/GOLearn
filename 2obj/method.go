package main

import "fmt"

type Bag struct {
	items []int
}

// 不是函数但胜似函数
/*
func (接收器变量 接收器类型) 方法名(参数列表) (返回参数) {
    函数体
}

*/

// 注意：由于指针的特性，调用方法时，修改接收器（如下面的b）指针的任意成员变量，在方法结束后，修改都是有效的。
func (b *Bag) Insert(itemid int) {
	b.items = append(b.items, itemid)
}

func (b *Bag) GetFirstItem() int {
	return b.items[0]
}

// 若传递的不是指针，那么对其的修改无效，相当于浅拷贝了一份。
func (b Bag) GetFirstItem2() int {
	return b.items[0]
}

func main() {
	b := new(Bag)
	b.Insert(1001) // 调用方法
	b.Insert(1002) // 调用方法
	fmt.Println(b.GetFirstItem())
}
