## Range(范围) 补充说明
`GO`语言当中`range`关键字用于`for`循环迭代数组(array),切片(slice),集合(map),通到(channel)的元素。
在数组和切片中返回元素的**索引(key)**,在集合当返回**键值对(key-value)的key值**

#### 简单使用

```
/*实现对一与slice求和*/
nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
arr := [11]int{1, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
sum := 0
maps := map[string]string{"a": "app", "b": "bpp", "c": "cpp"}
/*这里们不需要索引,所以使用_省略*/
for _, num := range nums {
    sum += num
}
fmt.Printf("求和的结果为:%d\n", sum)
/*数组的使用与索引*/
for i, num := range arr {
    if num == 30 {
        fmt.Printf("当前的索引为:%d\n", i)
    }
}
/*作用与map*/
for k, v := range maps {
    fmt.Printf("%s ---> %s\n", k, v)
}
```
结果:

```
求和的结果为:55
当前的索引为:3
a ---> app
b ---> bpp
c ---> cpp
```
