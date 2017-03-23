## 结构(interface)

GO当中的接口是非常灵活的。
它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口。
接口定义了一组方法，如果某个对象实现了某个接口的所有方法，则此对象就实现了该接口。
接口的定义与使用
```
/*接口*/
type interface_name interface {
    method_name1 [return type]
    .
    .
    .
}
/*结构体 */
type struct_name struct{
    var_name var type
}
/*实现接口的方法*/
func (struct_name_var struct_name) method_name1() [return_type]{

}
```
简单事例
```
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
}
```
结果:

```
{iphone5}
正在进行连接
iphone5 已成功连接

正在读取
OK
正在写入
OK
iphone5 成功断开连接
成功断开连接 iphone5
```
1. 只要实现了接口的当中的方法,那么我们就可以说实现了该接口
2. GO所有接口都继承与空接口
3. 接口可以实现转换,上面的`Usb`可以换成`Connecter`
4. 接口是值类型的(也就是说会拷贝一份新),不会对原有对象进行修改
