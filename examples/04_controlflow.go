package main

import "fmt"

func main() {
	fmt.Println("=== Contoh Kontrol Alur (Control Flow) ===\n")

	// IF Statement
	fmt.Println("=== IF Statement ===")
	age := 20
	if age >= 18 {
		fmt.Printf("Umur %d: Dewasa\n", age)
	} else {
		fmt.Printf("Umur %d: Anak-anak\n", age)
	}

	// IF dengan else if
	score := 85
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else if score >= 60 {
		fmt.Println("Grade: D")
	} else {
		fmt.Println("Grade: F")
	}

	// IF dengan short statement
	if value := 100; value > 50 {
		fmt.Printf("Nilai %d lebih besar dari 50\n", value)
	}

	// SWITCH Statement
	fmt.Println("\n=== SWITCH Statement ===")
	day := "Rabu"
	switch day {
	case "Senin":
		fmt.Println("Awal minggu")
	case "Selasa", "Rabu", "Kamis":
		fmt.Println("Pertengahan minggu")
	case "Jumat":
		fmt.Println("Akhir minggu kerja")
	case "Sabtu", "Minggu":
		fmt.Println("Akhir pekan")
	default:
		fmt.Println("Hari tidak valid")
	}

	// Switch tanpa kondisi (seperti if-else chain)
	hour := 14
	switch {
	case hour < 12:
		fmt.Println("Selamat pagi")
	case hour < 17:
		fmt.Println("Selamat siang")
	case hour < 21:
		fmt.Println("Selamat sore")
	default:
		fmt.Println("Selamat malam")
	}

	// Switch dengan short statement
	switch time := "malam"; time {
	case "pagi":
		fmt.Println("Sarapan")
	case "siang":
		fmt.Println("Makan siang")
	case "malam":
		fmt.Println("Makan malam")
	}

	// FOR Loop
	fmt.Println("\n=== FOR Loop ===")

	// For loop standar
	fmt.Println("Counting 1-5:")
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// For sebagai while loop
	count := 0
	fmt.Println("\nWhile-style loop:")
	for count < 3 {
		fmt.Printf("Count: %d\n", count)
		count++
	}

	// Infinite loop dengan break
	fmt.Println("\nInfinite loop dengan break:")
	n := 0
	for {
		if n >= 3 {
			break
		}
		fmt.Printf("n = %d\n", n)
		n++
	}

	// Continue statement
	fmt.Println("\nContinue - skip genap:")
	for i := 1; i <= 5; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// Range loop dengan array
	fmt.Println("\n=== Range Loop ===")
	numbers := []int{10, 20, 30, 40, 50}
	fmt.Println("Iterasi dengan index dan value:")
	for index, value := range numbers {
		fmt.Printf("Index %d: Value %d\n", index, value)
	}

	// Range hanya index
	fmt.Println("\nHanya index:")
	for index := range numbers {
		fmt.Printf("%d ", index)
	}
	fmt.Println()

	// Range hanya value (skip index dengan _)
	fmt.Println("\nHanya value:")
	for _, value := range numbers {
		fmt.Printf("%d ", value)
	}
	fmt.Println()

	// Range dengan string
	fmt.Println("\nRange dengan string:")
	text := "Hello"
	for index, char := range text {
		fmt.Printf("Index %d: %c\n", index, char)
	}

	// Range dengan map
	fmt.Println("\nRange dengan map:")
	fruits := map[string]int{
		"apel":   5,
		"pisang": 3,
		"jeruk":  7,
	}
	for name, quantity := range fruits {
		fmt.Printf("%s: %d\n", name, quantity)
	}

	// Nested loops
	fmt.Println("\n=== Nested Loops ===")
	fmt.Println("Tabel perkalian 3x3:")
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("%d x %d = %d\t", i, j, i*j)
		}
		fmt.Println()
	}

	// Label dan break ke label
	fmt.Println("\n=== Label dengan Break ===")
outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				fmt.Println("Breaking outer loop at i=1, j=1")
				break outer
			}
			fmt.Printf("i=%d, j=%d\n", i, j)
		}
	}

	// Goto (jarang digunakan, tapi tersedia)
	fmt.Println("\n=== GOTO Statement ===")
	x := 0
loop:
	fmt.Printf("x = %d\n", x)
	x++
	if x < 3 {
		goto loop
	}
	fmt.Println("Done with goto")
}
