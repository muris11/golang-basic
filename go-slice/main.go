package main

import "fmt"

func main(){
	var iniSlice1 = []int{1, 2, 3, 4, 5} // deklarasi slice dengan tipe data integer
	fmt.Println(iniSlice1) // mencetak nilai slice iniSlice1
	fmt.Println(len(iniSlice1)) // mencetak panjang slice iniSlice1
	fmt.Println(cap(iniSlice1)) // mencetak kapasitas slice iniSlice1

	iniSlice2 := []string{
		"Rifqy", 
		"Saputra",
		"Rifqy Saputra",
	} // deklarasi slice dengan tipe data string

	fmt.Println(iniSlice2) // mencetak nilai slice iniSlice2
	fmt.Println(len(iniSlice2)) // mencetak panjang slice iniSlice2
	fmt.Println(cap(iniSlice2)) // mencetak kapasitas slice iniSlice2

	iniArray := [8]int{1, 2, 3, 4, 5, 6, 7, 8} // deklarasi array dengan tipe data integer
	iniSlice3 := iniArray[3:5] // deklarasi slice dengan tipe data integer menggunakan array iniArray

	fmt.Printf("Ini Slice 3 =%v\n", iniSlice3) // mencetak nilai slice iniSlice3
	fmt.Printf("Length %d\n", iniSlice3) // mencetak panjang slice iniSlice3
	fmt.Printf("Capacity %d\n", iniSlice3) // mencetak kapasitas slice iniSlice3

	iniSlice3[0] = 20 // mengubah nilai elemen pertama slice iniSlice3
	fmt.Printf("IniSlice3 = %v\n", iniSlice3) // mencetak nilai slice iniSlice3 setelah diubah
	fmt.Printf("IniArray = %v\n", iniArray) // mencetak nilai array iniArray setelah slice diubah

	iniSlice4 := make([]int, 2, 5) // deklarasi slice dengan tipe data integer menggunakan fungsi make
	iniSlice4[0] = 10 // mengisi elemen pertama slice iniSlice4
	iniSlice4[1] = 20 // mengisi elemen kedua slice iniSlice4

	fmt.Printf("IniSlice 4 = %v\n", iniSlice4) // mencetak nilai slice iniSlice4
	fmt.Printf("Length %d\n", len(iniSlice4)) // mencetak panjang slice iniSlice4
	fmt.Printf("Capacity %d\n", cap(iniSlice4)) // mencetak kapasitas slice iniSlice4

}