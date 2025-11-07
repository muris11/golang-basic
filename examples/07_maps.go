package main

import "fmt"

func main() {
	fmt.Println("=== Contoh Map di Golang ===\n")

	// Membuat map dengan make
	fmt.Println("=== Membuat Map ===")
	ages := make(map[string]int)
	ages["Alice"] = 25
	ages["Bob"] = 30
	ages["Charlie"] = 35
	fmt.Printf("Ages: %v\n", ages)

	// Map literal
	scores := map[string]int{
		"Math":    90,
		"Physics": 85,
		"Chemistry": 88,
	}
	fmt.Printf("Scores: %v\n", scores)

	// Map dengan value yang lebih kompleks
	students := map[string]map[string]interface{}{
		"Alice": {
			"age":   20,
			"grade": "A",
		},
		"Bob": {
			"age":   22,
			"grade": "B",
		},
	}
	fmt.Printf("Students: %v\n", students)

	// Akses value
	fmt.Println("\n=== Akses Value ===")
	aliceAge := ages["Alice"]
	fmt.Printf("Alice's age: %d\n", aliceAge)

	mathScore := scores["Math"]
	fmt.Printf("Math score: %d\n", mathScore)

	// Check key existence
	fmt.Println("\n=== Check Key Existence ===")
	age, exists := ages["David"]
	if exists {
		fmt.Printf("David's age: %d\n", age)
	} else {
		fmt.Println("David not found in ages map")
	}

	// Menggunakan value tanpa check (return zero value jika tidak ada)
	unknownAge := ages["Unknown"]
	fmt.Printf("Unknown age: %d (zero value)\n", unknownAge)

	// Update value
	fmt.Println("\n=== Update Value ===")
	fmt.Printf("Bob's age before: %d\n", ages["Bob"])
	ages["Bob"] = 31
	fmt.Printf("Bob's age after: %d\n", ages["Bob"])

	// Delete key
	fmt.Println("\n=== Delete Key ===")
	fmt.Printf("Before delete: %v\n", ages)
	delete(ages, "Charlie")
	fmt.Printf("After delete Charlie: %v\n", ages)

	// Iterate map
	fmt.Println("\n=== Iterate Map ===")
	fmt.Println("Scores:")
	for subject, score := range scores {
		fmt.Printf("  %s: %d\n", subject, score)
	}

	// Iterate hanya keys
	fmt.Println("\nSubjects (keys only):")
	for subject := range scores {
		fmt.Printf("  %s\n", subject)
	}

	// Map length
	fmt.Printf("\nNumber of students: %d\n", len(ages))
	fmt.Printf("Number of subjects: %d\n", len(scores))

	// Map of slices
	fmt.Println("\n=== Map of Slices ===")
	classes := map[string][]string{
		"Math":    {"Alice", "Bob", "Charlie"},
		"Physics": {"Alice", "David"},
		"Chemistry": {"Bob", "Eve"},
	}
	fmt.Println("Classes:")
	for subject, students := range classes {
		fmt.Printf("  %s: %v\n", subject, students)
	}

	// Nested maps
	fmt.Println("\n=== Nested Maps ===")
	userProfiles := map[string]map[string]string{
		"alice": {
			"email": "alice@example.com",
			"city":  "Jakarta",
		},
		"bob": {
			"email": "bob@example.com",
			"city":  "Bandung",
		},
	}
	fmt.Println("User Profiles:")
	for username, profile := range userProfiles {
		fmt.Printf("  User: %s\n", username)
		for key, value := range profile {
			fmt.Printf("    %s: %s\n", key, value)
		}
	}

	// Map dengan struct sebagai value
	fmt.Println("\n=== Map with Struct Values ===")
	type Person struct {
		Name string
		Age  int
		City string
	}

	people := map[int]Person{
		1: {Name: "Alice", Age: 25, City: "Jakarta"},
		2: {Name: "Bob", Age: 30, City: "Bandung"},
		3: {Name: "Charlie", Age: 35, City: "Surabaya"},
	}

	fmt.Println("People:")
	for id, person := range people {
		fmt.Printf("  ID %d: %s (%d years) from %s\n", id, person.Name, person.Age, person.City)
	}

	// Check and add if not exists
	fmt.Println("\n=== Check and Add ===")
	inventory := map[string]int{
		"apple":  10,
		"banana": 5,
	}

	// Add orange if not exists
	if _, exists := inventory["orange"]; !exists {
		inventory["orange"] = 7
		fmt.Println("Added orange to inventory")
	}
	fmt.Printf("Inventory: %v\n", inventory)

	// Increment value
	fmt.Println("\n=== Increment Value ===")
	counter := make(map[string]int)
	words := []string{"hello", "world", "hello", "golang", "world", "hello"}

	for _, word := range words {
		counter[word]++
	}

	fmt.Println("Word counter:")
	for word, count := range counter {
		fmt.Printf("  %s: %d\n", word, count)
	}

	// Clear all entries (Go doesn't have built-in clear for maps before Go 1.21)
	fmt.Println("\n=== Clear Map ===")
	temp := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Printf("Before clear: %v\n", temp)
	
	// Method 1: Assign new empty map
	temp = make(map[string]int)
	fmt.Printf("After clear: %v\n", temp)

	// Method 2: Delete all keys
	temp2 := map[string]int{"x": 1, "y": 2, "z": 3}
	fmt.Printf("Before delete all: %v\n", temp2)
	for key := range temp2 {
		delete(temp2, key)
	}
	fmt.Printf("After delete all: %v\n", temp2)

	// Nil map vs empty map
	fmt.Println("\n=== Nil Map vs Empty Map ===")
	var nilMap map[string]int
	emptyMap := make(map[string]int)
	
	fmt.Printf("Nil map: %v (len: %d, is nil: %t)\n", nilMap, len(nilMap), nilMap == nil)
	fmt.Printf("Empty map: %v (len: %d, is nil: %t)\n", emptyMap, len(emptyMap), emptyMap == nil)
	
	// Writing to nil map will panic!
	// nilMap["key"] = 1 // This would cause panic
	emptyMap["key"] = 1 // This is OK
	fmt.Printf("After adding to empty map: %v\n", emptyMap)
}
