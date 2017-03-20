package main

import "fmt"

func main() {
	/*定义一个简单切片*/
	var s1 []int
	/*截取数组成为切片,使有起始与结束*/
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s2 := a[9]
	s3 := a[4:6]
	s4 := a[3:]
	s5 := a[:6]
	s6 := a[:]
	/*make方法创建 第一个参数是 类型,第二个当前存放的元素个数,第三个分配的空间大小*/
	s7 := make([]int, 3, 10)

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)
	fmt.Println(s6)
	fmt.Println(s7)
	fmt.Printf("%v,%p\n", s7, s7)
	/*追加*/
	s7 = append(s7, 1, 2, 3, 4, 5, 6)
	fmt.Printf("%v,%p\n", s7, s7)
	/*追加,当超过了设置之好了的空间大小之后,会重新分配一个空间*/
	s7 = append(s7, 1, 2, 3, 4, 5, 6)
	fmt.Printf("%v,%p\n", s7, s7)
	/*slice是指是向地址空间的*/
	b := []int{1, 2, 3, 4, 5}
	s8 := b[1:5]
	s9 := b[1:3]
	fmt.Println(s8, s9)
	s8[0] = 111
	fmt.Println(s8, s9)
	s9 = append(s9, 5, 5, 5, 5, 5, 5, 5, 5, 5)
	s8[1] = 222
	fmt.Println(s8, s9)
	/*copy, 把后面的参数复制前一个参数上*/
	s11 := []int{1, 2, 3, 4, 5, 6, 7}
	s12 := []int{11, 21, 31}
	copy(s11, s12)
	fmt.Println(s11)
	s13 := []int{1, 2, 3, 4}
	s14 := []int{11, 12, 14, 15, 15, 19}
	copy(s13, s14)
	fmt.Println(s13)
}
