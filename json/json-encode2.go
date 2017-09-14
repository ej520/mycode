/*
go使用encoding/json进行JSON处理，常用到的内容进行如下整理：

项目	        详细	    使用方式
标准包	      encoding/json	        import “encoding/json”
Decoding	json数据-〉go数据结构	func Unmarshal(data []byte, v interface{}) error
Encoding	go数据结构-〉json数据	func Marshal(v interface{}) ([]byte, error)
*/
package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Id      int
	Name    string
	Country string
}

func main() {
	ej520 := person{1, "ej520", "China"}
	fmt.Println("EJ520=", ej520)

	result, err := json.Marshal(ej520)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("JSON result:", result)
		fmt.Println("JSON string(result):", string(result))
	}
}

/*
输出结果：
EJ520= {1 ej520 China}
JSON result: [123 34 73 100 34 58 49 44 34 78 97 109 101 34 58 34 101 106 53 50 48 34 44 34 67 111 117 110 116 114 121 34 58 34 67 104 105 110 97 34 125]
JSON string(result): {"Id":1,"Name":"ej520","Country":"China"}
*/
