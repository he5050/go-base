## defter

即延迟（defer）语句，你可以在函数中添加多个defer语句。当函数执行到最后时，这些defer语句会按照逆序执行，最后该函数返回。特别是当你在进行一些打开资源的操作时，遇到错误需要提前返回，在返回前你需要关闭相应的资源，不然很容易造成资源泄露等问题

1. 执行方式类似与java当中的析构函数,在函数体执行结束后,按调用顺序**相反顺序**逐一执行,也就相当于我们的`堆`,先进后出
2. 及时函数发生了**严重错误**也会逐一执行
3. 支持匿名函数调用
4. 常用于资源清理、文件关闭、解锁以及记录时间等操作
5. 能过与匿名函数配合可以return之后**修改**函数的计算结果
6. 如果函数体内某个变量做`defer`时匿名函数的参数,则在定义时已经获得了拷贝,否则是引用某个变量的地址

```
func ReadWrite() bool {
    file.Open("file")
    defer file.Close()
    if failureX {
          return false
    }
    if failureY {
          return false
    }
    return true
}
```

具体的实例如下:
```
    func main() {
    	fmt.Println("a")
    	for i := 1; i < 3; i++ {
            /*最后一个执行*/
    		defer fmt.Println(i)
    		defer func() {
                /*这里倒数第二执行,但是这里一个闭包引用的是i,*/
    			fmt.Println(i)
    		}()
    	}
        /*第二执行*/
    	defer fmt.Println("b")
        /*第一个执行*/
    	defer fmt.Println("c")
    }
```
结果：
```
a
c
b
3
2
3
1
```
> 为什么我们要使用defer

- 这里我们提一下,在`GO`当中是不骨异常机制的`try...catch...`是没有的,但是有`panic/recover`,这两种模式来处理
- `panic`可以是在任何地方引发,但是`recover`**只有**在`defter`调用的函数中有效

看下面的实例
```
func A() {
	fmt.Println("Func A")
}
func B() {
    /*如果,我们使用defer recover,那么程序就是在这里中断*/
	defer func() {
		//判断是否发生了panic,我们在panic之前注册这个恢复方法
		if err := recover(); err != nil {
			fmt.Println("Rrcover in B")
		}
	}()
	panic("Panic in B")
}

func C() {
	fmt.Println("Func C")
}
```
结果:
```
Func A
panic: Panic in B
 
goroutine 1 [running]:
main.B()
        D:/demo/go/demo/goDemo/demo09/defer.go:27 +0x6b
main.main()
        D:/demo/go/demo/goDemo/demo09/defer.go:19 +0x324
exit status 2
PS D:\demo\go\demo\goDemo\demo09> go run defer.go
Func A
Rrcover in B
Func C

```
