// from :http://blog.csdn.net/tiaotiaoyly/article/details/38942311

/*
结构体
结构体必须是大写字母开头的成员才会被JSON处理到，小写字母开头的成员不会有影响。

Mashal时，结构体的成员变量名将会直接作为JSON Object的key打包成JSON；Unmashal时，会自动匹配对应的变量名进行赋值，大小写不敏感。
Unmarshal时，如果JSON中有多余的字段，会被直接抛弃掉；如果JSON缺少某个字段，则直接忽略不对结构体中变量赋值，不会报错。
*/

package main

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Name  string
	Body  string
	Time  int64
	inner string
}

var m = Message{
	Name:  "ej520",
	Body:  "hello",
	Time:  1294706395881547000,
	inner: "ok",
}

func main() {
	b := []byte(`{"nAmE":"Bob","Food":"Pickle", "inner":"changed"}`) //反单引号``
	err := json.Unmarshal(b, &m)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(m)
	}
}

/*
Output:
{Bob hello 1294706395881547000 ok}

json只允许双引号
struct：inner: 小写了，所以没有被替换，值还是ok
json串中的字段不区分大小写，nAmE和Name进行配对，结果为Bob
json Food字段无匹配，丢弃
struct：Time 没有配对，值未发生变化。
*/
