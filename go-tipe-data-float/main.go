package main

import "fmt"

func main() {
	var a float32 = 255.5 // deklarasi variabel dengan tipe data float32
	var b float32 = 3.4e+38 // deklarasi variabel dengan tipe data float32 menggunakan notasi ilmiah
	var c float32 // deklarasi variabel tanpa inisialisasi (nilai default adalah 0.0)

	var d float64 = 225.55 // deklarasi variabel dengan tipe data float64
	var e float64 = 1.7e+308 // deklarasi variabel dengan tipe data float64 menggunakan notasi ilmiah

	fmt.Printf("Value: %v, Tipe Data: %T\n", a, a) // mencetak nilai variabel a
	fmt.Printf("Value: %v, Tipe Data: %T\n", b, b) // mencetak nilai variabel b
	fmt.Printf("Value: %v, Tipe Data: %T\n", c, c) // mencetak nilai variabel c
	fmt.Printf("Value: %v, Tipe Data: %T\n", d, d) // mencetak nilai variabel d
	fmt.Printf("Value: %v, Tipe Data: %T\n", e, e) // mencetak nilai variabel e
}