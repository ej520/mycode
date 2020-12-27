package lession3

import "fmt"

func FuncMap() {
	//map: key value

	//map定义1
	var m1 map[string]string
	m1 = map[string]string{}

	m1["name"] = "ej520"
	m1["sex"] = "男"
	fmt.Println("FuncMap m1:", m1)

	//map定义2
	m2 := map[string]string{}
	m2["name"] = "wyj"
	m2["age"] = "4"
	fmt.Println("FuncMap m2:", m2)

	//map定义3
	m3 := make(map[string]string)
	m3["name"] = "ej520"
	m3["sex"] = "男"
	fmt.Println("FuncMap m3:", m3)

	//map常用   interface{} 和 struct
	//m := map[int]bool{}
	// m[1] = true
	// m[2] = false
	m := map[string]interface{}{}
	m["a"] = 1
	m["b"] = false
	m["c"] = "str"
	m["d"] = []int{1, 2, 3, 4}

	fmt.Println(m)
}

func FuncMapDelete() {
	m := map[int]interface{}{}
	m[1] = true
	m[2] = false
	m[3] = 3
	m[4] = "ej520"
	m[5] = []int{1, 2, 3}
	m[6] = []string{"a", "b"}
	fmt.Println("FuncMapDelete:Before", m)
	delete(m, 5)
	fmt.Println("FuncMapDelete:5 After", m)
	delete(m, 6)
	fmt.Println("FuncMapDelete:6 After", m)
}

func FuncMapDeleteRange() {

	m := map[int]interface{}{}
	m[1] = true
	m[2] = false
	m[3] = 3
	m[4] = "ej520"
	m[5] = []int{1, 2, 3}
	m[6] = []string{"a", "b"}

	for k, v := range m {
		fmt.Println("FuncMapDeleteRange:", k, v)
	}

	//map是无序的
	for k := range m {
		//fmt.Println("FuncMapDeleteRange:Before", m)
		delete(m, k)
		fmt.Println("FuncMapDeleteRange:After", m)
	}

}
