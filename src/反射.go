/*
	反射
*/

package main

import (
	"fmt"
	"reflect"
)

func reflectNum(arg interface{}) {
	fmt.Println(reflect.TypeOf(arg))
	fmt.Println(reflect.ValueOf(arg))
}

func main() {
	var num float64 = 12.221
	reflectNum(num)
}
