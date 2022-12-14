package main

import "fmt"

var PackVar = "Package Scope"

func main() {

	var name = "Tugberk"
	name, surname := "Tugberk", "Atasoy" //normalde tekrar kısa gösterimle tanımlasam hata alıcaktım ama yanına ekstra bi değişken atadım hata olmadı.

	fmt.Println(name, surname)
}

func myFunc() {

	fmt.Println(PackVar)

}
