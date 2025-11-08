package main

import "fmt"

func main(){
	var a = 15 // contoh penggunaan operator logika

	fmt.Println(a > 16 && a < 5) // operator AND
	fmt.Println(a > 16 || a < 20) // operator OR
	fmt.Println(!(a > 16 || a < 20)) // operator NOT

	var nilaiUjian = 90 // contoh studi kasus penggunaan operator logika
	var nilaiAbsensi = 80 // 100

	var lulusUjian = nilaiUjian > 80 // nilai ujian lebih dari 80
	var lulusAbsen = nilaiAbsensi > 81 // nilai absensi lebih dari 80
	fmt.Println(lulusUjian) // menampilkan hasil kelulusan ujian
	fmt.Println(lulusAbsen) // menampilkan hasil kelulusan absensi

	var lulus = lulusUjian && lulusAbsen // lulus jika nilai ujian dan absensi memenuhi syarat
	

	fmt.Println(nilaiUjian > 80 && nilaiAbsensi > 80) // menampilkan hasil kelulusan jika nilai ujian dan absensi memenuhi syarat (nilaiUjian > 80 && nilaiAbsensi > 80)
	fmt.Println(lulus) // menampilkan hasil kelulusan
}