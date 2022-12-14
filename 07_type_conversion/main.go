package main

import "fmt"

func main() {

	/* x := 10
	y := 10.0

	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", y, y) */

	/* var x int16 = 128
	var y int8

	y = int8(x) //type(value)

	fmt.Println(y)
	fmt.Printf("%T \n", y)
	fmt.Printf("%T \n", x) */

	/* x := 10
	y := "10"

	fmt.Printf("%T \n", y)
	fmt.Printf("%T \n", x)

	fmt.Println(x + int(y))
	*/
	num1 := 35
	str1 := string(num1)

	fmt.Printf("%v %T\n", str1, str1)
	fmt.Println()
	fmt.Printf("%v %T\n", num1, num1)

	// type conversion kullanarak
	// string'i int'e dönüştüremeyiz
}
