package basic

import (
	"fmt"
	"testing"
)

/*
	结构体定义
	结构体是值类型，因可以通过 new 函数来创建
	type T struct {a, b int}

	 // 构造函数
	T{a:1,b:1} : 使用{}定义结构体变量，值类型,&获取指针
	new(T): 返回指针类型

	结构体的内存布局
	Go 语言中，结构体和它所包含的数据在内存中是以连续块的形式存在的，即使结构体中嵌套有其他的结构体，这在性能上带来了很大的优势。
	如果是引用类型的结构体或者map,存放的是指针，内存不连续
*/

func TestStruct(t *testing.T) {
	// var t *T = new(T) 变量 t 是一个指向 T 的指针 等同于&T
	p := new(person)
	//var t T 也会给 t 分配内存，并零值化内存，t 是类型 T的变量
	var p1 person
	fmt.Println("struct new ", p, p1)
	p.name = "cc"
	p.age = 100
	p1.name = "cqu"
	fmt.Println("struct new init", p, p1)

	// 初始化结构体
	p2 := person{"gowb", 12}
	// 结构体变量
	p3 := person{age: 10, name: "wq"}

	fmt.Println("struct init : ", p2, p3)
	// 指针变量
	p4 := &person{age: 20}
	fmt.Println("struct init : ", p4)

	ms := new(struct1)
	ms.i1 = 10
	ms.f1 = 15.5
	ms.str = "Chris"

	fmt.Printf("The int is: %d\n", ms.i1)
	fmt.Printf("The float is: %f\n", ms.f1)
	fmt.Printf("The string is: %s\n", ms.str)
	fmt.Println(ms)

	// 结构体类型转换
	n := Number{5.0}
	n2 := nr{6.0}
	var c = Number(n2)
	fmt.Println("struct conversion: ", n, n2, c)

	testAlias()
}

type e Employee

func testAlias() {

	e2 := e{"gowb", 12, 100000}
	employee := Employee(e2)
	fmt.Println("struct alias", employee)
}
