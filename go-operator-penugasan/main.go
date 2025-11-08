package main

import "fmt"

func main() {
	var a uint8 = 3 // contoh penggunaan operator penugasan sederhana
	var b, c, d, e, f uint8 = 4, 5, 6, 7, 8 // contoh penggunaan operator penugasan gabungan

	b += 5 // b = b + 5
	c -= 5 // c = c - 5
	d *= 5 // d = d * 5
	e /= 5 // e = e / 5
	f %= 5 // f = f % 5

	fmt.Println("Nilai a: ", a) // menampilkan nilai a
	fmt.Println("Nilai b: ", b) // menampilkan nilai b
	fmt.Println("Nilai c: ", c) // menampilkan nilai c
	fmt.Println("Nilai d: ", d) // menampilkan nilai d
	fmt.Println("Nilai e: ", e) // menampilkan nilai e
	fmt.Println("Nilai f: ", f) // menampilkan nilai f
}