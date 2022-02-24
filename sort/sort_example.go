package main

import (
	"sort"
)

/*
	标准库提供了 sort 包来实现常见的搜索和排序操作
	sort 包中的函数 func Ints(a []int) 来实现对 int 类型的切片排序
	函数 IntsAreSorted(a []int) bool 来检查数组是否已经被排序
 	func Float64s(a []float64) 来排序 float64 的元素
	数组或切片中搜索一个元素，该数组或切片必须先被排序（因为标准库的搜索算法使用的是二分法
	函数 func SearchInts(a []int, n int) int 进行搜索，并返回对应结果的索引值
	func SearchFloat64s(a []float64, x float64) int
	func SearchStrings(a []string, x string) int
*/

func main() {
	data := []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	var a sort.IntSlice = data[:] //conversion to type IntArray
	sort.Sort(a)

}
