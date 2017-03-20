package main

import "fmt"

func main() {
	var m1 map[int]string
	m2 := make(map[int]string)
    /*map的使用*/
    m2[1] = 'ok'
	fmt.Println(m1)
	fmt.Println(m2)
    /*输出*/
    fmt.Println()
}
