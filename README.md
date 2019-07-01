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
 
 3. 关于切片的容量问题。
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