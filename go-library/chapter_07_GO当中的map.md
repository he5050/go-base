## map(集合)

`Map` 是一种无序的键值对的集合。`Map` 最重要的一点是通过 `key` 来快速检索数据，`key` 类似于索引，指向数据的值 `Map` 是一种集合，所以我们可以像迭代数组和切片那样迭代它。不过，`Map` 是无序的，我们无法决定它的返回顺序，这是因为 `Map` 是使用 `hash` 表来实现的，也是引用类型

使用`map`过程中需要注意的几点：
- `map`是无序的，每次打印出来的`map`都会不一样，它不能通过`index`获取，而必 须通过`key`获取
- `map`的长度是不固定的，也就是和`slice`一样，也是一种引用类型
- 内置的`len`函数同样适用于 `map`，返回`map`拥有的`key`的数量
- `map`的值可以很方便的修改，通过`numbers["one"]=11`可以很容易的把`key`为 `one`的字值改为`11`

1. map的创建
`var m1 map[string] string`
`m2 := make(map[int]string)`
>  上面是两种map 创建的常用形式,m是map的变量名`[string]`键的类型,string表示值的类型
>  如果不初始化 map，那么就会创建一个 nil map。nil map 不能用来存放键值对

2. map的简单使用

```
    package main

    import "fmt"
    import "sort"

    func main() {
    	var m1 map[int]string
    	m2 := make(map[int]string)
    	/*map的使用*/
    	m2[1] = "ok"
    	fmt.Println(m1)
    	fmt.Println(m2)
    	/*输出*/
    	fmt.Println()
    	/*删除*/
    	delete(m2, 1)
    	fmt.Println(m2)
    	/*查找与判断注意使用make创建map的时候,也是可以像生成slice一样分配一个空间大小的*/
    	/*有多少个map就得初始化多少次*/
    	m3 := make(map[int]map[int]string)
    	m3[1] = make(map[int]string)
    	m3[1][1] = "OKOK"
    	c := m3[1][1]
    	fmt.Println(c)
    	/*使用多返回值为判断是否传在key*/
    	b, ok := m3[2][1]
    	fmt.Println(b, ok)
    	/*ok的值只能是bool的*/
    	if !ok {
    		m3[2] = make(map[int]string)
    		m3[2][1] = "KOKO"
    	}
    	b, ok = m3[2][1]
    	fmt.Println(b, ok)
    	/*迭代*/
    	/*定义一个map类型的分配了5个*/
    	sm := make([]map[int]string, 5)
    	/*使用数组的迭代方法进行循环*/
    	/*使用这种方式是对sm对象拷贝之后进行操作*/
    	for _, v := range sm {
    		v = make(map[int]string, 1)
    		v[1] = "ok"
    		fmt.Println(v)
    	}
    	fmt.Println(sm)
    	fmt.Println("我们换一种方法进行迭代")
    	/*这里会直接进行操作*/
    	for i := range sm {
    		sm[i] = make(map[int]string, 1)
    		sm[i][1] = "KO"
    		fmt.Println(sm[i])
    	}
    	fmt.Println(sm)

    	/*迭代map当中元素*/
    	m7 := map[int]string{1: "a", 5: "b", 4: "c", 3: "d", 2: "e"}
    	/*slice*/
    	s1 := make([]int, len(m7))
    	/*循环中的i*/
    	i := 0
    	/*把map当中的所有key存到s1当中*/
    	for k := range m7 {
    		s1[i] = k
    		i++
    	}
    	/*对s1进行排序*/
    	sort.Ints(s1)
    	fmt.Println(s1)
    }

```
结果：

```
map[]
map[1:ok]
 
map[]
OKOK
 false
KOKO true
map[1:ok]
map[1:ok]
map[1:ok]
map[1:ok]
map[1:ok]
[map[] map[] map[] map[] map[]]
我们换一种方法进行迭代
map[1:KO]
map[1:KO]
map[1:KO]
map[1:KO]
map[1:KO]
[map[1:KO] map[1:KO] map[1:KO] map[1:KO] map[1:KO]]
[1 2 3 4 5]
```
这里我们使用`delete`来删除,使用`ok`来获取`key/value`是否存在
另外我们的`map`是无序的,如果要对`map`排序的话,我们一般使用的是`slice`,来排序之后存放对应的key-value
