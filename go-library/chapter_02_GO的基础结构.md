## 1. GO的一般结构

- Go的程序是通过`package`来组织的
- 只有一个`package`名称为`main`的包可以包含`main`函数
- 一个可执行的程序**有且仅有一个**`main`包

- 通过`import`关键字来导入其它非`main`包
- 通过`const`关键字来进行常量的定义
- 通过在函数体外部使用`var`关键字进行全局变量的声明与赋值
- 通过`type`关键字进行结构`struct`或接口`interface`的声明
- 能过`func`关键字来进行函数的声明

### 2. 包的别名与`.`的使用

- 我们在导入包的时候可以为我们使用包取一个别名或是使用`.`省略到包名,具体如下:

            import 'fmt'
            fmt.Println('我不用别名')
            /*------------------*/
            import ss 'fmt'
            ss.Println('我用了别名了')
            /*------------------*/
            import . 'fmt'
            Println('我省略了')

>  不建议使用省略调用,同时 省略与别名不能同时出现
- `_`操作,其实是引入该包，而不直接使用包里面的函数，而是调用了该包里面的init函数
```
            import (
              "database/sql"
              _ "github.com/ziutek/mymysql/godrv"
            )
```
- 在同一个包下面的文件属于同一个工程文件，不用`import`包，可以直接使用
- 在同一个包下面的所有文件的`package`名，都是一样的
- 在同一个包下面的文件`package`名都建议设为是该目录名，但也可以不是

### 3. 可见现规则
  在Go语言当中,有个叫可见性规则。使用`大小写`来决定该**常量、变量、类型、接口、结构、函数** 是否可以被外部包调用
  - 根据约定,`函数名`的首字母`小写`即为**private**,也就是我们学的别的语言当中类的私有变量
  - 根据约定,`函数名`的首字母`大写`即为**public**,也就是我们常说的公共

### 4. 变量的声明

> 变量与常量的命名规则: 以字母或下划线开头，由一个或多个字母、数字、下划线组成

- 变量的声音使用`var v1 int = 10`,解释一下`var`变量的声明的关键字,`v1`变量名,`int`变量的类型,`= 10`对变量进行赋值

> 这样写太复杂了是不,所以我们可以简写,因为go或自动判断变量的类型

            var v1 int =10 /**显示**/
            var v2 = 10  /**隐示**/
            v3 := 10

```
    `package` main
    import 'fmt'
    var a = 'hello'
    var b string = 'world'
    var c bool
    func main(){
        fmt.Println(a,b,c);
    }
    /**结果如下: hello world flase**/
```
- 批量声明如下:

        var (
            name1 type1
            name2 type2
            )

> 不能在函数体当中使用这种方法进行赋值   

- 变量的引用与赋值,以及多重赋值
```
  package main
  import "fmt"
  func main(){
    x: = 140
    fmt.Println(&x)
    x,y := 200 ,"abcd"
    fmt.Println(&x,x)
    fmt.Println(y)
}
/**
 * 0xc04200a2b0
 * 0xc04200a2b0 200
 * abcd
 */
```
> 能过上面我不难发现,在同一个作用域中，已存在同名的变量，则之后的声明初始化，则变成为赋值操作。但这个前提是，最少要有一个新的变量被定义，且在同一作用域，例如，上面的y就是新定义的变量,使用`&`表示引用的该变量的地址,可以能过输出结果看出

- 变量的简单使用
```
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
        fmt.Printf("面积为:%d\n",area)
        fmt.Printf("体积为:%d\n",volume)
        fmt.Printf("HEIGHT的类型是:%T,值为:%v\n",HEIGHT,HEIGHT)
        fmt.Println(a,b,c)
    }
    /**
     * 面积为:50
     * 体积为: 250
     * HEIGHT的类型是: int ,值为 5
     * 1 false string
     */
```
> 这里要补充一下:常量组中如不指定类型和初始化值，则与上一行非空常量右值相同,比如上面的HEIGHT就是,它的值就是为5


- 关于变量的的默认值所有变量的`值类型`默认为`0`,`bool`默认是`false`,`string`默认为一个空字符串(长度为0)

- **iota**
  iota，特殊常量，可以认为是一个可以被编译器修改的常量
  > 在每一个const关键字出现时，被重置为0，然后再下一个const出现之前，每出现一次iota，其所代表的数字会自动增加1
```
      package main

      import "fmt"

      func main() {
          const (
                  a = iota   //0
                  b          //1
                  c          //2
                  d = "h"   //独立值，iota += 1
                  e          //"ha"   iota += 1
                  f = 100    //iota +=1
                  g          //100  iota +=1
                  h = iota   //7,恢复计数
                  i          //8
          )
          fmt.Println(a,b,c,d,e,f,g,h,i)
      }
```
`0 1 2 h h 100 100 7 8`

### 5. `main`与`init`的区分

- 这两个函数在定义时不能有任何的参数和返回值
- 虽然一个`package`里面可以写任意多个`init`函数，但推荐只用**一个**
- Go程序会自动调用`init()`和`main()`
- 每个`package`中的`init`函数都是可选的，但`package main`就必须包含一个`main`函数
- 调用顺序是,先调用`init`函数，再调用`main`函数
- 运行程序，必须要运行存在`main`函数的`go`文件

> 程序的初始化和执行都起始于main包。如果main包还导入了其它的包，那么就会在编译时将它们依次导入。有时一个包会被多个包同时导入，那么它只会被导入一次（例如很多包可能都会用到fmt包，但它只会被导入一次，因为没有必要导入多次）。当一个包被导入时，如果该包还导入了其它的包，那么会先将其它包导入进来，然后再对这些包中的包级常量和变量进行初始化，接着执行init函数（如果有的话），依次类推。等所有被导入的包都加载完毕了，就会开始对main包中的包级常量和变量进行初始化，然后执行main包中的init函数（如果存在的话），最后执行main函数。

### GO指针的简单应用
```
 package main
 import "fmt"
 func main(){
    a := 1
    var p *int = &a
    fmt.Println(p)
    fmt.Println(*p)
 }
```
结果:
`0xc04200e240 1`
