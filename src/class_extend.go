/*
	类的继承
*/
package main

import "fmt"

// 定义一个父类
type Human struct {
	name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("人会吃..........")
}

func (this *Human) Walk() {
	fmt.Println("人会走路啊..........")
}

type SuperMan struct {
	Human
	Level int
}

//重写父类方法
func (this *SuperMan) Walk() {
	fmt.Println("超人会快速的走路..........")
}

func (this *SuperMan) Fly() {
	fmt.Println("超人会飞..........")
}

func main() {
	//方法一
	// superMan := SuperMan{Human{name: "小红", sex: "女"}, 2}
	//方法二
	var superMan SuperMan
	superMan.name = "小红"
	superMan.sex = "女"
	superMan.Level = 2

	superMan.Eat()
	superMan.Walk()
	superMan.Fly()
}
