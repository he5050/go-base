** 数组与slice
`Go` 语言提供了数组类型的数据结构。 数组是具有相同唯一类型的一组已编号且长度固定的数据项序列，这种类型可以是任意的原始类型例如整形、字符串或者自定义类型。
数组元素可以通过索引（位置）来读取（或者修改），索引从0开始，第一个元素索引为 0，第二个索引为 1，以此类推。

1. 数组的声明
```
    var arr [10]int
    var arr = [10]int{1,2}
```
2. 数组的使用
```
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
    }
```

结果：

```
声明数组了
[0 0 0 0 0 0 0 0 0 0]
[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1]
[10 20 30 40]
[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 22]
关于指针
&[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 22]
[0xc042008240 0xc042008248]
&[0 0 0 0 0 0 0 0 0 0]
[[0 0 9] [0 4 4]]
Array[0]=10
Array[1]=11
Array[2]=12
Array[3]=13
Array[4]=14
Array[5]=15
Array[6]=16
Array[7]=17
Array[8]=18
Array[9]=19
```
