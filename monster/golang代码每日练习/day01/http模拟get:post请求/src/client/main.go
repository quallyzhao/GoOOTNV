package main

// 客户端随机模拟10次请求
// 请求有2种
// 1. get请求      /images/{{date}}/{{md5-hash-of-timestamp}}.jpg    {{date}}替换为当前日期，比如2018-02-22 ，后面那个为时间戳的md5 hash
// 2. post请求    requestBody附带用户名，response返回"hello word---【用户名】"
// flage 参数可以指定2个 -i  -c    i是请求的主机ip c是模拟请求次数
import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var ip string
var SimuCount int

// 默认请求localhost:8000 请求10次
func main() {
	flag.StringVar(&ip, "i", "http://115.231.182.109:8000", "testIp")
	flag.IntVar(&SimuCount, "c", 1000, "testCount")
	flag.Parse()

	fmt.Println("开始模拟请求")
	time.Sleep(time.Duration(time.Second))
	// 获取日期字符串
	dataStr := DataStr()
	//获取md5 时间戳
	md5TimeStr := Md5TimeStr()
	//拼接url
	url := fmt.Sprintf("%s/images/%s/%s.jpg", ip, dataStr, md5TimeStr)
	//模拟10次请求
	rand.Seed(time.Now().Unix())
	for i := 0; i < SimuCount; i++ {
		if rand.Intn(1000)%2 == 0 {
			//模拟get请求
			SimuGetRequest(url)

		} else {
			//模拟post请求
			postUrl := ip + "/user/signup"
			postUrl2 := ip + "/user/login"
			SimuPostRequest(postUrl, "zhouchengbin")
			SimuPostRequest(postUrl2, "zhouchengbin")
		}
	}
}
func SimuPostRequest(url string, name string) {
	postStr := fmt.Sprintf("name=%s", name)
	resPost, err := http.Post(url,
		"application/x-www-form-urlencoded",
		strings.NewReader(postStr))
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resPost.Body)
	defer resPost.Body.Close()
	if err != nil {
		// handle error
	}
	fmt.Println(string(body))
}
func SimuGetRequest(url string) {
	//发起请求
	resGet, err := http.Get(url)
	defer resGet.Body.Close()
	if err != nil {
		//handle error
		fmt.Println(err)
	}
	bodyGet, err := ioutil.ReadAll(resGet.Body)
	fmt.Println(string(bodyGet))
}

//获取日期字符串
func DataStr() string {
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	// 获取日期字符串
	timestr := tm.Format("2006-01-02")
	return timestr
}

//获取md5加密过的timestamp
func Md5TimeStr() string {
	var t int64 = time.Now().Unix()
	var s string = time.Unix(t, 0).Format("2006-01-02 15:04:05")
	h := md5.New()
	h.Write([]byte(s))
	cipherStr := h.Sum(nil)
	md5TimeStr := hex.EncodeToString(cipherStr)
	return md5TimeStr

}
