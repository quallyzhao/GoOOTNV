package main

import (
	"log"
	"time"
	"../worker"
	"sync"
	"unsafe"
)

var names =[]string{
	"ddfd",
	"dsafadf",
	"fadsffa",
	"fadsfd",
}

type namePrinter struct {
	name string
}
func (m *namePrinter)Task(){
	log.Println(m.name)
	time.Sleep(time.Second)
}

func main(){
	p:=worker.New(2)
	var wg sync.WaitGroup
	wg.Add(100*len(names))
	for i:=0;i<100;i++{
		for _,name := range names{
			np:= namePrinter{
				name: name,
			}
			go func(){
				p.Run(&np)
				wg.Done()
			}()
			}
	}
	wg.Wait()
	p.Shutdown()
}