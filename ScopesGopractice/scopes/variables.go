package scopes

import "fmt"

func Boolvariables() {

	x := true
	y := false
	fmt.Println("x is", x, "y is", y)

	z := x && y
	fmt.Println("z is", z)

}

func Stringvariables() {

	var Name string
	Name = "Rahul"
	fmt.Println("Name is", Name)

}

func IntegerVariables() {

	var r int = 50
	fmt.Println("r is", r)

	var value1 int8 = 50
	fmt.Println("value1 is", value1)

	// var value2 int8 = 129
	// fmt.Println("value2 is", value2)   // cannot use this sice range of 8 bit integer is from -128 to 127

	var value2 int16 = 129
	fmt.Println("value2 is", value2) // the same can be acheived by using int16 since 16 bit has -32768 to +32767

	// var value3 int16 = 32768
	// fmt.Println("value3 is", value3) // will give you an error sice we declared value3 out of range so we need to move to int32

	var value3 int32 = 32768
	fmt.Println("value3 is", value3) // the range of the 32bit is from -2147483648 to 2147483647

	// var value4 int32 = 2147483648
	// fmt.Println("value4 is", value4) // the range of the 32bit is from -2147483648 to 2147483647 // out of range again

	var value4 int64 = 2147483648
	fmt.Println("value4 is", value4) // int64 limit is from -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807

	var value5 int64 = 1
	fmt.Println("value5 is", value5)

}
