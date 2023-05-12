package main

import "fmt"

/*
	函数的创建与使用

	函数可以没有参数或接受多个参数。

	在本例中，add 接受两个 int 类型的参数。

	注意类型在变量名 之后。
*/

func add(x int, y int) int {
	return x + y
}

/*
	当连续两个或多个函数的已命名形参类型相同时，除最后一个类型以外，其它都可以省略。
*/

func add2(x, y int) int {
	return x + y
}

/*
	函数可以返回多个返回值

	在函数中，简洁赋值语句 := 可在类型明确的地方代替 var 声明。

    函数外的每个语句都必须以关键字开始（var, func 等等），因此 := 结构不能在函数外使用。
*/

func swap(firstname string, lastname string) (string, string) {
	return firstname, lastname
}

func main() {
	fmt.Printf("参数未省略函数add结果输出:")
	fmt.Println(add(1, 2))

	fmt.Printf("多个参数类型相同省略函数add2结果输出:")
	fmt.Println(add2(3, 4))

	fmt.Printf("返回多个返回值的函数swap结果输出：")
	var c, d = swap("test1", "test2")
	a, b := swap("Pan", "YiCheng") // :=可以使得变量不用声明就直接使用，只能在函数内使用
	test := 123
	fmt.Println(a, b)
	fmt.Println(c, d)
	fmt.Println(test)
}
