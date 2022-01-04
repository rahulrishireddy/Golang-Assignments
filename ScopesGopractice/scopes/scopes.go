package scopes

import "fmt"

var c int = 40

var Globalvar int = 50

func Scopes() {

	var a int = 10

	fmt.Println("a is", a)

}

func Scopes1() {

	b := 20

	fmt.Println("b is", b)
	//fmt.Println("a is", a) //variables created in one function is NOT accessible in another function

	fmt.Println("c is", c) // variables declared outside the function can still be accessed inside the function as it is called in the same package

}
