package __18

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"unsafe"
)

/*
https://mp.weixin.qq.com/s/tjHOd6jvGj7tpmf1K4wlYg

在 Go 中，string 类型的底层表示如下：
 type string struct {
 ptr unsafe.Pointer
 len int
}
*/
func TestClone(t *testing.T) {
	s := "abcdefghijklmn"
	s1 := s[:4]
	// 而 reflect.StringHeader 结构是对字符串底层结构的反射表示。
	sHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))
	s1Header := (*reflect.StringHeader)(unsafe.Pointer(&s1))
	// 子串和原字符串的 Data 部分指向相同的内存，因此整个字符串并不会被 GC 回收。
	fmt.Println(sHeader.Len == s1Header.Len)
	fmt.Println(sHeader.Data == s1Header.Data)
	//strings.Clone 函数就是为了解决这个问题的
	//它保证将 s 复制到一个新分配的副本中,对于长度为零的字符串，将返回字符串 ""，不进行内存分配。
	// 底层实现为make开辟新空间，copy方法进行复制
	s2 := strings.Clone(s[:4])
	s2Header := (*reflect.StringHeader)(unsafe.Pointer(&s2))
	fmt.Println(sHeader.Len == s2Header.Len)
	fmt.Println(sHeader.Data == s2Header.Data)

	title := strings.Title("i love you")
	fmt.Println("title", title)

	addrParse()
}

func addrParse() {
	addr := "192.168.1.1:8080"
	pos := strings.Index(addr, ":")
	pos = strings.LastIndex(addr, ":")

	if pos == -1 {
		panic("非法地址")
	}
	ip, port := addr[:pos], addr[pos+1:]
	fmt.Println("addr=", ip, "ip=", port)

	ok := false
	ip, port, ok = strings.Cut(addr, ":")
	if ok {
		fmt.Println("addr=", ip, "ip=", port)
	} else {
		fmt.Println("addr=", addr, " is not a ip address")
	}
}
