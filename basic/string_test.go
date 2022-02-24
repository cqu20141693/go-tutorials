package basic

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

/*
	string类型一种特殊的slice类型，可以使用内置函数len获取长度，同时支持 切片操作
len(s): 获取字符串长度
比较符： > < == != 可直接使用
    strings 包：

//   EqualFold 函数，计算 s 与 t 忽略字母大小写后是否相等。

// 子串 substr 在 s 中，返回 true
func Contains(s, substr string) bool

// chars 中任何一个 Unicode 代码点在 s 中，返回 true
func ContainsAny(s, chars string) bool

// Unicode 代码点 r 在 s 中，返回 true
func ContainsRune(s string, r rune) bool

在内存中，一个字符串实际上是一个双字结构，即一个指向实际数据的指针和记录字符串长度的整数

*/

func Compare(a, b []byte) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		switch {
		case a[i] > b[i]:
			return 1
		case a[i] < b[i]:
			return -1
		}
	}
	// 数组的长度可能不同
	switch {
	case len(a) < len(b):
		return -1
	case len(a) > len(b):
		return 1
	}
	return 0 // 数组相等
}
func TestStringCompare(t *testing.T) {
	fmt.Println("compare=", Compare([]byte("qwe"), []byte("qqwq")))
}

func TestStrings(t *testing.T) {
	testCompareAndEqual()
	// ASCII 一个字节，Unicode 字符会占用 2 个字节，有些甚至需要 3 个或者 4 个字节来进行表示
	s := "你好中国"
	runes := []rune(s)
	fmt.Println("runes len=", len(runes))
	for i := range runes {
		fmt.Printf("runes %d=%c\n", i, runes[i])
	}
	bytes := []byte(s)
	fmt.Println("bytes len=", len(bytes))
	for i := range bytes {
		fmt.Printf("bytes %d=%c\n", i, bytes[i])
	}
	// 获取字符串的某一部分,截取的是底层bytes slice
	s2 := s[0:3]
	s3 := s[1:4]
	fmt.Println("string [:] type ", reflect.TypeOf(s2), reflect.TypeOf(s3))
	fmt.Println("string sub ", s2, s3)
}

/*
	Compare 函数，用于比较两个字符串的大小，如果两个字符串相等，返回为 0。如果 a 小于 b ，返回 -1 ，反之返回 1 。
	EqualFold: 比较字符串，大小写不敏感
	不推荐使用这个函数，直接使用 == != > < >= <= 等一系列运算符
*/
func testCompareAndEqual() {
	a := "gopher"
	b := "hello world"
	fmt.Println(strings.Compare(a, b))
	fmt.Println(strings.Compare(a, a))
	fmt.Println(strings.Compare(b, a))
	fmt.Println(a > b)
	fmt.Println(a == a)
	fmt.Println(a < b)

	fmt.Println(strings.EqualFold("GO", "go"))
	fmt.Println("GO" == "go")
	fmt.Println(strings.EqualFold("GO", "GO"))
	fmt.Println(strings.EqualFold("壹", "一"))

}

func TestStringNull(t *testing.T) {
	var s string

	if s == "" {
		fmt.Println("this is a recommended way")
	}
	if len(s) > 0 {
		fmt.Println("this is another way")
	}
}

func TestIntToString(t *testing.T) {

	var i int = 12
	var s10 string = strconv.FormatInt(int64(i), 10)
	var s2 string = strconv.FormatInt(int64(i), 2)
	fmt.Println("int to String ", s10, s2)
}
