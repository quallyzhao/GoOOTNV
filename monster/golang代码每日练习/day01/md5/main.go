// md5
package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	h := md5.New()
	h.Write([]byte("123456")) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	// 常规打印 cipherStr ，是字节数组
	fmt.Println(cipherStr)
	// %x 以16进制打印字符串
	fmt.Printf("cipherStr以16进制打印字符串： \t%x\n",cipherStr)
	fmt.Printf("%s\n", hex.EncodeToString(cipherStr)) // 输出加密结果
}

