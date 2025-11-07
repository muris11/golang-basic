package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	fmt.Println("=== Contoh Error Handling di Golang ===\n")

	// Basic error handling
	fmt.Println("=== Basic Error ===")
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 / 2 = %.2f\n", result)
	}

	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println()

	// Errors.New
	fmt.Println("=== Creating Errors ===")
	err1 := errors.New("something went wrong")
	fmt.Println("Error 1:", err1)

	// fmt.Errorf untuk error dengan format
	value := 42
	err2 := fmt.Errorf("invalid value: %d", value)
	fmt.Println("Error 2:", err2)
	fmt.Println()

	// Custom error type
	fmt.Println("=== Custom Error Type ===")
	err3 := &ValidationError{
		Field:   "email",
		Message: "invalid email format",
	}
	fmt.Println("Error 3:", err3)

	// Type assertion untuk custom error
	if ve, ok := err3.(*ValidationError); ok {
		fmt.Printf("Validation failed on field '%s': %s\n", ve.Field, ve.Message)
	}
	fmt.Println()

	// Multiple return dengan error
	fmt.Println("=== Multiple Return with Error ===")
	user, err := getUser(1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("User: %+v\n", user)
	}

	_, err = getUser(999)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println()

	// Error wrapping (Go 1.13+)
	fmt.Println("=== Error Wrapping ===")
	err4 := processData()
	if err4 != nil {
		fmt.Println("Error:", err4)
		
		// Check wrapped error
		if errors.Is(err4, ErrNotFound) {
			fmt.Println("This is a 'not found' error")
		}
	}
	fmt.Println()

	// Defer for cleanup
	fmt.Println("=== Defer for Cleanup ===")
	err = writeToFile("/tmp/test.txt", "Hello, World!")
	if err != nil {
		fmt.Println("Error writing file:", err)
	} else {
		fmt.Println("File written successfully")
	}
	fmt.Println()

	// Panic and Recover
	fmt.Println("=== Panic and Recover ===")
	fmt.Println("Starting risky operation...")
	safeDivide(10, 2)
	safeDivide(10, 0)
	fmt.Println("Program continues after panic recovery")
	fmt.Println()

	// Multiple error checks
	fmt.Println("=== Multiple Error Checks ===")
	if err := step1(); err != nil {
		fmt.Println("Step 1 failed:", err)
		return
	}
	if err := step2(); err != nil {
		fmt.Println("Step 2 failed:", err)
		return
	}
	if err := step3(); err != nil {
		fmt.Println("Step 3 failed:", err)
		return
	}
	fmt.Println("All steps completed successfully")
	fmt.Println()

	// Error handling in loop
	fmt.Println("=== Error Handling in Loop ===")
	processItems()
	fmt.Println()

	// Sentinel errors
	fmt.Println("=== Sentinel Errors ===")
	err = performOperation()
	if errors.Is(err, ErrNotFound) {
		fmt.Println("Handling not found error")
	} else if errors.Is(err, ErrUnauthorized) {
		fmt.Println("Handling unauthorized error")
	} else if err != nil {
		fmt.Println("Unknown error:", err)
	}
	fmt.Println()

	// Custom error with additional context
	fmt.Println("=== Error with Context ===")
	opErr := &OperationError{
		Op:   "database query",
		Path: "/users/123",
		Err:  errors.New("connection timeout"),
	}
	fmt.Println("Operation error:", opErr)
	fmt.Println()

	// Best practices
	fmt.Println("=== Best Practices Demo ===")
	demonstrateBestPractices()
}

// Basic error function
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// Custom error type
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error on field '%s': %s", e.Field, e.Message)
}

// User struct untuk example
type User struct {
	ID   int
	Name string
}

func getUser(id int) (*User, error) {
	if id <= 0 {
		return nil, errors.New("invalid user ID")
	}
	if id == 999 {
		return nil, errors.New("user not found")
	}
	return &User{ID: id, Name: "John Doe"}, nil
}

// Sentinel errors
var (
	ErrNotFound     = errors.New("not found")
	ErrUnauthorized = errors.New("unauthorized")
	ErrInvalidInput = errors.New("invalid input")
)

func processData() error {
	// Simulasi error
	return fmt.Errorf("failed to process: %w", ErrNotFound)
}

func performOperation() error {
	// Simulasi error
	return ErrNotFound
}

// File operations dengan defer
func writeToFile(filename, data string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close() // Always close file

	_, err = file.WriteString(data)
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}

// Panic and recover
func safeDivide(a, b int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	result := a / b
	fmt.Printf("%d / %d = %d\n", a, b, result)
}

// Steps dengan error
func step1() error {
	fmt.Println("Executing step 1...")
	return nil
}

func step2() error {
	fmt.Println("Executing step 2...")
	return nil
}

func step3() error {
	fmt.Println("Executing step 3...")
	return nil
}

// Error handling in loop
func processItems() {
	items := []int{1, 2, 3, 4, 5}
	
	for i, item := range items {
		if err := processItem(item); err != nil {
			fmt.Printf("Error processing item %d: %v\n", i, err)
			continue // Skip to next item
		}
		fmt.Printf("Processed item %d successfully\n", i)
	}
}

func processItem(item int) error {
	if item == 3 {
		return errors.New("item 3 cannot be processed")
	}
	return nil
}

// Custom error with context
type OperationError struct {
	Op   string
	Path string
	Err  error
}

func (e *OperationError) Error() string {
	return fmt.Sprintf("%s %s: %v", e.Op, e.Path, e.Err)
}

func (e *OperationError) Unwrap() error {
	return e.Err
}

// Best practices demonstration
func demonstrateBestPractices() {
	// 1. Always check errors
	data, err := readData()
	if err != nil {
		fmt.Println("Error reading data:", err)
		return
	}
	fmt.Println("Data read:", data)

	// 2. Provide context in errors
	err = validateInput("")
	if err != nil {
		fmt.Println("Validation error:", err)
	}

	// 3. Use defer for cleanup
	cleanup()

	// 4. Don't panic in libraries, return errors
	result, err := libraryFunction(10)
	if err != nil {
		fmt.Println("Library error:", err)
	} else {
		fmt.Println("Library result:", result)
	}
}

func readData() (string, error) {
	// Simulasi membaca data
	success := true
	if !success {
		return "", errors.New("failed to read data")
	}
	return "sample data", nil
}

func validateInput(input string) error {
	if input == "" {
		return fmt.Errorf("validation failed: input cannot be empty")
	}
	return nil
}

func cleanup() {
	defer fmt.Println("Cleanup completed")
	fmt.Println("Performing cleanup...")
}

func libraryFunction(value int) (int, error) {
	if value < 0 {
		return 0, errors.New("value cannot be negative")
	}
	return value * 2, nil
}
