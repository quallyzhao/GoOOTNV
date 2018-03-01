package main

import (
	"fmt"
)

func main() {
	// 1. 各种数组的初始化
	a := [20]int{19: 1}
	fmt.Println(a)
	b := [...]int{1, 2, 3, 4, 5}
	fmt.Println(b)
	c := [...]int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	fmt.Println(c)
	d := [...]int{29: 1}
	fmt.Println(d)
	var p *[30]int = &d
	fmt.Printf("p的值%p，d的地址%p----总结：p的值是&d",p,&d)
	x, y := 1, 2
	// 2. 数组存的是int型指针
	q := [...]*int{&x, &y}
	fmt.Println(q)
	// 3. 这是一个存储了2个数组指针的数组，存储的类型是*[5]int
	r := [...]*[5]int{&b, &c} // b和c数组都是相同类型的数组(长度一样)
	fmt.Println(r)
	// 4. 指向数组的指针
	aa := new([10]int)
	aa[1] = 2
	fmt.Println(aa,"aa")
	bb := [10]int{}
	bb[1] = 2
	fmt.Println(bb,"bb")

	aaa := [2][3]int{
		{1, 1, 1},
		{2, 2, 2}} // 注意,最后一个}不能单独占一行
	fmt.Println(aaa,"aaa")

	bbb := [2][3]int{
		{1: 1},
		{2: 2}} // 注意,最后一个}不能单独占一行
	fmt.Println(bbb,"bbb")

	ccc := [...][3]int{ // NOTE: 第二维不能用...省略
		{1: 1},
		{2: 2}} // 注意,最后一个}不能单独占一行
	fmt.Println(ccc)

	ppp := [...]int{5, 2, 6, 3, 9}
	fmt.Println(ppp)
	cnt := len(ppp)
	for i := 0; i < cnt; i++ {
		for j := i + 1; j < cnt; j++ {
			if ppp[i] < ppp[j] {
				temp := ppp[i]
				ppp[i] = ppp[j]
				ppp[j] = temp
			}
		}
	}
	fmt.Println(ppp)
}
