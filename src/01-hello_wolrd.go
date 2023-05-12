package main

import (
	"fmt"
	"math"
	"math/rand"
)

/*
	每个 Go 程序都是由包构成的。
	程序从 main 包开始运行。
	import 导入路径 "fmt" 和 "math/rand" 来使用这两个包

	可以分开导入
	import "fmt"
	import "math"

	或者组合导入
	import (
		"fmt"
		"math/rand"
	)
*/

func main() {
	fmt.Printf("hello wolrd !\n") // printf 只可输出字符串  不会换行
	fmt.Printf("this is my first time to use go\n")
	fmt.Printf("let is generate a random int number\n")
	fmt.Println(rand.Intn(10)) // println  可以输出任意值 默认会换行
	fmt.Println("lets print Π value")
	fmt.Println(math.Pi)
	fmt.Print("end......")
	fmt.Print("end")
}
