package main

import (
	"io/ioutil"
	"fmt"
	"os"
)

func catFile(name string){
	buf ,err := ioutil.ReadFile(name)
	if err !=nil{
		fmt.Println(err)
		return
	}
	fmt.Println(string(buf))
}
func main(){
	catFile(os.Args[1])
}
