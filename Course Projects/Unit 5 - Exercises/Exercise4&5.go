package main

import (
	"fmt"
)

type cType int

var d cType
var e int

func main() {
	fmt.Println(d)
	fmt.Printf("%T", d)
	d = 42
	fmt.Println(d)
	e = int(d)
	fmt.Println(e)
	fmt.Printf("%T", e)

}
