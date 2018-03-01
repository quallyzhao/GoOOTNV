package main

import (
	"fmt"
)

func main() {

	//1. 使用make创建切片
	s := make([]int, 3, 10)
	fmt.Println("s",s, "\ts的长度:",len(s), "\ts的容量:","s",cap(s))

	//2.切片再切片, 下标可超过切片长度, 但不能超过底层数组容量
	a := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k'}
	sa := a[2:5]
	sb := a[3:5]
	fmt.Println(string(sa))
	fmt.Println(string(sb))
	fmt.Println(string(sa[3:5]))
	fmt.Println(string(sa[3:9]))
	// fmt.Println(string(sa[3:10])) // error 超出原始数组范围
	//3. 切片的内存分配
	ss := make([]int, 3, 6)
	fmt.Printf("%v %p---\n", ss,ss)
	ss = append(ss, 1, 2, 3)
	fmt.Printf("%v %p\n", ss, ss)
	ss = append(ss, 1, 2, 3) // 超过容量最大值,重新分配了内存，可以看到打印出来的地址不同
	fmt.Printf("%v %p\n", ss, ss)

	//4. 切片操作的影响, 多个切片内存重叠部分都会改变, 注意是内存重叠部分, 不是下标重叠部分
	aaa := []int{1, 2, 3, 4, 5, 6}
	sss1 := aaa[1:3]
	sss2 := aaa[2:5]
	fmt.Println(sss1, sss2)
	fmt.Printf("%p--- %p---\n", sss1, sss2)
	sss2 = append(sss2, 1) // 在改变sss1之前对sss2增加元素, 但是没有超出容量, 不重新分配内存
	fmt.Printf("%p--- %p---\n", sss1, sss2)
	// sss2 = append(sss2, 1,1) // 在改变sss1之前对sss2增加元素, 且超过原容器容量,会重新分配内存
	sss1[1] = 9 // 可能改变共同部分(sss2没有重新分配的话), 也可能不改变共同部分(如果sss2重新分配了就没有共同部分了)
	fmt.Println(sss1, sss2)

	//5.  切片拷贝,以2个切片中短的切片为准, 从左到右尽可能拷贝最多个元素, 直到其中一个切片结束,丢弃多余的部分
	x := []int{1, 2, 3, 4, 5, 6}
	y := []int{7, 8, 9, 10}
	// copy(y, x)
	copy(y, x[1:])
	fmt.Println(y,"将x拷贝到y")

	// z := x[0:len(x)]
	// z := x[0:]
	z := x[:]
	fmt.Println(z)

}
