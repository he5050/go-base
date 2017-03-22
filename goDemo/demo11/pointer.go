package main

import (
	"fmt"
)

type name int8
type p struct {
	age int
	ok  bool
	name
}

func main() {
	a := 10
	/*指针的初识*/
	fmt.Printf("变量地址%x,\n变量指针%p,\n变量的值:%v\n", &a, &a, a)
	/* 声明实际变量 */
	b := 20
	/*声明int型指针变量*/
	var ip *int
	/* 指针变量的存储地址 */
	ip = &b
	fmt.Printf("b 变量的地址为:%x\n", &b)
	/*指针变量的存储地址*/
	fmt.Printf("ip 指针变量的地址为:%x\n", ip)
	/*该指针所对应的值是*/
	fmt.Printf("*ip 变量指向的值为:%d\n", *ip)
	/**复杂应用**/
	p1 := p{1, false, 2}
	/*声明P2是p1的指针*/
	p2 := &p1
	/*取值的差异*/
	fmt.Println(p1, p1.age, p1.ok, p1.name)
	fmt.Println(p2, p2.age, p2.ok, p2.name)
	fmt.Println(&p2, (*p2).age, (*p2).ok, (*p2).name)
	/*指针数组*/
	const MAX int = 5
	/*声明一个slice*/
	s1 := []int{10, 100, 300, 600, 400}
	/*声明一个指针数组*/
	var ptr [MAX]*int
	/*但是我们需要保存这个数组*/
	for i := 0; i < MAX; i++ {
		fmt.Printf("d[%d] = %d\n", i, s1[i])
		/*整数地址赋值给指针数组*/
		ptr[i] = &s1[i]
	}
	/*输出指针数组*/
	for i := 0; i < MAX; i++ {
		/*输出变量的值*/
		fmt.Printf("d[%d] = %d\n", i, *ptr[i])
		/*输出地址*/
		fmt.Printf("d[%d] = %d\n", i, ptr[i])
	}
	/*指针变量*/
	var k int
	var ptr1 *int
	var ptr2 **int

	k = 3000
	ptr1 = &k
	ptr2 = &ptr1
	/*获取上面的*/
	fmt.Printf("变量 k = %d\n", k)
	fmt.Printf("指针变量 *ptr1 = %d\n", *ptr1)
	fmt.Println("指针与指针变量的地址", ptr1, ptr2, *ptr2)
	fmt.Printf("指向指针的指针变量 **ptr2 = %d\n", **ptr2)
	/*参数*/
	x := 100
	y := 200
	fmt.Printf("交换前x=%d,y=%d\n", x, y)
	swap(&x, &y)
	fmt.Printf("交换后x=%d,y=%d\n", x, y)
}

/**
 * 实现交换
 * 传入x,y的指针
 */
func swap(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}
