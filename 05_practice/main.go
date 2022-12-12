// 1-) StudentName >>> John Doe , grade >>> 77 IsPassed >>> true değişkenlerini 3 farklı yöntemle oluştur çıktısını yazdır.

package main

import "fmt"

func main() {
	/* 1 Cevabı :
	 var StudentName = "John Doe"
	StudentName = "John Doe" */
	StudentName := "John Doe"

	/* var grade int = 77
	grade = 77 */
	grade := 77

	/* var IsPassed bool = true
	IsPassed = true */
	IsPassed := true

	fmt.Println(StudentName)
	fmt.Println(grade)
	fmt.Println(IsPassed)

}

//2-) Yukardaki belirtilenleri değişiklikleri tek satır içerisinde tanımlayınız.
//2 Cevabı: var StudentName , grade , IsPassed = "John Doe" , 77 , true
/* fmt.Println(StudentName)
fmt.Println(grade)
fmt.Println(IsPassed) */

//3-) "Declaration" , "Assign" , "Initialization" , "Initial Value" kavramlarını açıklayınız. (Terminolojisi)
// Tanımlama yapmak data type belirlemek , "= aslında assign etmektir." Initialization tamamı olarak geçiyor. Initial Value başlangıç değeri

//4-) "Statically Typed" vs "Dynamically Typed" ifadelerini GO ve Phyton üzerinden gösteriniz.
//5-) ":=" vs "=" aradaki farkı gösteriniz.
