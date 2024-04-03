package main

import (
	"fmt"
)

func NewStruct() {
	/*
			type 类型名 struct {
		    字段1 字段1类型
		    字段2 字段2类型
		}

	*/

	// 实例化结构体1
	type Point struct {
		X int
		Y int
	}
	var p Point
	p.X = 10
	p.Y = 20
	fmt.Println(p)

	// 实例化结构体2
	type Player struct {
		Name        string
		HealthPoint int
		MagicPoint  int
	}
	tank := new(Player)
	tank.Name = "Canon"
	tank.HealthPoint = 300

	// 实例化3
	type Command struct { //定义 Command 结构体，表示命令行指令
		Name    string
		Var     *int //命令绑定的变量，使用整型指针绑定一个指针，指令的值可以与绑定的值随时保持同步。
		Comment string
	}
	var version int = 1 //命令绑定的目标整型变量：版本号。
	cmd := &Command{}   //对结构体取地址实例化。
	cmd.Name = "version"
	cmd.Var = &version
	cmd.Comment = "show version"

	// 实例化4
	newCommand := func(name string, varref *int, comment string) *Command {
		return &Command{
			Name:    name,
			Var:     varref,
			Comment: comment,
		}
	}
	cmd = newCommand(
		"version",
		&version,
		"show version",
	)
	fmt.Println(cmd) // 输出：&{version 0xc0000180c8 show version}

	// 嵌套初始化
	type People struct { //定义 People 结构体。
		name  string
		child *People //结构体的结构体指针字段，类型是 *People。
	}
	relation := &People{ //relation 由 People 类型取地址后，形成类型为 *People 的实例。
		name: "爷爷",
		child: &People{ //child 在初始化时，需要 *People 类型的值，使用取地址初始化一个 People。
			name: "爸爸",
			child: &People{
				name: "我",
			},
		},
	}
	fmt.Println(relation)

}

func CreateStruct() {
	type Cat struct {
		Color string
		Name  string
	}
	NewCatByName := func(name string) *Cat {
		return &Cat{
			Name: name,
		}
	}
	NewCatByColor := func(color string) *Cat {
		return &Cat{
			Color: color,
		}
	}
	fmt.Printf("%v\n", NewCatByName("Tom"))
	fmt.Printf("%v\n", NewCatByColor("black"))

	// 父子

	type BlackCat struct {
		Cat // Cat即上面定义的Cat
	}

	//构造子类
	NewBlackCat := func(color string) *BlackCat {
		cat := &BlackCat{}
		cat.Color = color
		return cat
	}

	fmt.Printf("%v\n", NewBlackCat("black")) // &{{black }} 注意多了一层维度

}

func main() {
	fmt.Println("struct")
	NewStruct()
	CreateStruct()
}
