## 函数

`go`语言至少有一个main函数

```
    func funcName(input1 type1, input2 type2) (output1 type1, output2 type2) {
    //这里是处理逻辑代码
    //返回多个值
    return value1, value2
    }
```
- func：函数由 func 开始声明
- funcName：函数名称，函数名和参数列表一起构成了函数签名。
- input1 type1, input2 type2：参数列表，参数就像一个占位符，当函数被调用时，你可以将值传递给参数，这个值被称为实际参数。参数列表指定的是参数类型、顺序、及参数个数。参数是可选的，也就是说函数也可以不包含参数。
- output1 type1, output2 type2：返回类型，函数返回一列值。return_types 是该列值的数据类型。有些功能不需要返回值，这种情况下 return_types 不是必须的。
- 上面返回值声明了两个变量output1和output2，如果你不想声明也可以，直接就两个类型。
- 如果只有一个返回值且不声明返回值变量，那么你可以省略包括返回值的括号（即一个返回值可以不声明返回类型）
- 函数体：函数定义的代码集合。

```
1. 函数的简单使用
/*定义了返回值的名称*/

    func main() {
    	/* 定义局部变量 */
    	a := 100
    	b := 200
    	var ret int

    	/* 调用函数并返回最大值 */
    	ret = max(a, b)

    	fmt.Printf("最大值是 : %d\n", ret)
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
```

结果:`最大值是 : 200`

2. Go函数返回多个值

```
    /* 调用函数并返回最大值 */
    ret = max(a, b)
    c, d := swap("XX", "YY")
    fmt.Println(c, d)
    fmt.Printf("最大值是 : %d\n", ret)
    }

    /*多返回值*/
    func swap(x, y string) (string, string) {
    return y, x
    }
```

结果: `YY,XX`

3. Go函数支持变参。接受变参的函数是有着不定数量的参数的。为了做到这点，首先需要定义函数使其接受变参：
  `func myfunc(arg ...int) {}`
`arg ...int`告诉`Go`这个函数接受不定数量的参数。注意，这些参数的类型全部是`int`。在函数体中，变量`arg`是一 个`int`的`slice`,然后我们就可以像遍历`slice`一样把传过来参数都遍历出来:
```
for _, n := range arg {
fmt.Printf("And the number is: %d\n", n)
}
```

4. go语言函数的参数也是存在`值传递`和`引用传递`
 - 值传递
 ```
     /*匿名函数*/
     getSquare := func(x float64) float64 {
         return math.Sqrt(x)
     }
     /*调用函数*/
     fmt.Println(getSquare(13))
 ```
 结果: `3.605551275463989`

5. 引用传递

这就牵扯到了所谓的指针。我们知道，变量在内存中是存放于一定地址上的，修改变量实际是修改变量地址处的`内存`。有add1函数知道x变量所在的地址，才能修改x变量的值。所以我们需要将x所在地址&x传入函数，并将函数的参数的类型由int改为*int，即改为指针类型，才能在函数中修改x变量的值。此时参数仍然是按copy传递的，只是copy的是一个指针。

```
func main(){
    x := 3
    /*这里结果为3*/
    fmt.Println("x = ", x)
    /*传地址,传入x的地址*/
    x1 := add1(&x)
    fmt.Println("x+1 = ", x1)
    /*我们在打印x的值*/
    fmt.Println("x =", x)
}
    /*简单的一个函数，实现了参数+1的操作*/
    func add1(a *int) int {
    	*a = *a + 1
    	return *a
    }
```
结果：
```
x =  3
x+1 =  4
x = 4
```
> 在这里我们不难发现,x的值也相应的发生了改变,就因我们这里是传的地址

 - 1. 传指针使得多个函数能操作同一个对象。
 - 2. 传指针比较轻量级 (8bytes),只是传内存地址，我们可以用指针传递体积大的结构体。如果用参数值传递的话, 在每次copy上面就会花费相对较多的系统开销（内存和时间）。所以当你要传递大的结构体的时候，用指针是一个明智的选择。
 - 3. **Go语言中slice，map,闭包**的实现机制类似指针，所以可以直接传递，而不用取地址后传递 指针。（注：若函数需改变slice的长度，则仍需要取地址传递指针）

6. 闭包
`Go` 语言支持匿名函数，可作为闭包。匿名函数是一个"内联"语句或表达式。匿名函数的优越性在于可以直接使用函数内的变量，不必申明,学个`js`对这个闭包因该很熟悉。

```
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
    }

    /*定义一个包函数,返回为一个函数*/
    func getSe() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
    }
```
结果如下:
```
1
2
3
4
1
2
```
7. 把函数做为值进行传递

```

    func main(){
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
```
结果如下:

```
    /*函数做为值进行传递*/
    s1 := []int{1, 2, 3, 4, 5, 6, 7}
    fmt.Println("s1= ", s1)
    /*函数当做值来传递了*/
    odd1 := filter(s1, isOdd1)
    fmt.Println("odd1 是:", odd1)
    /*函数当做值来传递了*/
    odd2 := filter(s1, isOdd2)
    fmt.Println("odd2 是:", odd2)
```
能过上面的例子,我们可以看出使用函数传递的一般需要进行以下操作

 - 统一定义一个函数类型 参数类型、个数、顺序相同，返回值相同
 - 实现定义的函数类型
 - 把函数做来参数调用

 ***有没有发现,有点像接口的功能,实现接口,然后在实现接口所对应的方法***
 
> 函数当做值和类型在我们写一些通用接口的时候非常有用，通过上面例子我们看到testInt这个类型是一个函数类型，然后两个filter函数的参数和返回值与testInt类型是一样的，但是我们可以实现很多种的逻辑，这样使得我们的程序变得非常的灵活
