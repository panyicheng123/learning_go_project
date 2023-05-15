package main

import (
	"fmt"
	"math/rand"
)

/*
	if 语句可以实现对选择功能的是实现

	if 语句与 for 循环类似，表达式外无需小括号 ( ) ，而大括号 { } 则是必须的。
*/

func main() {

	var index int = 11
	if index < 10 {
		fmt.Println(index)
	} else {
		fmt.Println("index 大于10")
	}

	// 同 for 一样， if 语句可以在条件表达式前执行一个简单的语句。
	if value := rand.Intn(10); value > 5 {
		fmt.Printf("value-%v 大于 5", value)

	} else {
		fmt.Printf("value-%v 小于或等于 5", value)
	}

	var today string = "周二"
	if today == "周一" {
		fmt.Printf("今天是周一")
	} else if today == "周二" {
		fmt.Printf("今天是周二")
	} else if today == "周三" {
		fmt.Printf("今天是周三")
	} else if today == "周四" {
		fmt.Printf("今天是周四")
	} else if today == "周五" {
		fmt.Printf("今天是周五")
	} else {
		fmt.Printf("今天是周末")
	}
}
