package main

import "fmt"

/*
	一个结构体（struct）就是一组字段（field）

	结构体字段使用点号来访问
*/

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})

	var vex = Vertex{3, 4}

	fmt.Println(vex.X, vex.Y)

	/*
		如果我们有一个指向结构体的指针 p，那么可以通过 (*p).X 来访问其字段 X。
		不过这么写太啰嗦了，所以语言也允许我们使用隐式间接引用，直接写 p.X 就可以。
	*/

	var vex2 = Vertex{5, 6}
	p := &vex2
	fmt.Println((*p).X, (*p).Y)
	fmt.Println(p.X, p.Y)

	var (
		v1 = Vertex{1, 2}  // 创建一个 Vertex 类型的结构体
		v2 = Vertex{X: 1}  // Y:0 被隐式地赋予
		v3 = Vertex{}      // X:0 Y:0
		gg = &Vertex{1, 2} // 创建一个 *Vertex 类型的结构体（指针）
	)
	fmt.Println(v1, gg, v2, v3)
}
