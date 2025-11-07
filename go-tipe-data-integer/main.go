package main

import "fmt"

func main() {
	var a int64 = 600   // deklarasi variabel dengan tipe data integer
	var b int64 = -4600 // deklarasi variabel dengan tipe data integer

	var c uint = 600	// deklarasi variabel dengan tipe data unsigned integer
	var d uint = 4600	// deklarasi variabel dengan tipe data unsigned integer

	fmt.Printf("Tipe data: %T, Value: %v\n", a, a) // mencetak nilai variabel a
	fmt.Printf("Tipe data: %T, Value: %v\n", b, b) // mencetak nilai variabel b

	fmt.Printf("Tipe data: %T, Value: %v\n", c, c) // mencetak nilai variabel c
	fmt.Printf("Tipe data: %T, Value: %v\n", d, d) // mencetak nilai variabel d
}