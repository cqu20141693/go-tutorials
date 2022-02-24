package basic

import (
	"fmt"
	"math"
	"testing"
)

/*
	口定义了一组方法（方法集），但是这些方法不包含（实现）代码：它们没有被实现（它们是抽象的）。接口里也不能包含变量。
	接口的名字由方法名加 [e]r 后缀组成，例如 Printer、Reader、Writer、Logger、Converter 等等。还有一些不常用的方式（当后缀 er 不合适时），比如 Recoverable，此时接口名以 able 结尾，或者以 I 开头
	类型不需要显式声明它实现了某个接口，但是必须实现全部的接口方法
	类型的变量赋值给一个接口类型的变量

	多态 的 Go 版本:同一种类型在不同的实例上似乎表现出不同的行为
	接口嵌套接口: 一个接口可以包含一个或多个其他的接口，这相当于直接将这些内嵌接口的方法列举在外层接口中


 	类型断言：如何检测和转换接口变量的类型
	v := varI.(T)       // varI 必须是一个接口变量

	类型判断：type-switch,在 type-switch 不允许有 fallthrough
	测试一个值是否实现了某个接口: sv, ok := v.(Stringer)

	方法接收者指针不能使用值类型进行调用

	在接口上调用方法时，必须有和方法定义时相同的接收者类型或者是可以从具体类型 P 直接可以辨识的：
	指针方法可以通过指针调用
	值方法可以通过值调用
	接收者是值的方法可以通过指针调用，因为指针会首先被解引用
	接收者是指针的方法不可以通过值调用，因为存储在接口中的值没有地址

	Sort 包中的排序算法，提供了排序算法需要的行为接口，算法实现逻辑公用实现，对于其他要排序的类型可以进行类似实现

	空接口
	空接口或者最小接口 不包含任何方法，它对实现不做任何要求,任何其他类型都实现了空接口，可以理解为基类.
	每个 interface {} 变量在内存中占据两个字长：一个用来存储它包含的类型，另一个用来存储它包含的数据或者指向数据的指针

	类型可以先出现方法，如果出现相同的类型时，可以定义接口抽象处理；并且某个类型要新增接口能力，只需要给类型实现新的方法。
	当接口出现冲突时，可以定义ImplementsXXX()接口标识实现XXX接口
*/

type Shaper interface {
	Area() float32
	Name() string
}

type Square struct {
	side float32
}

func (sq *Square) Name() string {
	return "square"
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

type Rectangle struct {
	length float32
	width  float32
}

type Circle struct {
	radius float32
}

func (ci Circle) Area() float32 {
	return ci.radius * ci.radius * math.Pi
}

func (c Circle) Name() string {
	return "Circle"
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

func (r Rectangle) Name() string {
	return "rectangle"
}

type Drawer interface {
	draw(shaper Shaper)
}

type IDrawBoard interface {
	Drawer
	Close()
}

type Computer struct {
	name string
}

func (c Computer) draw(shaper Shaper) {
	fmt.Printf("draw %s area %f \n", shaper.Name(), shaper.Area())
}

func (c Computer) Close() {
	fmt.Printf("%s computer closed", c.name)
}

func TestInterface(t *testing.T) {
	r := &Rectangle{5, 3} // Area() of Rectangle needs a value
	q := &Square{5}       // Area() of Square needs a pointer

	shapes := []Shaper{r, q}
	fmt.Println("Looping through shapes for area ...")
	for n, shape := range shapes {

		if _, ok := shape.(*Square); ok {
			fmt.Printf("type assert square ")
		}
		if _, ok := shape.(*Rectangle); ok {
			fmt.Printf("type assert Rectangle ")
		}
		fmt.Println("Shape details: ", shapes[n])
		fmt.Printf("The %s has area: %f\n", shapes[n].Name(), shapes[n].Area())
	}

	// 测试嵌套接口
	var drawBoard IDrawBoard
	drawBoard = Computer{"huawei"}
	drawBoard.draw(r)

	// testSwitchType
	testTypeSwitch(r)

}

func testTypeSwitch(shaper Shaper) {
	switch t := shaper.(type) {
	case *Square:
		fmt.Printf("Type Square %T with value %v\n", t, t)
	case *Circle:
		fmt.Printf("Type Circle %T with value %v\n", t, t)
	case nil:
		fmt.Printf("nil value: nothing to check?\n")
	default:
		fmt.Printf("Unexpected type %T\n", t)
	}
}

type List []int

func (l List) Len() int {
	return len(l)
}

func (l *List) Append(val int) {
	*l = append(*l, val)
}

type Appender interface {
	Append(int)
}

func CountInto(a Appender, start, end int) {
	for i := start; i <= end; i++ {
		a.Append(i)
	}
}

type Lener interface {
	Len() int
}

func LongEnough(l Lener) bool {
	return l.Len()*10 > 42
}

func TestValueAndPointerMethod(t *testing.T) {
	// A bare value
	var lst List
	// compiler error:
	// cannot use lst (type List) as type Appender in argument to CountInto:
	//       List does not implement Appender (Append method has pointer receiver)
	// CountInto(lst, 1, 10)
	if LongEnough(lst) { // VALID:Identical receiver type
		fmt.Printf("- lst is long enough\n")
	}

	// A pointer value
	plst := new(List)
	CountInto(plst, 1, 10) //VALID:Identical receiver type
	if LongEnough(plst) {
		// VALID: a *List can be dereferenced for the receiver
		fmt.Printf("- plst is long enough\n")
	}
}

var i = 5
var str = "ABC"

type Person struct {
	name string
	age  int
}

type Any interface{}

func TestEmptyInterface(t *testing.T) {
	var val Any
	val = 5
	fmt.Printf("val has the value: %v\n", val)
	val = str
	fmt.Printf("val has the value: %v\n", val)
	pers1 := new(Person)
	pers1.name = "Rob Pike"
	pers1.age = 55
	val = pers1
	fmt.Printf("val has the value: %v\n", val)
	// 利用type switch 进行类型处理
	switch t := val.(type) {
	case int:
		fmt.Printf("Type int %T\n", t)
	case string:
		fmt.Printf("Type string %T\n", t)
	case bool:
		fmt.Printf("Type boolean %T\n", t)
	case *Person:
		fmt.Printf("Type pointer to Person %T\n", t)
	default:
		fmt.Printf("Unexpected type %T", t)
	}
}

// test interface and struct impl

// 不管接收者类型是值类型还是指针类型，都可以通过值类型或指针类型调用，这里面实际上通过语法糖起作用的
// 类型断言
//<目标类型的值>，<布尔参数> := <表达式>.( 目标类型 ) // 安全类型断言
// <目标类型的值> := <表达式>.( 目标类型 )　　//非安全类型断言

type animal interface {
	say()
}
type dog struct {
}

func (d *dog) say() {
	fmt.Printf("You are dog dog")
}

type people struct {
	animal *animal
}

type cat struct {
}

func (c cat) say() {
	fmt.Println("You are cat cat")
}

func (p *people) walkTheDog() {

	// 安全断言
	animal2, ok := (*p.animal).(*dog)
	if ok {
		animal2.say()
	} else {
		fmt.Println("animal is not a dog")
		animal2.say()
	}

}

func TestInterfacePointer(t *testing.T) {
	d := dog{}
	var a animal = &d
	p := people{
		animal: &a,
	}
	p.walkTheDog()
	c := cat{}
	var a1 animal = &c
	p = people{animal: &a1}
	p.walkTheDog()
}
