package main

import (
	"fmt"
)

/*定义一个结构体*/
type person struct {
	Name string
	Age  int
}

/*结构嵌套*/
type school struct {
	Name    string
	Address string
	No      int
	Contact struct {
		Phone string
		City  string
	}
}

/*匿名字段*/
type teacher struct {
	string
	int
}

/*结构组合*/
type human struct {
	Sex int
}
type students struct {
	human
	string
	int
}

func main() {
	/*初始化*/
	a := person{}
	/*操作结构体属性*/
	a.Name = "GO"
	a.Age = 19
	/*第二种方式*/
	b := person{
		Name: "GO",
		Age:  20,
	}
	/*使用引用*/
	c := new(person)
	c.Name = "KKKKKK"
	/*匿名结构*/
	d := struct {
		Name string
		Age  int
	}{
		Name: "GOGOG",
		Age:  55,
	}
	/*嵌套*/
	e := school{}
	e.Address = "大坪"
	e.Name = "大坪小学"
	e.No = 1
	e.Contact.City = "重庆"
	e.Contact.Phone = "12546542"
	fmt.Println("sturct")
	/*匿名字段*/
	f := teacher{"老师", 50}
	/*结构组合*/
	s := students{human{Sex: 1}, "我是学生", 14}
	s.Sex = 10
	/*值拷贝*/
	fmt.Println(a)
	A(a)
	fmt.Println(a)
	fmt.Println(b)
	/*引用传递*/
	fmt.Println(c)
	B(c)
	fmt.Println(c)
	/*匿名结构*/
	fmt.Println(d)
	/*嵌套结构*/
	fmt.Println(e)
	/*匿名字段*/
	fmt.Println(f)
	/*结构组合*/
	fmt.Println(s)
}

/*判断sturct是引用传递还是值传递*/

func A(per person) {
	per.Age = 15
	fmt.Println("A", per)
}
func B(per *person) {
	per.Age = 100
	fmt.Println("B", per)
}
