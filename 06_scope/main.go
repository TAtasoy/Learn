package main

import "fmt"

var PackVar = "Package Scope"

func main() {

	/* if true {
		var BlockVar = "Block Scope"
		fmt.Println(BlockVar)
	} */

	var funcVar = "Func Scope"

	fmt.Println(funcVar)
	fmt.Println(PackVar)
	myFunc()
}

func myFunc() {

	fmt.Println(PackVar)

}
