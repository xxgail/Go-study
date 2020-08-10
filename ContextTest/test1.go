package main

import (
	"github.com/gorilla/context"
	"log"
	"net/http"
	"strconv"
)

func main() {
	//http.Handle("/",context.ClearHandler(http.HandlerFunc(myHander)))
	// 自动清除
	http.HandleFunc("/", myHander)
	log.Fatal(http.ListenAndServe(":9999", nil))
}

func myHander(rw http.ResponseWriter, r *http.Request) {
	context.Set(r, "user", "Gai")
	context.Set(r, "age", 22)
	context.Set(r, "sex", 1)

	context.Delete(r, "sex")
	doHander(rw, r)
}

func doHander(rw http.ResponseWriter, r *http.Request) {
	allParams := context.GetAll(r)
	log.Println(allParams)
	// 分析GetAll() 通过遍历返回map，因为map是一个引用类型，如果我们直接返回map，调用者就可能会对这个map进行修改，
	// 破坏了map的存储，所以必须要返回一个map的拷贝
	user := context.Get(r, "user").(string)
	age := context.Get(r, "age").(int)
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("the user is " + user + ",age is " + strconv.Itoa(age)))
}
