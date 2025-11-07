package main

import "fmt"

var x int = 11 // variabel global
// y := 20      // variabel global dengan short variable declaration (tidak valid di level paket)

func main() {

	var myName string  // deklarasi dengan tipe data eksplisit
	myName = "Rifqy Saputra" // pengubahan nilai variabel 
	var nickName = "Rifqy" // deklarasi dengan inferensi tipe data
	myAge := 20 // deklarasi dengan short variable declaration

	var a string // deklarasi dengan tipe data eksplisit dan tanpa inisialisasi
	var b int   // deklarasi dengan tipe data eksplisit dan tanpa inisialisasi
	var c bool // deklarasi dengan tipe data eksplisit dan tanpa inisialisasi

	var d, e, f, g  = 1, 2, 3, "Hello Rifqy" // deklarasi multiple variable dengan tipe data 
	
	var (
		firstName string = "Rifqy" // deklarasi multiple variable dengan blok var
		lastName string = "Saputra" // deklarasi multiple variable dengan blok var
		height int = 170 // deklarasi multiple variable dengan blok var
	)

	var (
		_ string = "Rifqy" // deklarasi multiple variable tanpa inisialisasi
		_ string = "Saputra" // deklarasi multiple variable tanpa inisialisasi
		_ int = 170 // deklarasi multiple variable tanpa inisialisasi
	)

	fmt.Println(myName) // penggunaan variabel
	fmt.Println(nickName) // penggunaan variabel
	fmt.Println(myAge) // penggunaan variabel

	fmt.Println(x) // penggunaan variabel 
	// fmt.Println(y) // penggunaan variabel

	fmt.Println(a) // penggunaan variabel dengan nilai default
	fmt.Println(b) // penggunaan variabel dengan nilai default
	fmt.Println(c) // penggunaan variabel dengan nilai default

	fmt.Println(d) // penggunaan multiple variable
	fmt.Println(e) // penggunaan multiple variable
	fmt.Println(f) // penggunaan multiple variable
	fmt.Println(g) // penggunaan multiple variable

	fmt.Println(firstName) // penggunaan multiple variable dengan blok var
	fmt.Println(lastName) // penggunaan multiple variable dengan blok var
	fmt.Println(height) // penggunaan multiple variable dengan blok var
	
}