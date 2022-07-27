/*
	类的继承
*/
package main

import "fmt"

// 定义一个动物接口 ,实际是一个指针
type AnimalIF interface {
	Sleep()
	GetColor() string
	GetType() string
}

// 实现类
type Cat struct {
	color string
}

func (this *Cat) Sleep() {
	fmt.Println("猫在睡觉.....")
}
func (this *Cat) GetColor() string {
	return this.color
}
func (this *Cat) GetType() string {
	return "这是一只猫...."
}

// 实现类
type Dog struct {
	color string
}

func (this *Dog) Sleep() {
	fmt.Println("狗在睡觉.....")
}
func (this *Dog) GetColor() string {
	return this.color
}
func (this *Dog) GetType() string {
	return "这是一只狗...."
}

func main() {
	// var animal AnimalIF //父类指针
	// animal = &Cat{"yellow"}
	// animal = &Dog{"block"}

	showAnimal(&Cat{"yellow"})

	showAnimal(&Dog{"block"})
}

//AnimalIF是指针类型
func showAnimal(animal AnimalIF) {
	animal.Sleep()
	fmt.Println(animal.GetColor())
	fmt.Println(animal.GetType())
}
