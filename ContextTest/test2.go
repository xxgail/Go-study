package main

import (
	"context"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", myHander1)
	log.Fatal(http.ListenAndServe(":9998", nil))
}

func myHander1(rw http.ResponseWriter, r *http.Request) {
	userContext := context.WithValue(context.Background(), "user", "gai")
	ageContext := context.WithValue(userContext, "age", 18)
	rContext := r.WithContext(ageContext)
	// 可以用withContext函数，生成一个带有Context的*Request，这样存储有键值对的Context就跟着一起传递了

	doHander1(rw, rContext)
}

func doHander1(rw http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(string)
	age := r.Context().Value("age").(int)

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("the user is " + user + ",age is " + strconv.Itoa(age)))
}
