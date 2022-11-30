package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//var num int8 = 10
	num := 10
	fmt.Println("哈哈,a=", num)

	fmt.Println("====================分割线=========================")

	var n1, s2, n3 = 108, "嗨害嗨", 25
	fmt.Println(n1, s2, n3)

	fmt.Printf("类型：%T\n", s2)
	fmt.Println(unsafe.Sizeof(s2))

	var stuname uint8 = 26
	fmt.Println(stuname)

	var nfl float32 = 31415e-4
	fmt.Println(nfl)
}
