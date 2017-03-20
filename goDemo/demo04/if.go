package main

import "fmt"

func main() {
	/*if 判断语句*/

	/*定义局部变量*/
	var a = 10
	/*使用衣服语句判断布尔表达式*/
	if a < 20 {
		/*如果为true的时候*/
		fmt.Println("a 小于20")

	} else {
		fmt.Println("a 大于20")

	}
	/*使用块当中的a变量,在这里声明的a只能是在当前区块内使用,不建议这么使用*/
	if a := 1; a < 3 {
		fmt.Println("输出当前if块当中的a变量")
		fmt.Println(a)
		a++
		fmt.Println(a)
	}

	fmt.Printf("a的值为:%d\n", a)

	/*循环语句*/
	var b = 1
	/*简单看下这个是一个死循环,但是呢,我们要 跳出来*/
	for {
		b++
		fmt.Println("这里无限循环当中的b:")
		fmt.Println(b)
		if b > 5 {
			fmt.Println("over")
			break
		}
	}

	var c = 1
	for c <= 3 {
		c++
		fmt.Printf("这里是带条件循环的C：%d \n", c)
	}

	var d = 1
	for i := 1; i <= 4; i++ {
		d++
		fmt.Printf("我是第%d次循环,当前D的值:%d \n", i, d)
	}
	/*switch语句*/
	/*简单使用*/
	h := 2
	switch h {
	case 0:
		fmt.Println("h=0")
	case 1:
		fmt.Println("h=1")
	case 2:
		fmt.Println("h=2")
	default:
		fmt.Println("None")
	}
	/*fallthrough的使用*/
	k := 2
	switch {
	case k >= 0:
		fmt.Println("k>=0")
		/*继续进行后面的判断*/
		fallthrough
	case k >= 2:
		fmt.Println("h>=2")
	case k >= 5:
		fmt.Println("h>=5")
	default:
		fmt.Println("None")
	}

	/*switch带判断的使用*/

	switch m := 2; {
	case m >= 0:
		fmt.Println("m>=0")
		/*继续进行后面的判断*/
		fallthrough
	case m >= 2:
		fmt.Println("m>=2")
	case m >= 5:
		fmt.Println("m>=5")
	default:
		fmt.Println("None")
	}
	/*break 加标签跳转*/

LEABLE1:
	for {
		fmt.Println("进无限循环了")
		for j := 0; j < 10; j++ {
			fmt.Printf("我要在这里输出J的值:%d\n", j)
			if j > 3 {
				break LEABLE1
			}
		}
	}
	fmt.Println("我跳出循环了哦!")

	/*补充使用for循环数组*/
	numbers := [10]int{1, 2, 3, 4, 5, 6}
	for i, x := range numbers {
		fmt.Printf("数组的第%d位的值为:%d\n", i, x)
	}
}
