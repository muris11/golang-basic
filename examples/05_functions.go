package main

import (
	"fmt"
	"errors"
)

func main() {
	fmt.Println("=== Contoh Fungsi di Golang ===\n")

	// Fungsi sederhana tanpa parameter dan return
	greet()

	// Fungsi dengan parameter
	greetPerson("Alice")

	// Fungsi dengan return value
	sum := add(5, 3)
	fmt.Printf("5 + 3 = %d\n", sum)

	// Fungsi dengan multiple parameters
	fullName := createFullName("John", "Doe")
	fmt.Printf("Full name: %s\n", fullName)

	// Fungsi dengan multiple return values
	result, remainder := divide(17, 5)
	fmt.Printf("17 / 5 = %d dengan sisa %d\n", result, remainder)

	// Fungsi dengan error handling
	quotient, err := safeDivide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 / 2 = %.2f\n", quotient)
	}

	quotient, err = safeDivide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Named return values
	s, p := calculate(4, 5)
	fmt.Printf("Sum: %d, Product: %d\n", s, p)

	// Variadic function
	total := sumNumbers(1, 2, 3, 4, 5)
	fmt.Printf("Sum of 1,2,3,4,5 = %d\n", total)

	// Anonymous function
	fmt.Println("\n=== Anonymous Function ===")
	result = func(a, b int) int {
		return a * b
	}(6, 7)
	fmt.Printf("6 * 7 = %d\n", result)

	// Function sebagai variable
	multiply := func(a, b int) int {
		return a * b
	}
	fmt.Printf("multiply(3, 4) = %d\n", multiply(3, 4))

	// Function sebagai parameter
	fmt.Println("\n=== Function sebagai Parameter ===")
	numbers := []int{1, 2, 3, 4, 5}
	applyOperation(numbers, func(n int) int {
		return n * 2
	})

	// Closure
	fmt.Println("\n=== Closure ===")
	counter := makeCounter()
	fmt.Println("Counter:", counter()) // 1
	fmt.Println("Counter:", counter()) // 2
	fmt.Println("Counter:", counter()) // 3

	// Recursive function
	fmt.Println("\n=== Recursive Function ===")
	fmt.Printf("Factorial of 5: %d\n", factorial(5))
	fmt.Printf("Fibonacci(7): %d\n", fibonacci(7))

	// Defer
	fmt.Println("\n=== Defer ===")
	deferExample()
}

// Fungsi sederhana
func greet() {
	fmt.Println("Hello, World!")
}

// Fungsi dengan parameter
func greetPerson(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// Fungsi dengan return value
func add(a int, b int) int {
	return a + b
}

// Parameter dengan tipe yang sama bisa digabung
func createFullName(firstName, lastName string) string {
	return firstName + " " + lastName
}

// Multiple return values
func divide(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

// Return dengan error handling
func safeDivide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// Named return values
func calculate(a, b int) (sum int, product int) {
	sum = a + b
	product = a * b
	return // naked return
}

// Variadic function (parameter dengan jumlah variable)
func sumNumbers(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// Function yang menerima function sebagai parameter
func applyOperation(numbers []int, operation func(int) int) {
	fmt.Print("Apply operation: ")
	for _, num := range numbers {
		fmt.Printf("%d ", operation(num))
	}
	fmt.Println()
}

// Closure - function yang mengembalikan function
func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// Recursive function - Factorial
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// Recursive function - Fibonacci
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// Defer example
func deferExample() {
	fmt.Println("Start")
	defer fmt.Println("Deferred 1") // dijalankan paling akhir (LIFO)
	defer fmt.Println("Deferred 2")
	defer fmt.Println("Deferred 3")
	fmt.Println("End")
	// Output: Start, End, Deferred 3, Deferred 2, Deferred 1
}
