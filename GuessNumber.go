package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	maxNum := 100
	fmt.Printf("输入最大数：")
	fmt.Scanf("%d\n",&maxNum)
	if maxNum <= 0{
		maxNum = 100
		fmt.Println("您输入的数字不合规，最大数重置到100")
	}
	
	rand.Seed(time.Now().UnixNano())//随机数必须输入种子
	rndnum := rand.Intn(maxNum)

	var input int
	i := 1
	
	var lower int = 0
	var higher int = maxNum

	fmt.Println(maxNum)

	for {//for死循环
		fmt.Printf("输入数字,第%d次,\n数字在%d到%d之间\n",i,lower,higher)
		fmt.Scanf("%d\n", &input)
		
		if input > higher || input < lower {
			fmt.Println("您输入的数字不符合要求！")
			continue
		}else if input < rndnum {
			fmt.Println("数字过小,\n")
			lower = input
			i++
			continue

		} else if input > rndnum {
			fmt.Println("数字过大,\n")
			higher = input
			i++
			continue

		} else {
			fmt.Println("猜测成功！")
			break
		}
	}
	fmt.Printf("共猜测%d次\n",i)
	fmt.Scanf("\n")
	/*next：
	将结果记录到文件
	终止停顿
	*/
}
