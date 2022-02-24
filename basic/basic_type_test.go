package basic

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func TestBasicType(t *testing.T) {
	// bool
	var flag bool
	flag = true
	fmt.Println(flag)
	// byte
	var b byte = 1
	fmt.Println(b)
	var bytes []byte = []byte("/usr/bin/tso")
	fmt.Println("bytes", bytes)

	// 复数
	var c64 complex64 = 1 + 2i
	var c128 complex128 = 1 + 2i
	fmt.Println(c64)
	fmt.Println(c128)
	// int
	// uint
	// float
	// unitptr
	testNumber()

	// 结构体
	p := person{name: "gowb", age: 20}
	fmt.Println("struct packet visible ", p)

	// 类型转换
	testTypeConvert()
}

func testTypeConvert() {
	// 断言类型,接口断言 var s = x.(T)
	// 如果 T 不是接口类型，则要求 x 的类型就是 T，如果 T 是一个接口，要求 x 实现了 T 接口

	// switch type

	// T(x) 直接类型隐式转换

}

/*
	整数：int、uint 和 uintptr(指针)。
	int 和 uint 在 32 位操作系统上，它们均使用 32 位（4 个字节），在 64 位操作系统上，它们均使用 64 位（8 个字节）
	uintptr 的长度被设定为足够存放一个指针即可
int8（-128 -> 127）
int16（-32768 -> 32767）
int32（-2,147,483,648 -> 2,147,483,647）
int64（-9,223,372,036,854,775,808 -> 9,223,372,036,854,775,807）

uint8（0 -> 255）
uint16（0 -> 65,535）
uint32（0 -> 4,294,967,295）
uint64（0 -> 18,446,744,073,709,551,615）


*/
func testNumber() {
	// 有符号整数
	var i int = 1
	var i8 int8 = 1
	var i16 int16 = 1
	var i32 int32 = 1
	var i64 int64 = 1
	var ui64 uint64 = 18446744073709551615
	fmt.Println(i)
	fmt.Println(i8)
	fmt.Println(i16)
	fmt.Println(i32)
	fmt.Println(i64)
	fmt.Println(ui64)

	var f32 float32 = 12.2121
	var f64 float64 = 12.233242432
	fmt.Println(f32)
	fmt.Println(f64)

	testUintptr()
}

func testUintptr() {
	//创建一个变量
	var i int8 = 10
	i2 := &i
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.TypeOf(i2))

	//创建一个变量转化成Pointer 和 uintptr
	p := unsafe.Pointer(&i) //入参必须是指针类型的
	fmt.Println(p)          //是内存地址
	u := uintptr(i)
	// 取地址的值
	fmt.Println(u) //结果就是10

	//Pointer转换成uintptr
	temp := uintptr(p)
	fmt.Println(temp)
	//uintptr转Pointer
	p = unsafe.Pointer(u)
	//获取指针大小
	u = unsafe.Sizeof(p) //传入指针，获取的是指针的大小
	fmt.Println(u)       // 打印u是：8
	//获取的是变量的大小
	u = unsafe.Sizeof(i)
	fmt.Println(u) //打印u是：1
}
