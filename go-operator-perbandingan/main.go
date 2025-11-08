package main

import "fmt"

func main(){
	var x uint8 = 10 // contoh penggunaan operator perbandingan
	var y uint8 = 10// contoh penggunaan operator perbandingan

	var txt1 = "Hello"
	var txt2 = "Go"

	fmt.Println(x == y) // mengecek apakah x sama dengan y
	fmt.Println(x != y) // mengecek apakah x tidak sama dengan y
	fmt.Println(x > y) // mengecek apakah x lebih besar dari y
	fmt.Println(x < y) // mengecek apakah x lebih kecil dari y
	fmt.Println(x >= y) // mengecek apakah x lebih besar atau sama dengan y
	fmt.Println(x <= y) // mengecek apakah x lebih kecil atau sama dengan y

	var result bool = txt1 == txt2 // mengecek apakah txt1 sama dengan txt2
	fmt.Println(result)
}