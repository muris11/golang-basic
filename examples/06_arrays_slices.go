package main

import "fmt"

func main() {
	fmt.Println("=== Contoh Array dan Slice ===\n")

	// ARRAY
	fmt.Println("=== ARRAY ===")

	// Deklarasi array dengan ukuran tetap
	var numbers [5]int
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50
	fmt.Printf("Array numbers: %v\n", numbers)

	// Inisialisasi array dengan nilai
	fruits := [3]string{"Apple", "Banana", "Orange"}
	fmt.Printf("Array fruits: %v\n", fruits)

	// Array dengan ukuran otomatis
	colors := [...]string{"Red", "Green", "Blue", "Yellow"}
	fmt.Printf("Array colors: %v (length: %d)\n", colors, len(colors))

	// Akses element array
	fmt.Printf("First fruit: %s\n", fruits[0])
	fmt.Printf("Last fruit: %s\n", fruits[len(fruits)-1])

	// Iterasi array
	fmt.Print("Iterate colors: ")
	for i := 0; i < len(colors); i++ {
		fmt.Printf("%s ", colors[i])
	}
	fmt.Println()

	// Iterasi dengan range
	fmt.Println("Fruits dengan range:")
	for index, fruit := range fruits {
		fmt.Printf("  Index %d: %s\n", index, fruit)
	}

	// Multi-dimensional array
	var matrix [3][3]int
	matrix[0][0] = 1
	matrix[1][1] = 5
	matrix[2][2] = 9
	fmt.Printf("Matrix: %v\n", matrix)

	// SLICE
	fmt.Println("\n=== SLICE ===")

	// Membuat slice
	var slice []int
	fmt.Printf("Empty slice: %v (len: %d, cap: %d)\n", slice, len(slice), cap(slice))

	// Slice literal
	names := []string{"Alice", "Bob", "Charlie"}
	fmt.Printf("Slice names: %v (len: %d, cap: %d)\n", names, len(names), cap(names))

	// Append ke slice
	names = append(names, "David")
	fmt.Printf("After append: %v (len: %d, cap: %d)\n", names, len(names), cap(names))

	// Append multiple values
	names = append(names, "Eve", "Frank")
	fmt.Printf("After append multiple: %v (len: %d, cap: %d)\n", names, len(names), cap(names))

	// Make slice dengan capacity
	scores := make([]int, 5, 10) // length 5, capacity 10
	fmt.Printf("Slice with make: %v (len: %d, cap: %d)\n", scores, len(scores), cap(scores))

	// Set values
	for i := 0; i < len(scores); i++ {
		scores[i] = (i + 1) * 10
	}
	fmt.Printf("Scores after set: %v\n", scores)

	// Slicing operation
	fmt.Println("\n=== Slicing Operation ===")
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("Original: %v\n", nums)
	fmt.Printf("nums[2:5]: %v\n", nums[2:5])   // index 2 sampai 4
	fmt.Printf("nums[:4]: %v\n", nums[:4])     // dari awal sampai index 3
	fmt.Printf("nums[6:]: %v\n", nums[6:])     // dari index 6 sampai akhir
	fmt.Printf("nums[:]: %v\n", nums[:])       // semua elemen
	fmt.Printf("nums[1:8:9]: %v (cap: %d)\n", nums[1:8:9], cap(nums[1:8:9])) // [low:high:max]

	// Copy slice
	fmt.Println("\n=== Copy Slice ===")
	source := []int{1, 2, 3, 4, 5}
	destination := make([]int, len(source))
	copied := copy(destination, source)
	fmt.Printf("Source: %v\n", source)
	fmt.Printf("Destination: %v\n", destination)
	fmt.Printf("Elements copied: %d\n", copied)

	// Modifikasi copy tidak mempengaruhi original
	destination[0] = 100
	fmt.Printf("After modify destination[0]: source=%v, dest=%v\n", source, destination)

	// Slice of slices
	fmt.Println("\n=== 2D Slice ===")
	matrix2D := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("2D Matrix:")
	for i, row := range matrix2D {
		fmt.Printf("Row %d: %v\n", i, row)
	}

	// Remove element dari slice
	fmt.Println("\n=== Remove Element ===")
	items := []string{"a", "b", "c", "d", "e"}
	fmt.Printf("Original: %v\n", items)

	// Remove index 2 (element "c")
	indexToRemove := 2
	items = append(items[:indexToRemove], items[indexToRemove+1:]...)
	fmt.Printf("After remove index 2: %v\n", items)

	// Insert element
	fmt.Println("\n=== Insert Element ===")
	letters := []string{"a", "b", "d", "e"}
	fmt.Printf("Original: %v\n", letters)

	// Insert "c" at index 2
	insertIndex := 2
	letters = append(letters[:insertIndex], append([]string{"c"}, letters[insertIndex:]...)...)
	fmt.Printf("After insert 'c' at index 2: %v\n", letters)

	// Filter slice
	fmt.Println("\n=== Filter Slice ===")
	allNumbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var evenNumbers []int
	for _, num := range allNumbers {
		if num%2 == 0 {
			evenNumbers = append(evenNumbers, num)
		}
	}
	fmt.Printf("All: %v\n", allNumbers)
	fmt.Printf("Even: %v\n", evenNumbers)

	// Nil slice vs empty slice
	fmt.Println("\n=== Nil vs Empty Slice ===")
	var nilSlice []int
	emptySlice := []int{}
	fmt.Printf("Nil slice: %v (len: %d, cap: %d, is nil: %t)\n", nilSlice, len(nilSlice), cap(nilSlice), nilSlice == nil)
	fmt.Printf("Empty slice: %v (len: %d, cap: %d, is nil: %t)\n", emptySlice, len(emptySlice), cap(emptySlice), emptySlice == nil)

	// Append slices together
	fmt.Println("\n=== Append Slices ===")
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}
	combined := append(slice1, slice2...)
	fmt.Printf("Slice1: %v\n", slice1)
	fmt.Printf("Slice2: %v\n", slice2)
	fmt.Printf("Combined: %v\n", combined)
}
