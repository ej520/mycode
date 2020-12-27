package lession3

import (
	"fmt"
)

func FuncSlice() {

	a := [5]int{0, 1, 2, 3, 4}
	fmt.Println("array a:", a)
	//切片是数组的一部分
	//全部切片
	s1 := a[:]
	fmt.Println("a[:] : ", s1)
	//范围切片，前闭后开
	s2 := a[0:2]
	fmt.Println("a[0:2] : ", s2)
	//从index2开始 含2
	s3 := a[2:]
	fmt.Println("a[2:] : ", s3)
	//index3结束，不包含3
	s4 := a[:3]
	fmt.Println("a[:3] : ", s4)
}

func FuncSliceFuns() {
	// len cap append copy

	//array
	a := [5]int{0, 1, 2, 3, 4}

	//切片
	s1 := a[3:]
	//切片追加元素append
	s1 = append(s1, 6)
	s1 = append(s1, 6)
	s1 = append(s1, 6)

	//切片的长度len和容量cap, 切片的cap会自动扩容一倍
	fmt.Println(len(s1), cap(s1))
	//数组长度len和容量cap 保持不变
	fmt.Println(len(a), cap(a))

	//切片的拷贝copy
	s2 := a[:]
	s3 := a[2:4]
	//s2是[0 1 2 3 4] s3是[2 3]，copy会从index0开始用s3去覆盖s2，结果[2 3 2 3 4]
	copy(s2, s3)
	fmt.Println(s2)

	//切片的定义1,默认是空[]
	var ss []int
	//ss = append(ss, 12)
	//fmt.Println(ss)
	//切片的定义2,默认是[0,0,0,0]
	sss := make([]int, 4)
	fmt.Println(ss, sss) //[] [0 0 0 0]
}
