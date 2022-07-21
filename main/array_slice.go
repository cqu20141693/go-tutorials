package main

import "fmt"

// 数组切片

/*ArraySlice
数组： 值传递，长度不可变，类型=元素类型+数组长度
切片： 引用传递，可以追加，类型=元素类型
*/
func ArraySlice() {

	fmt.Println("....... ArraySlice.........")
	array()
	slice()
}

/*slice
切片： 引用类型
初始化: {}, make(),new()
获取: []
设置: []=
截取：[:]
长度：len
遍历： for,range

cap(s) 获取切片的容量
append(s, ...) 向切片追加内容
copy(s, s1) 向切片拷贝内容

*/
func slice() {
	// make
	s := make([]string, 3)
	fmt.Println("emp:", s)
	// set
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	//get
	fmt.Println("get:", s[2])
	//len
	fmt.Println("len:", len(s))
	// append
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	// copy
	copy(c, s)
	fmt.Println("cpy:", c)
	// sub
	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)
	// init
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)
	// 二维切片
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}

/*array
数组： 长度和类型都是数组类型的一部分

初始化: {}
获取: []
设置: []=
截取：[:]
长度：len
遍历： for,range

多维数组：[][]
数组函数传参是值传递，引用传递通过&
*/
func array() {
	// define
	var a [5]int
	fmt.Println("empty default values:", a)
	// set
	a[4] = 100
	fmt.Println("set:", a)
	//get
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))
	// init
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("初始化:", b)

	var twoD [2][3]int
	// 遍历
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	type user struct {
		name string
		age  byte
	}
	// 初始化
	users := [...]user{
		{"cc", 26},
		{"wx", 20},
	}
	// sub
	sub := users[:1]
	fmt.Println("origin=", users, "sub=", sub)
}
