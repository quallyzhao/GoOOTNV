package main

import (
	"net/http"

)

func main (){
	http.HandleFunc("/index",handleIndex)
	http.ListenAndServe(":4000",nil)
}
func handleIndex(w http.ResponseWriter, r*http.Request){
	w.Write([]byte("hello this is version 1"))
}