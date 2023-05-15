package main

import "fmt"

/*
	每个数组的大小都是固定的。而切片则为数组元素提供动态大小的、灵活的视角。在实践中，切片比数组更常用。
	类型 []T 表示一个元素类型为 T 的切片。
	它会选择一个半开区间，包括第一个元素，但排除最后一个元素。
	以下表达式创建了一个切片，它包含 a 中下标从 1 到 3 的元素：a[1:4]
*/

func main() {
	var array1 = [5]int{1, 2, 3, 4, 5}
	var slice []int = array1[1:4]
	fmt.Println(slice)

	/*
		切片就像数组的引用
		切片并不存储任何数据，它只是描述了底层数组中的一段。

		更改切片的元素会修改其底层数组中对应的元素。

		与它共享底层数组的切片都会观测到这些修改。
	*/

	var array2 = [3]int{1, 2, 3}
	var slice2 = array2[0:2]
	slice2[0] = 9
	slice2[1] = 9
	fmt.Println(array2)

	//切片文法类似于没有长度的数组文法。

	bool_array := []bool{true, false, true, false, true}
	fmt.Println(bool_array)

	string_array := []string{"my", "name", "is", "go", "lang"}
	fmt.Println(string_array)

	struct_array := []struct{ X, Y int }{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
	}
	fmt.Println(struct_array)

	//切片的默认行为,在进行切片时，你可以利用它的默认行为来忽略上下界。以下都是等价的
	var array3 = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(array3[:])
	fmt.Println(array3[0:])
	fmt.Println(array3[:10])
	fmt.Println(array3[0:10])

	//切片的长度与容量,切片的长度就是它所包含的元素个数。
}
