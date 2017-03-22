## 方法(method)

Go 语言中同时有函数和方法。一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或者是一个指针。所有给定类型的方法属于该类型的方法集
声明:
`func (struct_name struct_type) func_name(var_name var_type) (return_type)`
> 上面的方法的声明与函数的声明差不多,只是在函数名前面添加一个参数,就是对应结构体名与结构体的类型

```

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
```

结果:
```
A
我要到方法里面去了哦
sum的值为: 110
{AA 0}
{}
T
T
100
{我要初始化了 18 1}
&{我要初始化了 20 0}
我要初始化了 18 1
我要初始化了 20 0
Hi, I am Mark you can call me on 123456
Hi, I am Jum you can call me on 123456789
Hi, I am Tim you can call me on 000000000 my company is jsbn
```

1. 虽然method的名字一模一样，但是如果接收者不一样，那么method就不一样
2. method里面可以访问接收者的字段(方法可以直接方位该类型或是结构当中的数据)
3. 访求也是使用是值传递
4. 方法的声明与函数声明一样,只是在函数名前面多了一个类型
5. `(*T).Print(&t1)`,可以直接使用引用方法进行操作。注意这里必须是指针类型,不是变量,如果变量那就相当于重新开辟了一块内存空间
5. 一般使用NewXXX来定义一个函数用于初始化
6. 方法是可以继承的(也可以叫做匿名组合)
7. 方法是可以用来重写的
8. 存在继承关系时，按照就近原则，进行调用
