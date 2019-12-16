## 跟PHP的比较
 - :robot: 神奇的网站[php=>golang](https://www.php2golang.com/)
 
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