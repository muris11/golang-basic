package main

import "fmt"

func main(){
	fmt.Println(10 + 5) // operasi penjumlahan
	fmt.Println(10 - 5) // operasi pengurangan
	fmt.Println(10 * 5) // operasi perkalian
	fmt.Println(10 / 5) // operasi pembagian
	fmt.Println(10 % 5) // operasi modulus

	var a, b, c = 10, 5, 7
	fmt.Println(a + b + c) // operasi penjumlahan dengan lebih dari dua operand
	fmt.Println(a - b - c) // operasi pengurangan dengan lebih dari dua operand
	fmt.Println(a * b * c) // operasi perkalian dengan lebih dari dua operand
	fmt.Println(a / b / c) // operasi pembagian dengan lebih dari dua operand
	fmt.Println(a % b % c) // operasi modulus dengan lebih dari dua operand

	var d = a + b + c
	fmt.Println(d) // operasi penjumlahan dengan lebih dari dua operand

	i := 5 // deklarasi variabel dengan tipe data integer
	i++ // operasi penambahan 1	
	i = i + 1

	j := 5 // deklarasi variabel dengan tipe data integer
	j--

	fmt.Println(i) // mencetak nilai variabel i setelah operasi penambahan 1
	fmt.Println(j) // mencetak nilai variabel j setelah operasi pengurangan 1
}