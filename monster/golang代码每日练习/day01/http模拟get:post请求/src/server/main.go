package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

//http 服务端，接受user/signup && use/login && images/<imageName> 请求
//默认收到请求随机等待0-1000毫秒响应
//返回的内容的话，为一个简单字符串
var loger *log.Logger
var logFile *os.File

func init() {
	File, err := os.OpenFile("./"+time.Now().Format("20060102")+".txt", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
	}
	logFile = File
	loger = log.New(logFile, "test_", log.Ldate|log.Ltime)
}
func main() {
	http.HandleFunc("/user/signup", handlerPostSignup)
	http.HandleFunc("/user/login", handlerPostLogin)
	http.HandleFunc("/images/", handlerGet)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("success")
	}
	time.Sleep(time.Duration(time.Second * 100000))
}
func handlerPostSignup(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())
	rnd := rand.Intn(1000)
	time.Sleep(time.Duration(time.Millisecond * time.Duration(rnd)))

	log := fmt.Sprintf("%s请求,Hello Signup---%s", r.Method, r.PostFormValue("name"))
	loger.Output(2, log)

	w.Write([]byte(log))
	fmt.Println(r.URL.String())

}
func handlerPostLogin(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())
	rnd := rand.Intn(1000)
	time.Sleep(time.Duration(time.Millisecond * time.Duration(rnd)))
	// 打印日志
	log := fmt.Sprintf("%s请求,Hello Signin---%s", r.Method, r.PostFormValue("name"))
	loger.Output(2, log)
	//写入响应
	w.Write([]byte(log))
	fmt.Println(r.URL.String())

}
func handlerGet(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())
	rnd := rand.Intn(1000)
	time.Sleep(time.Duration(time.Millisecond * time.Duration(rnd)))
	// 打印日志
	log := fmt.Sprintf("%s请求,Hello Image---%s", r.Method, r.URL.String())
	loger.Output(2, log)
	//写入响应
	w.Write([]byte(log))
	fmt.Println(r.URL.String())
}
