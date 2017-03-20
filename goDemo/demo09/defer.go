package main

import (
	"fmt"
)

func main() {
	fmt.Println("a")
	for i := 1; i < 3; i++ {
		defer fmt.Println(i)
		defer func() {
			/*这里倒数第二执行,但是这里一个闭包引用的是i,*/
			fmt.Println(i)
		}()
	}
	defer fmt.Println("b")
	defer fmt.Println("c")
	A()
	B()
	C()
}

func A() {
	fmt.Println("Func A")
}
func B() {

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
