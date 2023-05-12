package main

/*
	var 语句用于声明一个变量列表，类型在最后。
*/

var yong, beautiful, helpful bool // 默认值 false

var age int // 默认值 0

var name, nation string = "Pan YiCheng", "China" // 同一个类型多个变量赋值

var boy, address = true, "SUZHOU" //不同类型多个变量赋值

func main() {
	println(yong, beautiful, helpful)
	println(age)
	println(name, nation)
	println(boy, address)
}
