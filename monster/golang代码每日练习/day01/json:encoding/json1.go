package main 
import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"io"
)
type Message struct {
    Name string
    Body string
    Time int64
}

func main(){
	//1. marshal
	m := Message{"Alice", "Hello", 1294706395881547000}
	b, err := json.Marshal(m)
	if err !=nil{

	}
	fmt.Println(string(b))
	//b == []byte(`{"Name":"Alice","Body":"Hello","Time":1294706395881547000}`)

	c := []byte(`{"Name":"Bob","Food":"Pickle"}`)
	var n Message
	err2 := json.Unmarshal(c, &n)
	if err2!=nil{
		log.Fatal(err2)
	}
	fmt.Printf("%#v\n",n)
	//n == main.Message{Name:"Bob", Body:"", Time:0}


	cc := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f interface{}
	err3 := json.Unmarshal(cc, &f)
	if err3!=nil{
		log.Fatal(err3)
	}
	fmt.Printf("%#v",f)
//	通过类型断言来访问
//	m := f.(map[string]interface{})
//	for k, v := range m {
//		switch vv := v.(type) {
//		case string:
//			fmt.Println(k, "is string", vv)
//		case int:
//			fmt.Println(k, "is int", vv)
//		case []interface{}:
//			fmt.Println(k, "is an array:")
//			for i, u := range vv {
//				fmt.Println(i, u)
//			}
//		default:
//			fmt.Println(k, "is of a type I don't know how to handle")
//		}
//	}


	//json 包提供了Decoder 和 Encoder 用来支持JSON 数据流的读写。函数NewDecoder 和 NewEncoder 封装了io.Reader和io.Writer 接口类型。
	//
	//func NewDecoder(r io.Reader) *Decoder
	//func NewEncoder(w io.Writer) *Encoder
	//下面是一个从标准输入读入连续的JSON object的实例程序，每个结构体只留下Name域，然后把objects写到标准输出：
	dec := json.NewDecoder(os.Stdin)
	enc := json.NewEncoder(os.Stdout)
	for {
		var v map[string]interface{}
		if err := dec.Decode(&v); err != nil {
			log.Println(err)
			return
		}
		for k := range v {
			if k != "Name" {
				delete(v, k)
			}
		}
		if err := enc.Encode(&v); err != nil {
			log.Println(err)
		}
	}
}
