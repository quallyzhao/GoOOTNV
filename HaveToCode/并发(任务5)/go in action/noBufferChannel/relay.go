// 这个例子展示如何用无缓冲通道来模拟接力赛
// 4个goroutine间的接力比赛
package main

import "sync"
// 用来等待程序结束
var wg sync.WaitGroup
