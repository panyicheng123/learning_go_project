package main

import "fmt"

/*
    Go 拥有指针。指针保存了值的内存地址。
	类型 *T 是指向 T 类型值的指针。其零值为 nil。
*/

func main() {
	var p *int
	fmt.Println(p)

	// & 操作符会生成一个指向其操作数的指针。
	i := 43
	p = &i
	fmt.Println(p)
	//* 操作符表示指针指向的底层值
	fmt.Println(*p)
	*p = 45
	fmt.Println(*p)

	var number int = 45            // 声明一个变量赋值为45
	var ponint_to_number = &number // 生成该指向该变量的指针
	*ponint_to_number = 66         // 将指针指向的值赋值为 66
	fmt.Println(number)

}
