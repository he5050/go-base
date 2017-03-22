package main

import (
	"fmt"
)

/*声明结构*/

type A struct {
	Name string
	Age  int
}
type B struct {
	Name string
}
type C struct {
	Name string
	age  int
	sex  int
}

/*定义类型*/
type T int

/*为结构提供方法*/
func (a *A) Print() {
	a.Name = "AA"
	fmt.Println("A")
}
func (b B) Print(str string) {
	b.Name = "BB"
	fmt.Println(str)
}
func (a A) Add(x int, y int) (sum int) {
	return x + y
}
func (a A) Increate(x *int) {
	*x += 100
}
func NewC(name string, age int, sex int) *C {
	return &C{name, age, sex}
}
func (c C) Print() {
	fmt.Println(c.Name, c.age, c.sex)
}

/*方法的继承与重载*/
type Human struct {
	name  string
	age   int
	phone string
}
type Student struct {
	Human
	school string
}
type Employee struct {
	Human
	company string
}

func (h *Human) SayHello() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}
func NewHuman(name string, age int, phone string) *Human {
	return &Human{name, age, phone}
}
func (e *Employee) SayHello() {
	fmt.Printf("Hi, I am %s you can call me on %s my company is %s\n", e.name, e.phone, e.company)
}

/*为类型提供方法*/
func (t *T) Print() {
	fmt.Println("T")
}
func main() {
	a := A{}
	b := B{}
	a.Print()
	sum := a.Add(10, 100)
	b.Print("我要到方法里面去了哦")
	fmt.Println("sum的值为:", sum)
	/*我们的结构体为值类型*/
	fmt.Println(a)
	fmt.Println(b)
	/*类型添加方法*/
	var t T
	t.Print()
	/*嵌入式*/
	var t1 T
	(*T).Print(&t1)
	d := 0
	a.Increate(&d)
	fmt.Println(d)
	/*初始化*/
	c := *NewC("我要初始化了", 18, 1)
	c1 := NewC("我要初始化了", 20, 0)
	fmt.Println(c)
	fmt.Println(c1)
	c.Print()
	c1.Print()
	/*继承*/
	h1 := NewHuman("Mark", 25, "123456")
	h2 := NewHuman("Jum", 26, "123456789")
	e := Employee{Human{"Tim", 28, "000000000"}, "jsbn"}
	h1.SayHello()
	h2.SayHello()
	/*重载/重写*/
	e.SayHello()
}
