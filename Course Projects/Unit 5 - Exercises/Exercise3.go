package main

import (
	"fmt"
)

var x = 42
var y = "James Bond"
var z = true

func main() {
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z) //%v displays just the value
	//Prints the values of x, y and z into one single string.
	fmt.Println(s) //Prints that string
}
