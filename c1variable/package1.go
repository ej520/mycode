package lession1

import "fmt"

var A string = "lession1 public A"
var b string = "lession1 private b"

func funca() {
	fmt.Println("priv func a", b)
}

func Funcb() {
	fmt.Println("public func b", b)
}
