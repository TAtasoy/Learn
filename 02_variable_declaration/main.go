//Package clause
package main

//Import Statement
import "fmt"

//My Code
func main() {
	var name string = "Tugberk" // var - name of variable "name" - static type "string" bu örnekte

	// name := "Tugberk" // değişken tipini belirtmeme gerek yok
	// age := 23 // kısa gösterim olarak geçiyor , static type özelliğini korumaya devam ediyor.
	var age int
	var isMarried bool
	isMarried = true
	age = 23
	// Firstname = "Tugberk" , tanımlanmış olan değişkene değer atıyoruz burda string olduğu için "Tugberk"

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isMarried)

}
