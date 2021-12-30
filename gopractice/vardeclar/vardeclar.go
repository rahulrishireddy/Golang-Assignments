package vardeclar

import "fmt"

func Vardeclar() {

	a := 10 // short handed notation
	fmt.Println("a is", a)
	a = 25
	b := 20 // short handed notation
	fmt.Println("b is", b)

	var c int
	c = 25

	fmt.Println("c is ", c)

	var d int
	d = 25
	fmt.Println("d is ", d)

	var e int = 25 // declaring variable with initial value
	fmt.Println("e is ", e)

	var f int = 35 //declaring variable with initial value
	fmt.Println("f is ", f)

	//Strings declaration

	fname := "rahul" // short handed notation
	fmt.Println("fname is", fname)
	fname = "rahul rishi"

	lname := "Vasikarla" // short handed notation
	fmt.Println("lname is", lname)

}
