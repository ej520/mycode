// from : http://blog.csdn.net/tiaotiaoyly/article/details/38942311
/*
Decode
将JSON数据解码
bool类型 转换为JSON的Boolean
整数，浮点数等数值类型 转换为JSON的Number
string 转换为JSON的字符串(带""引号)
struct 转换为JSON的Object，再根据各个成员的类型递归打包
数组或切片 转换为JSON的Array
[]byte 会先进行base64编码然后转换为JSON字符串
map 转换为JSON的Object，key必须是string
interface{} 按照内部的实际类型进行转换
nil 转为JSON的null
channel,func等类型 会返回UnsupportedTypeError

func Unmarshal(data []byte, v interface{}) error
类型转换规则和上面的规则类似
*/

package main

import (
	"encoding/json"
	"fmt"
)

type Animal struct {
	Name  string
	Order string
}

var json_Blob = []byte(`[{"Name":"an1","Order":"chifan"},{"Name":"an2","Order":"shuijiao"}]`)

func main() {
	var ans []Animal
	err := json.Unmarshal(json_Blob, &ans)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%+v", ans)
	}
}
