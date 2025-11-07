package main

import "fmt"

func main() {
	var firstName, lastName string = "Rifqy", "Saputra" // deklarasi variabel dengan tipe data string 
	var a, b = 100, 200 // deklarasi variabel tanpa tipe data (tipe data akan ditentukan secara otomatis oleh Go)

	var nickName, myAddress = "Rifqy", "Bekasi" // deklarasi beberapa variabel sekaligus tanpa tipe data

	var x string = "Hello, World!" // deklarasi variabel dengan tipe data string
	var y int = 42               // deklarasi variabel dengan tipe data integer

	var angka = 15.5 // deklarasi variabel tanpa tipe data (tipe data akan ditentukan secara otomatis oleh Go)
	var txt = "Hello, Go!" // deklarasi variabel tanpa tipe data (tipe data akan ditentukan secara otomatis oleh Go)
	

	fmt.Print(firstName, " ", lastName, "\n") // mencetak nilai variabel firstName dan lastName dengan baris baru di antara keduanya
	fmt.Print(a, b, "\n") // mencetak nilai variabel a dan b tanpa spasi atau baris baru di antara keduanya

	fmt.Println(nickName, myAddress) // mencetak nilai variabel nickName dan myAddress dengan spasi di antara keduanya dan menambahkan baris baru di akhir

	fmt.Printf("Value x adalah : %v dan tipe data x adalah : %T\n", x, x) // mencetak nilai dan tipe data variabel x menggunakan format spesifik 
	fmt.Printf("Value y adalah : %v dan tipe data y adalah : %T\n", y, y) // mencetak nilai dan tipe data variabel y menggunakan format spesifik

	fmt.Printf("%v%%\n", angka) // mencetak nilai variabel angka diikuti dengan simbol persen (%)
	fmt.Printf("%#v", txt) // mencetak nilai variabel txt dan menambahkan baris baru di akhir

}