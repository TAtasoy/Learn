package main

import (
	"fmt"
	"strconv"
)

func main() {

	/* var (
		name      string = "Tugberk"
		age       int    = 23
		isMarried bool   = false

		weight float64 = 72.5
		height int     = 176
	) */

	//var name, age, isMarried, weight, height = "Tugberk", 23, false, 68.8, 176 //initilaze etmek deniyor. 1 le 1 2'yle 2 eşleşiyor
	/* name, age, isMarried, weight, height := "Tugberk", 23, false, 68.8, 176 */

	/* var name string */
	var age int
	var isMarried bool
	var weight float32
	var height int = 65
	var name2 = strconv.Itoa(height) // data type change int to string aq

	fmt.Println(name2) //" " boş bir string değeri var buna zero values denir. stringi tanımladık ama initilaze etmedik.
	// yani bi değişkenimiz var name isminde ama onun içinde bi değer yok.
	fmt.Println(age)       // int değerler için default değerler 0
	fmt.Println(isMarried) // zero values bool >> false oluyor.
	fmt.Println(weight)    // float değerler için yine default değerler 0 oluyor
	fmt.Println(height)    // numeric olanlar >>> 0
	fmt.Printf("%T\n", name2)
}
