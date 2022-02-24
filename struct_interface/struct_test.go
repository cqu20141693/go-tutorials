package struct_interface

import (
	"fmt"
	"github.com/cqu20141693/go-tutorials/basic"
	"reflect"
	"runtime"
	"testing"
)

/*
结构体是一种聚合类型，里面可以包含任意类型的值，这些值就是我们定义的结构体的成员，也称为字段。在 Go 语言中，要自定义一个结构体，需要使用 type+struct 关键字组合。
*/
type person struct {
	name string

	age uint
}
type Student struct {
	person
}

/*
对象继承
*/
func TestExtends(t *testing.T) {
	stu := Student{person{"Panda", 20}}
	age := stu.age
	name := stu.name
	fmt.Println(stu, age, name)
}

func TestStruct(t *testing.T) {
	// person 不可见    new(person)
	person := basic.NewPerson("gowb", 23)
	fmt.Println("person", person)

	number := basic.NewNumber(10.0)
	b := new(basic.Number)
	fmt.Println("Number ", number, b)
}

func TestStructTag(t *testing.T) {
	tagType := basic.NewTagType(true, "gowb", 1)
	ttType := reflect.TypeOf(tagType)
	numField := ttType.NumField()
	for i := 0; i < numField; i++ {
		refTag(ttType, i)
	}
}

func refTag(ttType reflect.Type, index int) {

	field := ttType.Field(index)
	fmt.Println("struct tag ", field.Tag)
}

/*
	go 结构体通过匿名field 组合实现继承能力，具有匿名类型的能力
*/
func TestInnerFieldAndStruct(t *testing.T) {
	outer := new(basic.OuterS)
	outer.SetB(6)
	outer.SetC(7.5)
	outer.SetInt(60)
	outer.SetIn1(5)
	outer.SetIn2(10)

	fmt.Printf("outer.b is: %d\n", outer.B())
	fmt.Printf("outer.c is: %f\n", outer.C())
	fmt.Printf("outer.int is: %d\n", outer.GetInt())
	fmt.Printf("outer.in1 is: %d\n", outer.In2())
	fmt.Printf("outer.in2 is: %d\n", outer.In2())

	// 使用结构体字面量
	outer2 := basic.NewOuterS(6, 7.5, 60, basic.NewInnerS(5, 10))
	fmt.Println("outer2 is:", outer2)
}

type OuterAlias = basic.OuterS

type OuterExtend basic.OuterS

func (o *OuterExtend) extendFunc() {
	fmt.Printf("struct alias extend %v", o)
}

/*
	别名类型扩展
*/
func TestAliasExtend(t *testing.T) {
	outer := new(OuterAlias)
	outer.SetB(6)
	outer.SetC(7.5)
	outer.SetInt(60)
	outer.SetIn1(5)
	outer.SetIn2(10)
	fmt.Printf("OuterAlias is OutetS Alias :%v", outer)

	outer2 := new(OuterExtend)
	outer2.extendFunc()

}

func TestStructMeth(t *testing.T) {

	shop := basic.NewShop(10)
	shop2 := basic.NewShop(20)
	fmt.Println("Shop compare result=", shop2.CompareTo(shop))
}

func TestStructString(t *testing.T) {
	shop := basic.NewShop(10)
	runtime.SetFinalizer(shop, func(w *basic.Shop) {
		w.Finalizer()
	})
	fmt.Println("shop String", shop)
	fmt.Printf("shop String %v \n", shop)
	fmt.Printf("shop String %T\n", shop)
	fmt.Printf("shop String %#v\n", shop)

}
