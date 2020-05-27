package main

import (
	"fmt"
)

func main() {
	n, _ := fmt.Println("Hello, playground", 42, true)
	fmt.Println(n)

	x := 42 //Declares and assigns the value 42 to the variable X
	fmt.Println(x)
	x = 99 //Gives the variable X a new value of 99

	y := x + 1 //Declares and assigns the value of X plus one and assigns it to the variable Y

	z := x * y           //Declares and multiplies the value of x and y and assigns it to the variable Z
	fmt.Println(x, y, z) //Prints x, y and z.
}
