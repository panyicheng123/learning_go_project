package main

import "fmt"

/*
	类型 [n]T 表示拥有 n 个 T 类型的值的数组。
*/

func main() {
	var array1 [10]int

	for i := 0; i < len(array1); i++ {
		array1[i] = i
	}
	fmt.Println(array1)

	array2 := [3]int{3, 4, 5}
	fmt.Println(array2)

	array3 := [5]int{1, 2, 3, 4}
	fmt.Println(array3)
}
