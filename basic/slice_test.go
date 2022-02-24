package basic

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"testing"
)

/*
	切片（slice）是对数组一个连续片段的引用

	切片是一个引用类型
	给定项的切片索引可能比相关数组的相同元素的索引小。和数组不同的是，切片的长度可以在运行时修改，最小为 0 最大为相关数组的长度：切片是一个 长度可变的数组
	如果 s 是一个切片，cap(s) 就是从 s[0] 到数组末尾的数组长度。切片的长度永远不会超过它的容量
	多个切片如果表示同一个数组的片段，它们可以共享数据

	相较数组优点切片是引用，所以它们不需要使用额外的内存并且比使用数组更有效率，并且更加灵活，能共享数据又能类似数组操作
	声明切片的格式是： var identifier []type（不需要说明长度）
	切片在未初始化之前默认为 nil，长度为 0
	切片在内存中的组织方式实际上是一个有 3 个域的结构体：指向相关数组的指针，切片长度以及切片容量
	切片的容量是从数组的开始下标到数组结束的长度
	注意一般不要用指针指向 slice。切片本身已经是一个引用类型
	当需要对切片进行修改的时候可以使用指针进行操作，比如改变slice长度，

	append方法当slice没有超过容量，返回slice，否则重新创建slice返回
	copy 方法将src和dest中容取容量小值进行赋值
	切片和垃圾回收
	切片的底层指向一个数组，该数组的实际容量可能要大于切片所定义的容量。只有在没有任何切片指向的时候，底层的数组内存才会被释放，这种特性有时会导致程序占用多余的内存

	数组转切片
	a[low:high：max]
	切片转数组指针：得到的是底层数组的指针，不会出现内存拷贝
*/

func TestArraySliceConvert(t *testing.T) {
	ints := [5]int{1, 2, 3, 4, 5}
	slice := ints[1:2:3]
	fmt.Println("array[low:high：max] ", slice, len(slice), cap(slice))
	slice1 := ints[0:2]
	fmt.Println("array[low:high] ", slice1, len(slice1), cap(slice1))

	// 切片转数组
	slice3 := make([]byte, 2, 4)
	fmt.Println("slice make ", slice3, len(slice3), cap(slice3))

	// 数组指针类型,得到的是底层数组的指针，不会出现内存拷贝
	s0 := (*[0]byte)(slice3)
	s2 := (*[2]byte)(slice3)
	// s4 := (*[4]byte)(slice3) 切片长度不能小于数组长度，否则会panic
	fmt.Println("slice covert *array", s0, s2)

	// 将一个空slice转为0长度数组指针得到nil,非空slice转为0长度数组不是nil
	var ts []string
	t0 := (*[0]string)(ts)
	//t1 := (*[1]string)(ts)
	fmt.Println(t0)
}

var digitRegexp = regexp.MustCompile("[0-9]+")

func FindDigits(filename string) []byte {
	b, _ := ioutil.ReadFile(filename)
	//这种方式会导致底层使用的整个文件bytes
	//return digitRegexp.Find(b)
	b = digitRegexp.Find(b)
	// 新建一个[]byte 存储数字切片
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
func FindFileDigits(filename string) []byte {
	fileBytes, _ := ioutil.ReadFile(filename)
	b := digitRegexp.FindAll(fileBytes, len(fileBytes))
	c := make([]byte, 0)
	for _, bytes := range b {
		c = append(c, bytes...)
	}
	return c
}

func TestMemCopy(t *testing.T) {
	fmt.Println("file find digest ", FindDigits("slice_test.go"))
	fmt.Println("file find all digest ", FindFileDigits("slice_test.go"))
}
func TestSlice(t *testing.T) {
	// 切片定义
	var x = []int{2, 3, 5, 7, 11}

	slice0 := x[1:]

	var slice1 []int
	// 数组切片后时引用类型数据
	ints := [5]int{2, 3, 5, 7, 11}
	// {2, 3, 5, 7, 11}
	slice1 = ints[:]
	//{2, 3}
	slice2 := ints[:2]
	//{3, 5, 7, 11}
	slice3 := ints[1:]
	//{2, 3, 5, 7}
	slice4 := ints[0:4]

	//make和new 创建切片,len=cap=10
	slice5 := make([]int, 10)

	fmt.Println("slice init", x, slice0, slice1, slice2, slice3, slice4, slice5)
	fmt.Println("slice init cap", cap(x), cap(slice0), cap(slice1), cap(slice2), cap(slice3), cap(slice4), cap(slice5))
	fmt.Println("slice init len", len(x), len(slice0), len(slice1), len(slice2), len(slice3), len(slice4), len(slice5))
	// 切片使用

	// 修改slice 数据，数组和共享slice数据都会更改
	slice2[0] = 3
	fmt.Println("slice modify", slice1, slice2, slice3, slice4)
	// 修改一个没有初始化的slice
	slice5[0] = 4
	fmt.Println("slice modify", slice5)
	fmt.Println("slice modify len", len(slice5))
	var arr = [5]int{0, 1, 2, 3, 4}
	fmt.Println("sum ", sum(arr[:]))

	// 判断slice为空
	//绝对不能用 if slice == nil 这样的方式
	// &[], 一个slice的引用，不推荐使用
	slice6 := new([]int)
	fmt.Println("new slice ", slice6)
	if len(*slice6) == 0 {
		fmt.Println("slice6 is null", slice6)
	}
	// nil slice
	var slice7 []int
	if len(slice7) == 0 {
		fmt.Println("slice7 is null", slice6)
	}
	// empty slice
	slice8 := []int{}
	if len(slice8) == 0 {
		fmt.Println("slice8 is null", slice6)
	}

	// 切片扩容
	slice8 = append(slice8, 1)
	fmt.Println("slice append 扩容", slice8, cap(slice8), len(slice8))
	// copy 复制
	slice9 := make([]int, 10)
	copy(slice9, slice8)
	fmt.Println("slice copy", slice9)
	// slice 作为函数参数
	var buffer [256]byte
	slice := buffer[10:20]
	for i := 0; i < len(slice); i++ {
		slice[i] = byte(i)
	}
	fmt.Println("slice param before", slice)
	AddOneToEachElement(slice)
	fmt.Println("slice param after", slice)

	// 修改切片长度:利用切片指针
	fmt.Println("slice ptr Before: len(slice) =", len(slice))
	PtrSubtractOneFromLength(&slice)
	fmt.Println("slice ptr After:  len(slice) =", len(slice))
	// 多维切片

	// for forr
	for i := range slice9 {
		fmt.Printf("%d->%d", i, slice9[i])
	}
	fmt.Println()
	for i, i2 := range slice9 {
		fmt.Printf("%d->%d", i, i2)
	}
	fmt.Println()

	testMultiDimensionSlice(slice1)

}

func PtrSubtractOneFromLength(slicePtr *[]byte) {
	slice := *slicePtr
	*slicePtr = slice[0 : len(slice)-1]
}
func AddOneToEachElement(slice []byte) {
	for i := range slice {
		slice[i]++
	}
}
func testMultiDimensionSlice(slice1 []int) {
	/*
	   var sliceName[][][]...[]SliceType
	*/
	//声明一个二维切片
	var slice [][]int
	//为二维切片赋值
	slice = [][]int{{10}, {100, 200}}
	//声明并赋值
	slice7 := [][]int{{10}, {100, 200}}
	//为第一个切片追加值为20的元素
	slice7[0] = append(slice7[0], 20)
	fmt.Println(slice1)
	fmt.Println(slice)
}
func sum(a []int) (s int) {
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return
}
