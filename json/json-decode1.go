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
	var p person
	jdata := []byte(`{"id":1001,"name":"ej520","country":"China"}`)

	err := json.Unmarshal(jdata, &p)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(p)
	}

}

/*
输出结果：
{ json } master » go run json-decode1.go                                                  /d/develop/mygo/src/mycode/json
{1001 ej520 China}

*/

/*
注意点：
1、type struct中字段首字母需要大写， 否则，解码时无法存放进去。
2、json data中的字段是不区分大小写的。
*/
