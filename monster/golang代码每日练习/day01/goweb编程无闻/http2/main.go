package main

import (
	"net/http"

)
// 自己创建server mux
type myHandler struct{}
func (*myHandler)ServeHTTP(w http.ResponseWriter, r*http.Request){
	w.Write([]byte("hello this is version 2"+r.URL.String()))
}
func main (){
	mux := http.NewServeMux()
	mux.Handle("/",&myHandler{})

	http.ListenAndServe(":4000",mux)
}
