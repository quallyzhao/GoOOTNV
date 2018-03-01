package main

import (
	"net/http"
	"time"
)
// 自己创建server mux && server
type myHandler struct{}
func (*myHandler)ServeHTTP(w http.ResponseWriter, r*http.Request){
	w.Write([]byte("hello this is version 3"+r.URL.String()))
}
func main(){
	server :=&http.Server{
		Addr:":4000",
		WriteTimeout: 2*time.Second,
	}
	mux := http.NewServeMux()
	mux.Handle("/",&myHandler{})

	server.Handler = mux
	server.ListenAndServe()
}
