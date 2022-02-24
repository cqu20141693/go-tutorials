package basic

import (
	"fmt"
	"testing"
)

/*
	容器 , 它是可以包含大量条目（item）的数据结构，例如数组、切片和 map

	在 Go 程序中并不经常看到数组，因为数组的大小是数组类型的一部分，这限制了数组的表达能力。
	数组是具有相同 唯一类型 的一组以编号且长度固定的数据项序列
	数组元素类型为任意类型的可以使用空接口作为类型
	数组元素可以通过 索引（位置）来读取（或者修改）
	数组必须初始化时指定数组容量
	声明格式： var identifier [len]type
	数组支持for forr 遍历方式
	Go 语言中的数组是一种 值类型,可以通过 new() 来创建： var arr1 = new([5]int)
	当直接使用数组遍历赋值时是拷贝，需要使用指针进行传递才可以进行

	数组的优势：
	支持随机访问，内存连续，
	不支持动态扩容，且容量一定，对于数据少于容量情况下浪费内存
	对于插入和删除支持时间复杂度为N
*/
const (
	WIDTH  = 5
	HEIGHT = 5
)

type pixel int

var screen [WIDTH][HEIGHT]pixel

func TestArray(t *testing.T) {
	// 数组常量,初始化
	ints := [5]int{1, 2, 3, 4, 5}
	strings := [5]string{3: "name", 4: "age"}
	fmt.Println("ints=", ints)
	fmt.Println("strings=", strings)

	// new 创建数组,是一个指针类型变量
	array1 := new([100]int)
	fmt.Println("new array ", array1)
	fmt.Println("new array[1] ", array1[1])

	var arr1 [5]int
	// forr 遍历
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i * 2
	}
	for i := range arr1 {
		fmt.Printf("Array at index %d is %d\n", i, arr1[i])
	}

	for i, value := range arr1 {
		fmt.Printf("Array at index %d is %d\n", i, value)
	}

	// 多维数组
	MultidimensionalArrays()
	fmt.Println()
	arrayParameter()
}

func arrayParameter() {
	ints := [5]int{1, 2, 3, 4, 5}
	ints2 := ints
	ints2[0] = 2
	fmt.Println("ints=", ints)
	fmt.Println("ints2=", ints2)

	array := [3]float64{7.0, 8.5, 9.1}
	x := Sum(&array) // Note the explicit address-of operator
	// to pass a pointer to the array
	fmt.Printf("The sum of the array is: %f", x)

}
func Sum(a *[3]float64) (sum float64) {
	for _, v := range *a { // derefencing *a to get back to the array is not necessary!
		sum += v
	}
	return
}

func MultidimensionalArrays() {
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			screen[x][y] = 0
		}
	}
	for i := range screen {
		for j := range screen[i] {
			fmt.Printf(" %d", screen[i][j])
		}
	}
}
