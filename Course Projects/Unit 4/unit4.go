package main

import (
	"fmt" //The standard input/output package for Go
)

var f = 5 //Declares a global variable named A and assigns it to the value 5. (Initialization)
var g int //Declares that there is a global variable with the name Z and that the variable is of type int. Assigns the "zero value" of the integer type to "z"

//The "zero value" is essentially a default value. For instance, the zero value for numbers is 0 and the zero value for booleans is false.

//[NOTE] Only use var for variables outside of a function body

func main() {
	n, _ := fmt.Println("Hello, playground", 42, true)
	fmt.Println(n)
	x := 42 //Declares and assigns the value 42 to the variable X
	fmt.Println(x)
	x = 99 //Gives the variable X a new value of 99

	y := x + a //Declares and assigns the value of X plus one and assigns it to the variable Y

	z := x * y           //Declares and multiplies the value of x and y and assigns it to the variable Z
	fmt.Println(x, y, z) //Prints x, y and z.

	fmt.Printf("%T\n", x) //Returns the type of x (int)

	//[NOTE] Using 'single quotes' for a string essentially says give me everything that's inside of the quotes and
	//turn in into a string so you could have a double quote inside of a single to quote something somebody said for example.
}
