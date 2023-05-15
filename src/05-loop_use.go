package main

import (
	"fmt"
)

/*
	for 语句可以进行循环功能的实现
	基本的 for 循环由三部分组成，它们用分号隔开：

	初始化语句：在第一次迭代前执行
	条件表达式：在每次迭代前求值
	后置语句：在每次迭代的结尾执行
	初始化语句通常为一句短变量声明，该变量声明仅在 for 语句的作用域中可见。

	一旦条件表达式的布尔值为 false，循环迭代就会终止。
*/

func main() {
	var sum int = 0
	for i := 0; i < 10; i++ {
		sum = sum + i
	}
	fmt.Println(sum)

	//初始化语句和后置语句是可选的。 这样就相当于while语句了
	var sum2 int = 1

	for sum2 < 10 {
		sum2 = sum2 + 1
	}
	fmt.Println(sum2)

	// 无限循环
	for {
		fmt.Printf("无限循环中\n")
		// time.Sleep(1000)
	}
}
