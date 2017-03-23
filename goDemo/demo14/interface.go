package main

import (
	"fmt"
)

/*定义接口*/
type USB interface {
	Name() string
	Read() string
	Write() string
	Connecter
}

/*我果我引入了别的类别,那么我们这个接口就是嵌入式*/
type Connecter interface {
	Connect()
}
type Phone struct {
	name string
}

/*实现相应的方法*/
func (p Phone) Connect() {
	fmt.Println("正在进行连接")
	fmt.Println(p.name, "已成功连接")
}
func (p Phone) Read() string {
	fmt.Println("正在读取")
	return "OK"
}
func (p Phone) Write() string {
	fmt.Println("正在写入")
	return "OK"
}
func (p Phone) Name() string {
	return p.name
}
func Disconnect1(usb USB) {
	/*类型Phone 是不是use类型判断*/
	if p, ok := usb.(Phone); ok {
		fmt.Println(p.name, "成功断开连接")
		return
	}
	fmt.Println("发生了未知的错误!")
}

/*使用空接口*/
func Disconnect2(usb interface{}) {
	switch v := usb.(type) {
	case Phone:
		fmt.Println("成功断开连接", v.name)
	default:
		fmt.Println("发生了未知的错误!")
	}
}

/*指针接口*/
type USB1 interface {
	P(*int)
}
type T int

func (t T) P(age *int) {
	fmt.Println("我带了指针了哦", age)
}
func main() {
	ip := Phone{"iphone5"}
	fmt.Println(ip)

	ip.Connect()

	a := ip.Read()
	fmt.Println(a)

	b := ip.Write()
	fmt.Println(b)

	Disconnect1(ip)
	Disconnect2(ip)
	/*指针调用*/
	var t T
	k := 15
	t.P(&k)
}
