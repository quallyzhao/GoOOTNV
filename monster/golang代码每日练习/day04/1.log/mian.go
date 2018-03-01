package main

import (
	"os"
	"time"
	"fmt"
	"log"
)

func main() {
	logFile ,err := os.OpenFile("./"+time.Now().Format("20060102")+".txt",os.O_APPEND|os.O_RDWR|os.O_CREATE,0644)
	if err != nil{
		fmt.Println(err)
	}
	loger := log.New(logFile,"test_",log.Ldate|log.Ltime)
	loger.Output(2, "打印一条日志信息");

}
