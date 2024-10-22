# 基础知识

~~本章内容部分来自于《Go语言核心编程》第一章~~
#### 1. 基础数据类型
1. 布尔型数据和整形数据不能进行互相转换。

2. 声明的布尔型变量如果不指定初始化值，则默认为false。

3. Go在语言层面禁止指针运算。

4. 字符串可以和切片相互转换
    ```
    str := "hello,world!"   // 声明一个字符串
    a := []byte(str)        // 将字符串转化为 []byte 类型切片
    b := []rune(str)        // 将字符串转化为 []rune 类型切片
    ```

5. Go内置得map不是并发安全的，并发安全的map可以使用标准包sync中的map。

6. 建议用 &Struct 的方式进行初始化，否则一旦struct增加字段，整个初始化语句会报错。

#### 2. 控制结构
1. Go没有三元条件运算符。

2. Go中的for：
    - `for i:= 0; i < 10; i++ {}` 普通的循环
    - `for condition {}` 类似于while循环
    - `for {}` 类似于C语言中的while(1)，即无限循环！
    
3. 跳转：
    - goto 直接跳到标签处（标签：Label）
    - break 可以跳出本次的for、switch、select语句循环或者跳出标签开始的循环
    - continue 用于跳出当前所在的for循环或者跳出标签开始的for循环


