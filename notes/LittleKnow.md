## 记录学习过程中的小知识点

1. 初始化数组时 `arr := [2]int{1,2}` 。
    数组中 {} 中的元素个数不能大于 [] 中的数字。
    
2. :mute: 不想学指针所以先跳过。
    - 使用指针接收者的原因有二：
        - 首先，方法能够修改其接收者指向的值。
        - 其次，这样可以避免在每次调用方法时复制该值。若值的类型为大型结构体时，这样做会更加高效。
 
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
    可以保证退出函数时释放资源。
    
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
 
9. 方法的乱七八糟的概念：
    - 方法就是一类带特殊的 接收者 参数的函数。
    - 方法接收者在它自己的参数列表内，位于 func 关键字和方法名之间。
    - 方法只是个带接收者参数的函数。  
    - 只能为在同一包内定义的类型的接收者声明方法，而不能为其它包内定义的类型（包括 int 之类的内建类型）的接收者声明方法。就是接收者的类型定义和方法声明必须在同一包内；不能为内建类型声明方法。
 
10. 常用几个类型转换
    ```go
    // string到int
    int, err := strconv.Atoi(string)
    // string到int64
    int64, err := strconv.ParseInt(string, 10, 64)
    // int到string
    string := strconv.Itoa(int)
    // int64到string
    string := strconv.FormatInt(int64,10)
    // map转json
    m :=map[string]string{"type":"10","msg":"hello."}
    mjson,_ :=json.Marshal(m)
    mString :=string(mjson)
    fmt.Printf("print mString:%s",mString)
    // json转map
    mmap := make(map[string]interface{})
    json := 
    json.Unmarshal([]byte(),&auction)
    ```
    
11. 查看变量类型 `reflect.TypeOf(v)`。

12. **TIME:** 2006-01-02 15:04:05

13. 用 map[string]interface{} 可以存储值的类型不一样的映射数组。但是读取出来的都是接口类型，需要做进一步的转换。
    - v.(string)
    
14. 变量名采用驼峰法，不要有下划线，不要全部大写。

15. context 注意: 使用时遵循context规则
    - 不要将 Context放入结构体，Context应该作为第一个参数传入，命名为ctx。
    - 即使函数允许，也不要传入nil的 Context。如果不知道用哪种Context，可以使用context.TODO()。
    - 使用context的Value相关方法,只应该用于在程序和接口中传递和请求相关数据，不能用它来传递一些可选的参数。
    - 相同的 Context 可以传递给在不同的goroutine，Context 是并发安全的。

 
    

       