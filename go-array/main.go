package main

import "fmt"

func main() {
	var a1 = [3]int{100, 120, 300}   // cara deklarasi array dengan nilai awal
	a2 := [4]int{100, 120, 300, 400} // cara deklarasi array dengan nilai awal tanpa menyebutkan tipe data

	var a3 = [...]int{1, 3, 2} // cara deklarasi array dengan panjang otomatis
	a4 := [...]int{4, 5, 6} // cara deklarasi array dengan panjang otomatis tanpa menyebutkan tipe data

	var languages = [3]string{"GO", "Java", "Python"} // array bertipe string

	price := [3]int{1000, 2000, 3000} // array bertipe integer
	price[2] = 5000                     // mengubah nilai elemen array pada index ke-2

	a5 := [4]int{} 		// array dengan inisialisasi kosong
	a6 := [4]int{1, 2} // elemen yang tidak diinisialisasi akan bernilai 0 secara default
	a7 := [4]int{3, 4, 5, 6} // array dengan nilai awal

	fmt.Println(a1) // menampilkan seluruh isi array a1
	fmt.Println(a2) // menampilkan seluruh isi array a2
	fmt.Println(a3) // menampilkan seluruh isi array a3
	fmt.Println(a4) // menampilkan seluruh isi array a4
	fmt.Println(languages) // menampilkan seluruh isi array languages

	fmt.Println(price[0]) // menampilkan seluruh isi array price
	fmt.Println(price[2]) // menampilkan seluruh isi array price

	fmt.Println(a5) // menampilkan seluruh isi array a5
	fmt.Println(a6) // menampilkan seluruh isi array a6
	fmt.Println(a7) // menampilkan seluruh isi array a7

	fmt.Println(len(a2)) // menampilkan panjang array a2
	fmt.Println(len(languages)) // menampilkan panjang array languages
	
}