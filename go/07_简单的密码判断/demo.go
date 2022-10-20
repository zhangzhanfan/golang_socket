package main

import "fmt"

func main() {
	var a int
	var password int = 20221019

	fmt.Println("请输入你的密码：")
	fmt.Scanln(&a)
	if a == password {
		fmt.Println("请再次输入你的密码：")
		fmt.Scanln(&a)
		if a == password {
			fmt.Println("密码正确，登录成功！")
		} else {
			fmt.Println("密码错误！")
		}
	} else {
		fmt.Println("密码错误！")
	}
}
