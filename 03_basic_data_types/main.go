package main

import "fmt"

func main() {
	var name string = "Tugberk"
	var age int = 23
	var isMarried bool = false
	var weight float64 = 72.5

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isMarried)
	fmt.Println(weight)

	fmt.Printf("%T\n", name) //%T data type'ini yazdırıyor
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isMarried)
	fmt.Printf("%T\n", weight)

}
