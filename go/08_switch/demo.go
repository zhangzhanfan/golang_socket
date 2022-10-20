package main

import (
	"fmt"
)

func main() {
	var grade int = 60
	switch grade {
	case 60:
		fmt.Println("废物，不及格的废物！")
	case 70:
		fmt.Println("努力一点就可以摆脱废物的称号了！")
	case 80:
		fmt.Println("高不高，低不低，加油上90！")
	case 90:
		fmt.Println("有种超过90分啊！")
	default:
		fmt.Println("当你看到这句话的时候，你已经突破90大关了！")
	}

	// 不写 默认为true
	switch {
	case true:
		fmt.Println("对")
	case false:
		fmt.Println("错")
	default:
		fmt.Println("???")
	}

	b := false
	switch b {
	case false:
		fmt.Println("错")
		fallthrough //穿透语句，无论下一个条件是否成立，都会执行, break可以终止穿透
	case true:
		fmt.Println("对")
	default:
		fmt.Println("???")
	}
}
