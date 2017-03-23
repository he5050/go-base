## 反射 reflect

1. 所谓反射就是动态运行时的状态。我们一般用到的包是reflect包。
 - 反射使用`TypeOf`和`ValueOf`函数从接口中获取目标对象信息
 - 反射会把匿名字段作为独立字段
 - 可以使用反射修改对象状态
 - 能过反射可以`动态`调用方法

2. 具体实现
 1. 首先把要操作的内容转成reflect对象
 ```
 /*得到类型的元数据,通过t我们能获取类型定义里面的所有元素*/
 t := reflect.typeOf(i)
 /*得到实际的值,通过v我们获取存储在里面的值，还可以去改变值*/
 v := reflect.valueOf(i)
 ```
 2. 获取反射对象与相应的类型与值
  1. 普通类型
     ```
     x := 100
     v := reflect.valueOf(x)
     fmt.Printf("type:%s,kind:%s,value:%d",v.Type(),v.Kind(),v.Int())
     ```

   2. struct类型的
   ```
   type User struct{
        Id int
       Name string
   }

   u := User{1,"小强"}
   t := reflect.typeOf(u)
   v := reflect.valueOf(u)
   /*获取属性*/
   id := t.Elem().Filed(0).Id
   name := v.Elem().Filed(0).String()
   ```
3. 修改值(如果要进行修改必须传入地址,如果是struct则需要要使用Elem)
    ```
    x := 5.555
    p := reflect.valueOf(&x)
    p.Elem.SetFloat(7.77)
    ```
4. 简单案例

```
package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}
type Manager struct {
	User
	title string
}

func (u User) SayHello(name string) {
	fmt.Println(name, "Hello world by", u.Name)
}

/*最顶级的接口为空接口*/
/**
 * 打印输出结构当中的字段与方法信息
 */
func Info(i interface{}) {
	/*获取当前传过来结构信息*/
	t := reflect.TypeOf(i)
	/*获取结构名*/
	fmt.Println("Type", t.Name())
	/*判断是否是反射类型*/
	if k := t.Kind(); k != reflect.Struct {
		fmt.Println("不是指定类型")
		return
	}
	/*获取值 values*/
	v := reflect.ValueOf(i)
	fmt.Println("Fields:")
	/*循环输出,这里是按索引取值的*/
	for i := 0; i < t.NumField(); i++ {
		/*取得当前字段*/
		f := t.Field(i)
		/*取出字段的值*/
		val := v.Field(i).Interface()
		fmt.Printf("%6s,%v = %v\n", f.Name, f.Type, val)
	}
	/*获取方法*/
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%6s:%v \n", m.Name, m.Type)
	}
}

/*修改值*/
func Set(i interface{}) {
	v := reflect.ValueOf(i)
	/*判断是不是可以修改和是否是指针类型*/
	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("不能修改,或不是指针类型")
		return
	} else {
		v = v.Elem()
	}
	/*获取字段*/
	f := v.FieldByName("Name")
	/*判断字段是否存在*/
	if !f.IsValid() {
		fmt.Println("获取字段出错")
		return
	}
	if f.Kind() == reflect.String {
		f.SetString("我被修改了")
	}
}
func main() {
	u := User{1, "OK", 12}
	/*输出*/
	Info(&u)
	Info(u)
	fmt.Println("匿名字段")
	m := Manager{User: User{2, "err", 13}, title: "谁啊"}
	t := reflect.TypeOf(m)
	/*输出信息 %#v	值的Go语法表示*/
	fmt.Printf("%#v\n", t.Field(0))
	fmt.Printf("%#v\n", t.Field(1))
	/*取出匿名字段,传入的一个slice 第一个0表示字段,第二个0表示这个类型当中的第一个字段*/
	fmt.Printf("%v\n", t.FieldByIndex([]int{0, 0}))
	fmt.Printf("%v\n", t.FieldByIndex([]int{0, 1}))
	fmt.Printf("%v\n", t.FieldByIndex([]int{0, 2}))
	/*利用反射修改值*/
	x := 1111
	/*取得值,使用指针传递*/
	v := reflect.ValueOf(&x)
	v.Elem().SetInt(99)
	fmt.Println(x)
	/*使用方法进行修改,先判断类型是否可以修改,判断字段的类型是什么,然后在进行相应的修改*/
	u1 := User{22, "修改前", 50}
	fmt.Println(u1)
	Set(&u1)
	fmt.Println(u1)
	/*调用方法*/
	u2 := User{14, "OK", 15}
	u2.SayHello("111")
	/*获取方法*/
	v2 := reflect.ValueOf(u2)
	mv := v2.MethodByName("SayHello")
	args := []reflect.Value{reflect.ValueOf("我就是我")}
	mv.Call(args)
}

```
结果:
```
Type
不是指定类型

Type User
Fields:
    Id,int = 1
  Name,string = OK
   Age,int = 12
SayHello:func(main.User, string)
匿名字段
reflect.StructField{Name:"User", PkgPath:"", Type:(*reflect.rtype)(0x4d57c0), Tag:"", Offset:0x0, Index:[]int{0}
, Anonymous:true}
reflect.StructField{Name:"title", PkgPath:"main", Type:(*reflect.rtype)(0x4c50e0), Tag:"", Offset:0x20, Index:[]
int{1}, Anonymous:false}
{Id  int  0 [0] false}
{Name  string  8 [1] false}
{Age  int  24 [2] false}
99
{22 修改前 50}
{22 我被修改了 50}
111 Hello world by OK
我就是我 Hello world by OK
```
> 注意,我们要进行修改内容的时候,需要实现引用传递,同是要对类型进行判断。
> 在调用方法的时候,如果有参数的话需要传递:`[]reflect.Value{}`的slice
