# Go-study 

## 1. 跟PHP的比较
 
 1. PHP中的数组，在Go中分为两种
    ```
    //php
        $arr1 = [1,2,3];
    //Go
        arr1 := [3]int{1,2,3}
        
    //php
        $arr2 = [
            'a' => 1,
            'b' => 2,
        ]
    //Go
        arr2 := map[string]int{"a":1,"b":2}       
       
    ```
 2. Go 语言中变量可以在三个地方声明：
    - 函数内定义的变量称为局部变量
    - 函数外定义的变量称为全局变量
    - 函数定义中的变量称为形式参数
    
    
    
## 2.十分讲究的Go（讲究

 1. 初始化数组时 `arr := [2]int{1,2}` 。
    数组中 {} 中的元素个数不能大于 [] 中的数字。
    
 2. :mute: 不想学指针所以先跳过。
 
 3. 关于切片的容量问题。nil 切片的长度和容量为 0 且没有底层数组。
    切片在添加（append）的时候，长度直接增加，但是容量只能翻倍。增加过程中一旦容量不够，直接翻倍。
    ```Go
    func main() {    
        var numbers []int
        printSlice(numbers)
        // len=0,cap=0 
    
        /* 允许追加空切片 */
        numbers = append(numbers, 0)
        printSlice(numbers)
        // len=1,cap=1
     
        /* 向切片添加一个元素 */
        numbers = append(numbers, 1)
        printSlice(numbers)
        // len=2,cap=2 (1*2)
     
        /* 同时添加多个元素 */
        numbers = append(numbers, 2,3,4)
        printSlice(numbers)
        // len=5,cap=6 (2*3)
    
        numbers = append(numbers, 5,6,7,8)
        // len=9,cap=12 (6*2)
    }
    ```
    
 4. 函数外的每个语句都必须以关键字开始（var, func 等等），因此 **:=** 结构不能在函数外使用。
 
 5. 常量不能用 := 语法声明。
    
 6. defer 语句会将函数推迟到外层函数返回之后执行。
    推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用。
    
 7. `type name string` 定义一个新类型，该类型的名称是name，其实就是string类型。
 
 8. 接口的乱七八糟的概念：
    - 接口也是值。它们可以像其它值一样传递。
    - 接口值可以用作函数的参数或返回值。
    - 在内部，接口值可以看做包含值和具体类型的元组：(value, type)
    - 接口值保存了一个具体底层类型的具体值。 
    - 接口值调用方法时会执行其底层类型的同名方法。
    
    - 指定了零个方法的接口值被称为**空接口**`interface{}`
        - 空接口可保存任何类型的值。（因为每个类型都至少实现了零个方法。）
        - 空接口被用来处理未知类型的值。
        
        
 ## 3.Go中比较好玩的一些命令
 1. go env 查看go的配置信息（GOPATH等信息
 2. go build 运行main主程序，查看输出结果
 3. go install 生成一个.exe文件，作为可操作的命令
 4. go test 测试写好的包，返回OK。