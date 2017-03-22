## 结构体 struct

Go 语言中数组可以存储同一类型的数据，但在结构体中我们可以为不同项定义不同的数据类型。 结构体是由一系列具有相同类型或不同类型的数据构成的数据集合。
```
type school struct{
    name string
    class int
    students int
    teacher int
    address string
}
```
> 能过上面我们简单的定义了一个结构体,是不是有点像类的操作

1. 简单实例
```
/*定义一个结构体*/
type person struct {
	Name string
	Age  int
}


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
```
2. 嵌套
```

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

/*嵌套初始化*/
e := school{}
e.Address = "大坪"
e.Name = "大坪小学"
e.No = 1
e.Contact.City = "重庆"
e.Contact.Phone = "12546542"
```
3. 匿名字段

```
/*匿名字段*/
type teacher struct {
	string
	int
}

/*匿名字段*/
f := teacher{"老师", 50}
```
>  若存在匿名字段中的字段与非匿名字段名字相同，则最外层的优先访问，就近原则


4. 匿名结构

```
/*匿名结构*/
d := struct {
    Name string
    Age  int
}{
    Name: "GOGOG",
    Age:  55,
}
```
5. 结构的引用

```
/*使用引用*/
type person sturct{
    Name string
}
/*-------*/

d : = &person{Name : "KKKKK"}
c := new(person)
c.Name = "KKKKKK"
```

6. 结构组合

```
/*结构组合*/
type human struct {
	Sex int
}
type students struct {
	human
	string
	int
}

s := students{human{Sex: 1}, "我是学生", 14}

s.Sex = 10
```
> 这里要提示一下使用结构组合与结构嵌套有一个区别事,当你使用组合的,在这个结构当中的,所有字段都可以直接使用`s.Sex = 10`

7. 结构作为函数值

```
func A(per person) {
	per.Age = 15
	fmt.Println("A", per)
}
func B(per *person) {
	per.Age = 100
	fmt.Println("B", per)
}
```
8. 补充说明,关于类型的定义
```
/*声明一个类型*/
type Interger int
/*为类型添加一个方法,相当于是接口*/
func (a Interger) Less (b Interger) bool {
    return a<b
}
/*使用*/
var a Interger = 1
a.Less(2){
    fmt.Println(a)
}
```
