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
	/*交换map的键与值*/
	mis1 := map[int]string{1: "a", 5: "b", 4: "c", 3: "d", 2: "e"}
	fmt.Println(mis1)
	mis2 := make(map[string]int)
	for k, v := range mis1 {
		mis2[v] = k
	}
	fmt.Println(mis2)
}
