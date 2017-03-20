## 条件与循环语句

1. **if判断语句**

```
    package main

    import "fmt"

    func main() {
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
    }
```
> 这里我们要注意的一点是当我们在*if*这个条件处声明这个前面已声明的变量的时候,此时声明的变量则只能在当前的**if**块中使用

***

2. **switch条件语句**
`switch` 语句用于基于不同条件执行不同动作，每一个 `case` 分支都是唯一的，从上直下逐一测试，直到匹配为。 `switch` 语句执行的过程从上至下，直到找到匹配项，匹配项后面也不需要再加`break`。
`Go`里面`switch`默认相当于每个`case`最后带有`break`，匹配成功后不会自动向下执行其他`case`，而是跳出整个`switch`, 但是可以使用`fallthrough`强制执行后面的`case`代码。
而如果`switch`没有表达式，它会匹配`true`

- `switch`简单使用

```

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
```

- `switch`当中使用`fallthrough`继续后面的条件判断

```
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
```
- `switch`初始化语句
```
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
    }
```

> 在这里我们要注意在`switch`初始化的时候,我们写了一个`m :=2`,这里的m的作用域只能是在当前`switch`当中


***

3. **for循环语句**

在`go`语言当中的只一个循环语语句

- 第一种模式(无条件)

```
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
```
- 第二种带简单条件的模式

```
    var c = 1
    for c <= 3 {
        c++
        fmt.Printf("这里是带条件循环的C：%d \n", c)
    }
```

- 第三种经典模式
```
    var d = 1
    for i := 1; i <= 4; i++ {
        d++
        fmt.Printf("我是第%d次循环,当前D的值:%d \n", i, d)
    }
```
- 补充一点关于循环数组的
```
    /*补充使用for循环数组*/
    numbers := [10]int{1, 2, 3, 4, 5, 6}
    for i, x := range numbers {
        fmt.Printf("数组的第%d位的值为:%d\n", i, x)
    }
```
结果：
```
数组的第0位的值为:1
数组的第1位的值为:2
数组的第2位的值为:3
数组的第3位的值为:4
数组的第4位的值为:5
数组的第5位的值为:6
数组的第6位的值为:0
数组的第7位的值为:0
数组的第8位的值为:0
数组的第9位的值为:0
```

4. GO当中的跳转语句

- `break LEABLE1`的使用

```
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
    }

```
在这里我们定义了一个标签,当运行到条件判断的时候,我使用`break+标签`来调到标签所在层级
如果我们把break换成goto,continue那么依然会是一个死循环,goto表示是跳转到别处去,continue表示跳过本次循环继续进行下面的循环
