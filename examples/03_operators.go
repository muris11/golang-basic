package main

import "fmt"

func main() {
	fmt.Println("=== Contoh Operator di Golang ===\n")

	// Operator Aritmatika
	fmt.Println("=== Operator Aritmatika ===")
	a, b := 10, 3
	fmt.Printf("a = %d, b = %d\n", a, b)
	fmt.Printf("Penjumlahan: %d + %d = %d\n", a, b, a+b)
	fmt.Printf("Pengurangan: %d - %d = %d\n", a, b, a-b)
	fmt.Printf("Perkalian: %d * %d = %d\n", a, b, a*b)
	fmt.Printf("Pembagian: %d / %d = %d\n", a, b, a/b)
	fmt.Printf("Modulus: %d %% %d = %d\n", a, b, a%b)

	// Increment dan Decrement
	counter := 5
	fmt.Printf("\nCounter awal: %d\n", counter)
	counter++
	fmt.Printf("Setelah counter++: %d\n", counter)
	counter--
	fmt.Printf("Setelah counter--: %d\n", counter)

	// Operator Assignment
	fmt.Println("\n=== Operator Assignment ===")
	x := 10
	fmt.Printf("x awal: %d\n", x)
	x += 5
	fmt.Printf("x += 5: %d\n", x)
	x -= 3
	fmt.Printf("x -= 3: %d\n", x)
	x *= 2
	fmt.Printf("x *= 2: %d\n", x)
	x /= 4
	fmt.Printf("x /= 4: %d\n", x)
	x %= 3
	fmt.Printf("x %%= 3: %d\n", x)

	// Operator Perbandingan
	fmt.Println("\n=== Operator Perbandingan ===")
	num1, num2 := 15, 20
	fmt.Printf("num1 = %d, num2 = %d\n", num1, num2)
	fmt.Printf("num1 == num2: %t\n", num1 == num2)
	fmt.Printf("num1 != num2: %t\n", num1 != num2)
	fmt.Printf("num1 < num2: %t\n", num1 < num2)
	fmt.Printf("num1 > num2: %t\n", num1 > num2)
	fmt.Printf("num1 <= num2: %t\n", num1 <= num2)
	fmt.Printf("num1 >= num2: %t\n", num1 >= num2)

	// Operator Logika
	fmt.Println("\n=== Operator Logika ===")
	isAdult := true
	hasLicense := false
	fmt.Printf("isAdult = %t, hasLicense = %t\n", isAdult, hasLicense)
	fmt.Printf("isAdult && hasLicense: %t\n", isAdult && hasLicense)
	fmt.Printf("isAdult || hasLicense: %t\n", isAdult || hasLicense)
	fmt.Printf("!isAdult: %t\n", !isAdult)
	fmt.Printf("!hasLicense: %t\n", !hasLicense)

	// Operator Bitwise
	fmt.Println("\n=== Operator Bitwise ===")
	p, q := 12, 10 // 12 = 1100, 10 = 1010 dalam binary
	fmt.Printf("p = %d (binary: %b), q = %d (binary: %b)\n", p, p, q, q)
	fmt.Printf("p & q (AND): %d (binary: %b)\n", p&q, p&q)
	fmt.Printf("p | q (OR): %d (binary: %b)\n", p|q, p|q)
	fmt.Printf("p ^ q (XOR): %d (binary: %b)\n", p^q, p^q)
	fmt.Printf("p &^ q (AND NOT): %d (binary: %b)\n", p&^q, p&^q)
	fmt.Printf("p << 1 (Left Shift): %d (binary: %b)\n", p<<1, p<<1)
	fmt.Printf("p >> 1 (Right Shift): %d (binary: %b)\n", p>>1, p>>1)

	// Operator Precedence (urutan operasi)
	fmt.Println("\n=== Operator Precedence ===")
	result := 2 + 3*4
	fmt.Printf("2 + 3 * 4 = %d (perkalian dulu)\n", result)
	result = (2 + 3) * 4
	fmt.Printf("(2 + 3) * 4 = %d (kurung dulu)\n", result)

	// Complex expression
	age := 25
	hasID := true
	canVote := age >= 17 && hasID
	fmt.Printf("\nUmur: %d, Punya ID: %t\n", age, hasID)
	fmt.Printf("Bisa voting: %t\n", canVote)
}
