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

	var value21 uint8 = 129
	fmt.Println("value21 is", value21) // cannot use this sice range of 8 bit integer is from -128 to 127

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

	// var value6 int32 = -2147483649  //
	// fmt.Println("value6 is", value6) //cannot use it since it is out of range again this time it is negative integer range

}

func ArthOperation() {

	intA := 20
	intB := 27
	floatA := 22.5
	floatB := 12.7

	intsum := intA + intB

	fmt.Println("intsum is", intsum) // addition of the integers

	intsub := intA - intB

	fmt.Println("intsub is", intsub) // subtraction of the integers

	intmul := intA * intB

	fmt.Println("intmul is", intmul) // multiplication of the integers

	floatsum := floatA + floatB
	fmt.Println("floatsum is", floatsum) // addition of float

	floatsub := floatA - floatB
	fmt.Println("floatsub is", floatsub) // subtraction of float

	floatMul := floatA * floatB
	fmt.Println("floatMul is", floatMul) // subtraction of float

	i := 10
	j := 27.5

	k := i + int(j)
	fmt.Println("k is", k) // type conversion of float and adding it to int

	floatconv := j + float64(i)
	fmt.Println("floatconv is", floatconv) // type conversion of int and adding it to float
}
