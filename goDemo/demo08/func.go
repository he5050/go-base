package main

import (
	"fmt"
	"math"
)

/*定义了返回值的名称*/

func main() {
	/* 定义局部变量 */
	a := 100
	b := 200
	var ret int

	/* 调用函数并返回最大值 */
	ret = max(a, b)
	c, d := swap("XX", "YY")
	fmt.Println(c, d)
	fmt.Printf("最大值是 : %d\n", ret)
	/*匿名函数*/
	getSquare := func(x float64) float64 {
		return math.Sqrt(x)
	}
	/*调用函数*/
	fmt.Println(getSquare(13))
	x := 3
	/*这里结果为3*/
	fmt.Println("x = ", x)
	/*传地址,传入x的地址*/
	x1 := add1(&x)
	fmt.Println("x+1 = ", x1)
	/*我们在打印x的值*/
	fmt.Println("x =", x)
	/*闭包的使用*/
	num := getSe()
	/*我们连续调用几次看一下i的值*/
	fmt.Println(num())
	fmt.Println(num())
	fmt.Println(num())
	fmt.Println(num())
	/*重新创建一个*/
	num1 := getSe()
	fmt.Println(num1())
	fmt.Println(num1())
	/*函数做为值进行传递*/
	s1 := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("s1= ", s1)
	/*函数当做值来传递了*/
	odd1 := filter(s1, isOdd1)
	fmt.Println("odd1 是:", odd1)
	/*函数当做值来传递了*/
	odd2 := filter(s1, isOdd2)
	fmt.Println("odd2 是:", odd2)
}

/*声明一个函数类型*/
type textInt func(int) bool

/*定义两个判断是整除的方法*/
func isOdd1(x int) bool {
	if x%2 == 0 {
		return false
	}
	return true
}
func isOdd2(x int) bool {
	if x%2 != 0 {
		return false
	}
	return true
}

/*过滤函数*/
func filter(s []int, f textInt) []int {
	var result []int
	for _, value := range s {
		if f(value) {
			/*追加到结果slice当中*/
			result = append(result, value)
		}
	}
	return result
}

/*定义一个包函数,返回为一个函数*/
func getSe() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

/*简单的一个函数，实现了参数+1的操作*/
func add1(a *int) int {
	*a = *a + 1
	return *a
}

/*多返回值*/
func swap(x, y string) (string, string) {
	return y, x
}

/* 函数返回两个数的最大值 */
func max(num1, num2 int) int {
	/* 定义局部变量 */
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}
