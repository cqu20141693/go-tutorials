package main

import (
	"fmt"
	"math"
)

// 结构体

func StructAndInterface() {

	fmt.Println("..........struct_interface................")
	structs()

	interfaces()
}

type geometry interface {
	area() float64
	perim() float64
}
type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

/*interfaces

以指针类型接收者实现接口，只有对应的指针类型才被认为实现了接口
*/
func interfaces() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	// 其中一个方法的接收者是指针
	measure(&r)
	measure(c)
}

/*structs

结构体：type name struct{} ： 数据集合，值类型
可见性： 首字母大写可见
初始化： {}
method: receiver 类型为结构体类型,Go自动处理方法调用的值和指针之间的转换。您可能希望使用指针接收器类型来避免在方法调用中进行复制，或允许方法改变接收结构(interface)。
.： 获取属性和调用方法
结构体嵌套： 结构体字段为结构体
*/

type rect struct {
	width, height float64
}

func (r *rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func structs() {
	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
