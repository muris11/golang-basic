package main

import "fmt"

func main() {
	fmt.Println("=== Contoh Tipe Data di Golang ===\n")

	// Integer
	var age int = 25
	var smallNumber int8 = 127
	var bigNumber int64 = 9223372036854775807
	fmt.Printf("Integer: age=%d, smallNumber=%d, bigNumber=%d\n", age, smallNumber, bigNumber)

	// Unsigned Integer
	var count uint = 100
	var byteValue uint8 = 255
	fmt.Printf("Unsigned Integer: count=%d, byteValue=%d\n", count, byteValue)

	// Float
	var price float32 = 19.99
	var pi float64 = 3.14159265359
	fmt.Printf("Float: price=%.2f, pi=%.10f\n", price, pi)

	// Boolean
	var isActive bool = true
	var isComplete bool = false
	fmt.Printf("Boolean: isActive=%t, isComplete=%t\n", isActive, isComplete)

	// String
	var name string = "Golang"
	var message string = "Belajar pemrograman Go!"
	fmt.Printf("String: name='%s', message='%s'\n", name, message)

	// Byte (alias uint8)
	var b byte = 'A'
	fmt.Printf("Byte: b=%c (ASCII: %d)\n", b, b)

	// Rune (alias int32) - Unicode code point
	var r rune = '世'
	fmt.Printf("Rune: r=%c (Unicode: %d)\n", r, r)

	// Complex numbers
	var c complex64 = 1 + 2i
	var c2 complex128 = 3 + 4i
	fmt.Printf("Complex: c=%v, c2=%v\n", c, c2)

	// Zero values (nilai default)
	fmt.Println("\n=== Zero Values ===")
	var defaultInt int
	var defaultFloat float64
	var defaultBool bool
	var defaultString string
	fmt.Printf("Default int: %d\n", defaultInt)
	fmt.Printf("Default float: %f\n", defaultFloat)
	fmt.Printf("Default bool: %t\n", defaultBool)
	fmt.Printf("Default string: '%s'\n", defaultString)

	// Type conversion
	fmt.Println("\n=== Type Conversion ===")
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)
	fmt.Printf("int: %d -> float64: %f -> uint: %d\n", i, f, u)
}
