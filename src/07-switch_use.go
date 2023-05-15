package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
switch 是编写一连串 if - else 语句的简便方法。它运行第一个值等于条件表达式的 case 语句。

Go 只运行选定的 case，而非之后所有的 case。 实际上，Go 自动提供了在这些语言中每个 case 后面所需的 break 语句。

case 语句从上到下顺次执行，直到匹配成功时停止。
*/
func main() {
	var os string = runtime.GOOS
	switch os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	var today = time.Now().Weekday()

	switch time.Saturday {
	case today + 0:
		fmt.Println("今天就是周六")
	case today + 1:
		fmt.Println("明天就是周六")
	case today + 2:
		fmt.Println("距离周六还剩两天")
	default:
		fmt.Println("距离放假还有好多天")
	}

	// 没有条件的switch
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
