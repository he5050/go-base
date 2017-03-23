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
