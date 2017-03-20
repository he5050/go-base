package main
import "fmt"
func main(){
    const (
        LENGTH int = 10
        WIDTH int = 5
        HEIGHT
        )
    var area int
    var volume int
    const a,b,c = 1,false,"string"
    area = LENGTH*WIDTH
    volume = LENGTH*WIDTH*HEIGHT
    fmt.Printf("面积为:%d",area)
    fmt.Printf("体积为:%d",volume)
    fmt.Printf("HEIGHT的类型是:%T,值为:%v",HEIGHT,HEIGHT)
    fmt.Println(a,b,c)
}
