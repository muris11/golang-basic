package main

import "fmt"

func main(){
	numbers := []int{1, 2, 3} // deklarasi slice dengan tipe data integer

	fmt.Println(numbers[0]) // mencetak elemen pertama dari slice
	fmt.Println(numbers[1]) // mencetak elemen kedua dari slice

	numbers2 := []int{4, 5, 6} // deklarasi slice dengan tipe data integer
	numbers2[2] = 7 // mengubah elemen ketiga dari slice
	fmt.Println(numbers2) // mencetak elemen ketiga dari slice

	iniSlice := []int{10, 20, 30, 40, 50} // deklarasi slice dengan tipe data integer
	fmt.Printf("Ini Slice = %v\n", iniSlice) // mencetak nilai slice iniSlice
	fmt.Printf("Length %d\n", len(iniSlice)) // mencetak panjang slice iniSlice
	fmt.Printf("Capacity %d\n", cap(iniSlice)) // mencetak kapasitas slice iniSlice
	
	iniSlice = append(iniSlice, 70, 80)
	fmt.Printf("Ini Slice = %v\n", iniSlice) // mencetak nilai slice iniSlice
	fmt.Printf("Length %d\n", len(iniSlice)) // mencetak panjang slice iniSlice
	fmt.Printf("Capacity %d\n", cap(iniSlice)) // mencetak kapasitas slice iniSlice


	iniArray := [6]int{10, 20, 30, 40, 50, 60} // deklarasi array dengan tipe data integer
	iniSliceArray := iniArray[1:5] // deklarasi slice dengan tipe data integer menggunakan array iniArray

	fmt.Printf("Ini Slice Array = %v\n", iniSliceArray) // mencetak nilai slice iniSliceArray
	fmt.Printf("Length %d\n", len(iniSliceArray)) // mencetak panjang slice iniSliceArray
	fmt.Printf("Capacity %d\n", cap(iniSliceArray)) // mencetak kapasitas slice iniSliceArray

	iniSliceArray = iniArray[1:3] // mengubah range slice iniSliceArray
	fmt.Printf("Ini Slice Array = %v\n", iniSliceArray)
	fmt.Printf("Length %d\n", len(iniSliceArray)) // mencetak panjang slice iniSliceArray
	fmt.Printf("Capacity %d\n", cap(iniSliceArray)) // mencetak kapasitas slice iniSliceArray

	iniSliceArray = append(iniSliceArray, 20, 21, 22, 23) // menambahkan elemen pada slice iniSliceArray
	fmt.Printf("Ini Slice Array = %v\n", iniSliceArray) // mencetak nilai slice iniSliceArray
	fmt.Printf("Length %d\n", len(iniSliceArray)) // mencetak panjang slice iniSliceArray
	fmt.Printf("Capacity %d\n", cap(iniSliceArray)) // mencetak kapasitas slice iniSliceArray

	angka := []int{1, 2,3 , 4, 5, 6, 7, 8, 9, 10} // deklarasi slice dengan tipe data integer
	fmt.Printf("Angka = %v\n", angka) // mencetak nilai slice angka
	fmt.Printf("Length %d\n", len(angka)) // mencetak panjang slice angka	
	fmt.Printf("Capacity %d\n", cap(angka)) // mencetak kapasitas slice angka

	copyAngka := make([]int, len(angka), cap(angka)) // membuat slice baru dengan panjang dan kapasitas sama dengan slice angka
	copy(copyAngka, angka) // menyalin isi slice angka ke slice copyAngka
	fmt.Printf("Copy Angka = %v\n", copyAngka) // mencetak nilai slice copyAngka
	fmt.Printf("Length %d\n", len(copyAngka)) // mencetak panjang slice copyAngka
	fmt.Printf("Capacity %d\n", cap(copyAngka)) // mencetak kapasitas slice copyAngka




}