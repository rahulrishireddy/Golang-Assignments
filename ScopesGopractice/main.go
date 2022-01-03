package main

import (
	"fmt"

	"github.com/rahulrishireddy/ScopesGopractice/scopes"
)

func main() {

	scopes.Scopes1()
	//fmt.Println("c is", c) // since c is declared in different package we are not able to use c in this package

	fmt.Println("Gvari is", scopes.Gvari) // Global declaration
}
