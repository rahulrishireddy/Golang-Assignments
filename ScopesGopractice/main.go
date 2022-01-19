package main

import (
	"fmt"

	"github.com/rahulrishireddy/ScopesGopractice/scopes"
)

func main() {
	scopes.Scopes()
	scopes.Scopes1()
	scopes.Boolvariables()    //Boolean variable decleration
	scopes.Stringvariables()  // Used to define string variables
	scopes.IntegerVariables() //  various integers types are defined.
	scopes.ArthOperation()    // arthemeticoperations are defined.

	//fmt.Println("c is", c) // since c is declared in different package we are not able to use c in this package

	fmt.Println("Globalvar is", scopes.Globalvar) // Global declaration
}
