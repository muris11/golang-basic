package main

import "fmt"

func main() {
	fmt.Println("=== Contoh Variabel dan Konstanta ===\n")

	// Deklarasi variabel dengan tipe data eksplisit
	var name string = "John Doe"
	var age int = 25
	fmt.Printf("Nama: %s, Umur: %d\n", name, age)

	// Type inference (Go menentukan tipe otomatis)
	var city = "Jakarta"
	var temperature = 28.5
	fmt.Printf("Kota: %s, Suhu: %.1f°C\n", city, temperature)

	// Short declaration (hanya bisa di dalam fungsi)
	country := "Indonesia"
	population := 273000000
	fmt.Printf("Negara: %s, Populasi: %d\n", country, population)

	// Multiple variable declaration
	var x, y int = 10, 20
	fmt.Printf("x: %d, y: %d\n", x, y)

	// Multiple variable dengan tipe berbeda
	var (
		firstName string = "Alice"
		lastName  string = "Smith"
		userAge   int    = 30
		isActive  bool   = true
	)
	fmt.Printf("User: %s %s, Age: %d, Active: %t\n", firstName, lastName, userAge, isActive)

	// Short declaration multiple variables
	a, b, c := 1, 2, 3
	fmt.Printf("a: %d, b: %d, c: %d\n", a, b, c)

	// Reassignment
	name = "Jane Doe"
	age = 28
	fmt.Printf("Updated - Nama: %s, Umur: %d\n", name, age)

	// Konstanta
	fmt.Println("\n=== Konstanta ===")
	const Pi = 3.14159
	const AppName = "Golang Basic"
	const Version = 1.0
	fmt.Printf("App: %s v%.1f\n", AppName, Version)
	fmt.Printf("Pi: %f\n", Pi)

	// Konstanta dengan blok
	const (
		StatusOK          = 200
		StatusNotFound    = 404
		StatusServerError = 500
	)
	fmt.Printf("HTTP Status: OK=%d, NotFound=%d, Error=%d\n", StatusOK, StatusNotFound, StatusServerError)

	// Konstanta dengan iota (auto increment)
	const (
		Monday    = iota + 1 // 1
		Tuesday              // 2
		Wednesday            // 3
		Thursday             // 4
		Friday               // 5
		Saturday             // 6
		Sunday               // 7
	)
	fmt.Printf("Days: Mon=%d, Tue=%d, Wed=%d, Thu=%d, Fri=%d, Sat=%d, Sun=%d\n",
		Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday)

	// iota dengan operasi
	const (
		_  = iota             // 0 (diabaikan)
		KB = 1 << (10 * iota) // 1 << 10 = 1024
		MB                    // 1 << 20 = 1048576
		GB                    // 1 << 30 = 1073741824
	)
	fmt.Printf("Storage: 1 KB=%d bytes, 1 MB=%d bytes, 1 GB=%d bytes\n", KB, MB, GB)

	// Pointer
	fmt.Println("\n=== Pointer ===")
	num := 42
	ptr := &num
	fmt.Printf("Nilai: %d, Alamat: %p, Pointer: %p, Nilai via Pointer: %d\n", num, &num, ptr, *ptr)

	// Mengubah nilai via pointer
	*ptr = 100
	fmt.Printf("Nilai setelah diubah via pointer: %d\n", num)
}
