package main

import (
	"fmt"
	"math/cmplx"
)

/*
	Go 的基本类型有
	bool
    string
	int  int8  int16  int32  int64
	uint uint8 uint16 uint32 uint64 uintptr
	byte // uint8 的别名
	rune // int32 的别名

	float32 float64
	complex64 complex128

	int, uint 和 uintptr 在 32 位系统上通常为 32 位宽，在 64 位系统上则为 64 位宽。 当你需要一个整数值时应使用 int 类型，除非你有特殊的理由使用固定大小或无符号的整数类型。
*/

func main() {

	//同导入语句一样，变量声明也可以“分组”成一个语法块。
	var (
		ToBe   bool       = true
		MaxInt uint64     = 1<<64 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)
	//初始化变量没有赋值时，会存在默认值，数值类型为 0， 布尔类型为 false， 字符串为 ""（空字符串）。
	// Printf格式化输出
	fmt.Printf("Type:%T,Value:%v\n", ToBe, ToBe)
	fmt.Printf("Type:%T,Value:%v\n", MaxInt, MaxInt)
	fmt.Printf("Type:%T,Value:%v\n", z, z)

	//变量类型转换
	var number1 int = 10
	var number2 float32 = float32(number1)
	var number3 uint = uint(number2)
	fmt.Printf("number1:%v,number2:%v,number3:%v\n", number1, number2, number3)
	fmt.Printf("number1:%T,number2:%T,number3:%T\n", number1, number2, number3)

	//类型推导 变量类型根据右边值决定
	number4 := 10
	number5 := 3.14
	number6 := "hello"
	fmt.Printf("number1:%T,number2:%T,number3:%T\n", number4, number5, number6)

	// 常量 常量可以是字符、字符串、布尔值或数值。 常量不能用 := 语法声明。
	// 常量赋值之后就无法修改值
	const number7 int = 999999
	fmt.Printf("number1:%v,\n", number7)

	// 数值常量
	// << 左移多少位相当于 * 2^n次
	// >> 右移多少位相当于 / 2^n 次
	const number8 = 2 << 2
	const number9 = 40 >> 2
}
