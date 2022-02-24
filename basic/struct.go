package basic

import (
	"fmt"
	"strconv"
)

/*
	结构体依然遵循可见性原则
	需要通过工厂方法暴露器构造器 New
	自定义结构使用是通过模块间导入
	结构体field 具有标签能力，通过reflect 访问
	结构体具有匿名filed ，匿名结构体具有组合能力，具有OOP 继承特性

	Meth
	方法定义需要指定结构体 ，并遵循可见性原则，方法和类型必须在同一个包
	当方法中不适用对象时，可以用_ 代替 ，该对象具有OOP中的this特性
	方法必须有一个接收者，可以是所有的类型，但是不能是接口

	func
	String() 打印输出方法

	对象垃圾回收
	调用 runtime.GC() 函数可以显式的触发 GC
	SetFinalizer: 对象 obj 被从内存移除前执行一些特殊操作，比如写到日志文件中
	runtime.SetFinalizer(obj, func(obj *typeObj))

	如果要包外扩展方法，可以定义别名类型，对别名类型新增方法
*/

type person struct {
	name string
	age  int8
}

func (p *person) Name() string {
	return p.name
}

func (p *person) SetName(name string) {
	p.name = name
}

func (p *person) Age() int8 {
	return p.age
}

func (p *person) SetAge(age int8) {
	p.age = age
}

type struct1 struct {
	i1  int
	f1  float32
	str string
}

func NewPerson(name string, age int8) *person {
	return &person{name, age}
}

type Number struct {
	f float32
}

func NewNumber(f float32) *Number {
	return &Number{f: f}
}

type nr Number // alias type

type TagType struct {
	answer bool   "An important answer"
	name   string "The name of the thing"
	price  int    "How much there are"
}

func NewTagType(answer bool, name string, price int) TagType {
	return TagType{answer: answer, name: name, price: price}
}

type InnerS struct {
	in1 int
	in2 int
}

func NewInnerS(in1 int, in2 int) InnerS {
	return InnerS{in1: in1, in2: in2}
}

func (i *InnerS) In1() int {
	return i.in1
}

func (i *InnerS) SetIn1(in1 int) {
	i.in1 = in1
}

func (i *InnerS) In2() int {
	return i.in2
}

func (i *InnerS) SetIn2(in2 int) {
	i.in2 = in2
}

type OuterS struct {
	b      int
	c      float32
	int    // anonymous field
	InnerS // anonymous struct
}

func NewOuterS(b int, c float32, int int, innerS InnerS) *OuterS {
	return &OuterS{b: b, c: c, int: int, InnerS: innerS}
}
func (o *OuterS) NewOuterS(b int, c float32, int int, innerS InnerS) *OuterS {
	return &OuterS{b: b, c: c, int: int, InnerS: innerS}
}

func (o *OuterS) B() int {
	return o.b
}

func (o *OuterS) SetB(b int) {
	o.b = b
}

func (o *OuterS) C() float32 {
	return o.c
}

func (o *OuterS) SetC(c float32) {
	o.c = c
}

func (o *OuterS) SetInt(i int) {
	o.int = i
}

func (o *OuterS) GetInt() int {
	return o.int
}

//func (o * OuterS) SetIn1(in1 int)  {
//	o.in1=in1
//}
//
//func (o * OuterS) SetIn2(in2 int)  {
//	o.in2=in2
//}

//func (o *OuterS) GetIn1() int  {
//	return o.in1
//}
//func (o *OuterS) GetIn2() int {
//	return o.in2
//}

type Shop struct {
	price int
}

func NewShop(price int) *Shop {
	return &Shop{price: price}
}

func (s Shop) CompareTo(shop2 *Shop) bool {
	return s.price > shop2.price
}

func (s Shop) String() string {
	return "shop=(price=" + strconv.FormatInt(int64(s.price), 10) + ")"
}

func (s Shop) Finalizer() {
	fmt.Println("shop is recover")
}

type Employee struct {
	name   string
	age    int
	salary float32
}

func (e *Employee) Name() string {
	return e.name
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func (e *Employee) Age() int {
	return e.age
}

func (e *Employee) SetAge(age int) {
	e.age = age
}

func (e *Employee) Salary() float32 {
	return e.salary
}

func (e *Employee) SetSalary(salary float32) {
	e.salary = salary
}
