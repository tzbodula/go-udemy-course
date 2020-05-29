package main

import "fmt"

var y = 42 //The inital value of Y

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)  //returns the type of Y
	fmt.Printf("%b\n", y)  //returns the value of Y in binary
	fmt.Printf("%x\n", y)  //returns the value of Y in hexadecimal
	fmt.Printf("%#x\n", y) //returns the value of Y in hexadecimal but with a 0x in front of it. (The 0x signals that it's hexadecimal and not plain decimal)
	y = 911                //Assigns a new value to Y
	fmt.Printf("%#x\n", y) //returns the value of Y in hexadecimal but with a 0x in front of it. (The 0x signals that it's hexadecimal and not plain decimal)

	s := fmt.Sprintf("%#x\t%b\t%x", y, y, y) //This line using Sprintf returns the formatted string and assigns it to the variable S
	fmt.Println(s)
}
