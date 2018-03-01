package main

import (
	"os"
	"log"
	"fmt"
)

func main(){

	//openFile()
	readFile()
}

func createFile(){
	// os.create  存在清空，不存在创建
	f,err := os.Create("a.txt")
	if err !=nil{
		log.Fatal(err)
	}
	f.WriteString("hello\n")
	f.Close()
}
func openFile(){
	f,err := os.OpenFile("a.txt",os.O_APPEND|os.O_CREATE|os.O_RDWR,0644)
	if err !=nil{
		log.Fatal(err)
	}
	f.WriteString("hello\n")
	f.Close()
}

func readFile(){
	f,err := os.OpenFile("a.txt",os.O_APPEND|os.O_CREATE|os.O_RDWR,0644)
	if err !=nil{
		log.Fatal(err)
	}
	buf := make([]byte,2014)
	f.Read(buf)
	fmt.Println(string(buf))
	f.Close()
}
