package main

import "fmt"

func main() {
	/*声明一个长度为10的int型数组,默认的数组当中的所有元素的值为0*/
	var arr [10]int
	/*声明一个长度为20的int型数组,同时给最后一个元素设置其值为1*/
	arr1 := [20]int{19: 1}
	/*使有...来自动获取数组的长度*/
	arr2 := [...]int{0: 10, 1: 20, 2: 30, 3: 40}
	arr3 := [...]int{22: 22}
	/*指向数组的指针*/
	//var p *[23]int = &arr3
	p := &arr3

	/*使用 new返回数组指针*/
	p1 := new([10]int)
	/*指针数组*/
	x, y := 1, 2
	k := [...]*int{&x, &y}
	/*多维数组*/
	p2 := [2][3]int{
		{2: 9},
		{1: 4, 2: 4}}
	fmt.Println("声明数组了")
	fmt.Println(arr)
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println("关于指针")
	fmt.Println(p)
	fmt.Println(k)
	fmt.Println(p1)
	fmt.Println(p2)

	/*数组的传值与打印*/

	var n [10]int
	var i, j int

	/*为数组初始化值*/
	for i = 0; i < 10; i++ {
		n[i] = i + 10
	}
	/*输出数组*/
	for j = 0; j < 10; j++ {
		fmt.Printf("Array[%d]=%d\n", j, n[j])
	}
	/*冒泡法*/
	arr4 := [...]int{5, 2, 9, 1, 14, 3, 5, 9, 15, 26, 6}
	fmt.Println(arr4)
	/*获取长度*/
	num := len(arr4)
	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			if arr4[i] < arr4[j] {
				temp := arr4[i]
				arr4[i] = arr4[j]
				arr4[j] = temp
			}
		}
	}
	fmt.Println(arr4)
}
