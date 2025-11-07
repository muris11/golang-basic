package main

import "fmt"

func main(){
	var a bool = true // deklarasi variabel dengan tipe data boolean
	var b = true // deklarasi variabel tanpa tipe data (tipe data akan ditentukan secara otomatis oleh Go)
	var c bool // deklarasi variabel tanpa inisialisasi (nilai default adalah false)
	d := true // deklarasi variabel menggunakan short declaration

	fmt.Println(a) // mencetak nilai variabel a
	fmt.Println(b) // mencetak nilai variabel b
	fmt.Println(c) // mencetak nilai variabel c
	fmt.Println(d) // mencetak nilai variabel d
	fmt.Printf("%v, %T\n," , b, b) // mencetak nilai dan tipe data variabel b
}