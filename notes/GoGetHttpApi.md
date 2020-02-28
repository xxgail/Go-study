# go中处理各种请求方式以及处理接口请求参数

1.处理请求方式
```go
package main

import (
"fmt"
"io/ioutil"
"net/http"
"net/url"
"strings"
)

// Get 请求
func httpGet()  {
    resp,err := http.Get("https://aaa.com?id=1")
    if err != nil{
    }
    
    defer resp.Body.Close() // 一定要关闭返回的response中的body
    body,err := ioutil.ReadAll(resp.Body) // 读取body中的信息
    if err != nil { 
    }

    fmt.Println(string(body))
}

// Post 请求
func httpPost()  {
	resp,err := http.Post("https://aaa.com",
        "application/x-www-form-urlencoded",
        strings.NewReader("id=1"))
    // 这里的第二个参数是传入参数的类型
    // 第三个参数固定类型为io.Reader类型，因此调用了strings包中的func NewReader(s string) *Reader 转化为io.Reader类型
 
    if err != nil {
    }
    
    defer resp.Body.Close()
    body,err := ioutil.ReadAll(resp.Body)
    if err != nil {
    }
    
    fmt.Println(string(body))
}

// Post Form 表单请求
func httpPostForm()  {
    resp,err := http.PostForm("https://aaa.com",
        url.Values{"key":{"value"},"id":{"123"}})
    // 第二个参数规定的类型是URL包中的Values type Values map[string][]string
    if err != nil {
    }
    
    defer resp.Body.Close()
    body,err := ioutil.ReadAll(resp.Body)
    if err != nil {
    }
    
    fmt.Println(string(body))
}

// 这里一般是处理复杂的请求，比如要设置请求头以及一些请求信息时调用
func httpDo(){
    client := &http.Client{} // 实例化client结构体
    
    // func NewRequest(method, urlStr string, body io.Reader) (*Request, error) 
    // 第一个是请求方法，第二个是请求地址，第三个是请求的参数，这里依旧调用了string保重的对应方法转化为对应得数据类型
    req,err := http.NewRequest("POST",
        "https://aaa.com",
        strings.NewReader("id=1"))
    if err != nil {
    }

    // 设置请求头信息
    req.Header.Set("Content-Type","application/x-www-form-urlencoded")
    req.Header.Set("Cookie","name=an")

    resp,err := client.Do(req)

    if err != nil {
    }
        
    defer resp.Body.Close()
    body,err := ioutil.ReadAll(resp.Body)
    if err != nil {
    }
        
    fmt.Println(string(body))
}

func main()  {
    httpGet()
    httpPost()
}
```