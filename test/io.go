package main

import (
	"fmt"
)

func main() {

	var str string
	fmt.Printf("请输入一个字符串：")
	fmt.Scanf("%s", &str)
	fmt.Println("输出：", str)

	var a int
	fmt.Scan(&a)
	fmt.Println(a)
}
