# 类型系统

~~本章内容部分来自于《Go语言核心编程》第三章 类型系统~~
## 1. 类型简介
#### 1-1. 命名类型和未命名类型
1. 命名类型：类型可以通过标识符来表示。
    - 预声明类型
        - 布尔型   bool
        - 整型    int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,uintptr
        - 浮点型   float32,float64
        - 复数    complex64,complex128
        - 字符    byte(int8别名),rune(int32别名)
        - 字符串   string
        - 错误    error
    - 自定义类型 type new_type old_type
2. 未命名类型：一个类型由预声明类型、关键字和操作符组合而成。也称为类型字面量。
    - 类型字面量（复合类型）
        - 数组    array
        - 切片    slice
        - 字典    map
        - 指针    pointer
        - 通道    channel
        - 结构    struct
        - 接口    interface{}
        - 函数    function
        
#### 1-2. 底层类型
1. 预声明类型和类型字面量的底层类型是它们自身。
2. 自定义类型的底层类型是逐层递归向下查找的，知道查到的old_type是预声明类型或者类型字面量为止。

#### 1-3. 类型赋值
1. 类型赋值的条件
- 参考[test-1-3.go](../TypeSys/test-1-3.go)
- `var b T2 = a` 条件：
    - T1 和 T2 的类型相同
    - T1 和 T2 具有相同的底层类型，并且 T1 和 T2 里面至少有一个是未命名类型。
    - T2 是接口类型，T1 是具体类型，T1 的方法集是 T2 方法集的超集。
    - T1 和 T2 都是通道类型，它们拥有相同的元素类型，并且 T1 和 T2 中至少有一个是未命名类型。
    - a 是预声明标识符 nil，T2 是pointer、function、slice、map、channel、interface类型中的一个
    - a 是一个字面常量值，可以用来表示类型 T 的值。
    
#### 1-4. 类型强制转换
