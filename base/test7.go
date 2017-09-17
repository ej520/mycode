package main

import (
	"fmt"
)

func slice1() {
	// define array a ,length 10
	a := [10]int{}
	fmt.Println(a)

	//slice s1
	s1 := a[5]
	fmt.Println(s1)
	// slice s2  a index5-10  [5 6 7 8 9 ] no 10
	s2 := a[5:10]
	fmt.Println(s2)

	s3 := a[4:] //index >= 4
	s4 := a[:4] // index <4
	s5 := a[:]  //slice 全部copy
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)

}
func slice2() {
	//长度len，cap 容量
	fmt.Println("================slice2=================")
	s1 := make([]int, 10, 20)
	fmt.Println(len(s1), cap(s1))
	fmt.Println(s1)
}

func slice3() {
	fmt.Println("----------slice3--------reslice--------")
	//create array
	a := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g'}

	s1 := a[2:5]
	fmt.Println(s1)
	fmt.Println(string(s1))
	//reslice
	s2 := s1[:3]
	s3 := s1[3:] //当修改元数组同一元素时，s2，s3都会发生变化。除非s2或s3因为append超过cap而指向的内存位置发生变化
	//s4 := s1[3:15]
	//fmt.Println(s2, s3, s4)
	fmt.Println(s2, s3)
}

func slice4() {
	fmt.Println("----------slice4-----append-----------")
	s1 := make([]int, 3, 6)
	fmt.Printf("%p \n", s1)
	s1 = append(s1, 1, 2, 3)
	fmt.Printf("%v %p\n", s1, s1)
	s1 = append(s1, 1, 2, 3)
	fmt.Printf("%v %p\n", s1, s1) //地址发生变化，重新扩展了一倍容量
	fmt.Println(cap(s1))          //cap 6 => cap 12

}

func slice5() {
	fmt.Println("----------slice5-----copy-----------")
	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := []int{7, 8, 9, 10}
	fmt.Println(s1, s2)
	//copy(s1, s2) //copy s2的4各元素覆盖s1的前4个元素
	//copy(s2, s1) //copy s1 覆盖s2 ，超过容量的部分丢弃
	copy(s1[1:3], s2[2:4]) //将s2索引为2，3的元素 copy到s1的index为1，2的位置。
	fmt.Println(s1, s2)

}

func main() {
	slice1()
	slice2()
	slice3()
	slice4()
	slice5()
}
