package main

import (
	"fmt"
)

var a int       //Declares a variable a to be of a type int
type hotdog int //Declares a brand new type hotdog that accepts integers
var b hotdog    //Declares a variable b to be of a type hotdog

func main() {
	a = 42                //Assign the number 42 to a
	b = 43                //Assign the number 43 to b, don't forget the new hotdog type we just created accepts integers as it's underlying value.
	fmt.Println(a)        //Prints a
	fmt.Printf("%T\n", a) //Prints the type of A
	fmt.Println(b)        //Prints b
	fmt.Printf("%T\n", b) //Prints the type of B
	a = int(b)            //Converts B which of type hotdog to an integer, remeber that type hotdog accepts integers as it's underlying value (this is called conversion)

	//NOTE: Cannot do something like a=b because A is of type int and B is of type hotdog. Once you declare a variable's type, it will stay as that type.
}
