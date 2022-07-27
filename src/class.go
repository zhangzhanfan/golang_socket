/*
	封装类
	类名首字母大写表示public，方法名首字母大写表示public
*/
package main

import "fmt"

// 定义一个结构体
type Hero struct {
	Name  string
	Ad    int
	Level int
}

func (this *Hero) GetName() string {
	return this.Name
}

func (this *Hero) SetName(name string) {
	this.Name = name
}

func (this *Hero) Show() {
	fmt.Println(this.Name)
	fmt.Println(this.Ad)
	fmt.Println(this.Level)
}

func main() {
	hero := Hero{Name: "张三", Ad: 10, Level: 3}
	hero.Show()
	hero.SetName("李四")
	hero.Show()
}
