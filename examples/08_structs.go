package main

import "fmt"

// Definisi struct
type Person struct {
	Name string
	Age  int
	City string
}

// Struct dengan field tags
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// Nested struct
type Address struct {
	Street  string
	City    string
	ZipCode string
}

type Employee struct {
	Name    string
	Age     int
	Address Address
}

// Embedded struct
type Contact struct {
	Phone string
	Email string
}

type Customer struct {
	Name string
	Contact // embedded
}

func main() {
	fmt.Println("=== Contoh Struct di Golang ===\n")

	// Membuat instance struct
	fmt.Println("=== Membuat Struct ===")
	
	// Cara 1: Dengan field names
	person1 := Person{
		Name: "Alice",
		Age:  25,
		City: "Jakarta",
	}
	fmt.Printf("Person1: %+v\n", person1)

	// Cara 2: Tanpa field names (harus urut)
	person2 := Person{"Bob", 30, "Bandung"}
	fmt.Printf("Person2: %+v\n", person2)

	// Cara 3: Partial initialization
	person3 := Person{Name: "Charlie"}
	fmt.Printf("Person3: %+v\n", person3) // Age dan City akan jadi zero value

	// Cara 4: Zero value struct
	var person4 Person
	fmt.Printf("Person4 (zero value): %+v\n", person4)

	// Akses dan modifikasi field
	fmt.Println("\n=== Akses Field ===")
	fmt.Printf("Name: %s\n", person1.Name)
	fmt.Printf("Age: %d\n", person1.Age)
	fmt.Printf("City: %s\n", person1.City)

	person1.Age = 26
	fmt.Printf("Updated age: %d\n", person1.Age)

	// Pointer ke struct
	fmt.Println("\n=== Pointer ke Struct ===")
	person5 := &Person{
		Name: "David",
		Age:  35,
		City: "Surabaya",
	}
	fmt.Printf("Person5: %+v\n", person5)
	
	// Go otomatis dereference pointer
	person5.Age = 36
	fmt.Printf("Updated via pointer: %+v\n", person5)

	// Anonymous struct
	fmt.Println("\n=== Anonymous Struct ===")
	car := struct {
		Brand string
		Year  int
		Color string
	}{
		Brand: "Toyota",
		Year:  2023,
		Color: "Red",
	}
	fmt.Printf("Car: %+v\n", car)

	// Struct comparison
	fmt.Println("\n=== Struct Comparison ===")
	p1 := Person{Name: "Alice", Age: 25, City: "Jakarta"}
	p2 := Person{Name: "Alice", Age: 25, City: "Jakarta"}
	p3 := Person{Name: "Bob", Age: 30, City: "Bandung"}

	fmt.Printf("p1 == p2: %t\n", p1 == p2)
	fmt.Printf("p1 == p3: %t\n", p1 == p3)

	// Nested struct
	fmt.Println("\n=== Nested Struct ===")
	emp := Employee{
		Name: "John Doe",
		Age:  28,
		Address: Address{
			Street:  "Jl. Sudirman No. 123",
			City:    "Jakarta",
			ZipCode: "12345",
		},
	}
	fmt.Printf("Employee: %+v\n", emp)
	fmt.Printf("Employee city: %s\n", emp.Address.City)

	// Embedded struct
	fmt.Println("\n=== Embedded Struct ===")
	customer := Customer{
		Name: "Jane Smith",
		Contact: Contact{
			Phone: "08123456789",
			Email: "jane@example.com",
		},
	}
	fmt.Printf("Customer: %+v\n", customer)
	// Bisa akses field embedded langsung
	fmt.Printf("Customer phone: %s\n", customer.Phone)
	fmt.Printf("Customer email: %s\n", customer.Email)

	// Struct methods
	fmt.Println("\n=== Struct Methods ===")
	alice := Person{Name: "Alice", Age: 25, City: "Jakarta"}
	alice.Introduce()
	alice.HaveBirthday()
	alice.Introduce()

	// Method dengan pointer receiver
	fmt.Println("\nUsing pointer receiver:")
	bob := &Person{Name: "Bob", Age: 30, City: "Bandung"}
	bob.UpdateCity("Surabaya")
	fmt.Printf("Updated person: %+v\n", bob)

	// Slice of structs
	fmt.Println("\n=== Slice of Structs ===")
	people := []Person{
		{Name: "Alice", Age: 25, City: "Jakarta"},
		{Name: "Bob", Age: 30, City: "Bandung"},
		{Name: "Charlie", Age: 35, City: "Surabaya"},
	}

	fmt.Println("People:")
	for i, person := range people {
		fmt.Printf("  %d: %s (%d) from %s\n", i+1, person.Name, person.Age, person.City)
	}

	// Map of structs
	fmt.Println("\n=== Map of Structs ===")
	userMap := map[int]User{
		1: {ID: 1, Username: "alice", Email: "alice@example.com"},
		2: {ID: 2, Username: "bob", Email: "bob@example.com"},
	}

	fmt.Println("Users:")
	for id, user := range userMap {
		fmt.Printf("  ID %d: %s (%s)\n", id, user.Username, user.Email)
	}

	// Constructor pattern
	fmt.Println("\n=== Constructor Pattern ===")
	newPerson := NewPerson("Eve", 28, "Medan")
	fmt.Printf("New person: %+v\n", newPerson)

	// Struct with methods for calculations
	fmt.Println("\n=== Struct with Calculations ===")
	rect := Rectangle{Width: 10, Height: 5}
	fmt.Printf("Rectangle: %+v\n", rect)
	fmt.Printf("Area: %.2f\n", rect.Area())
	fmt.Printf("Perimeter: %.2f\n", rect.Perimeter())
}

// Method dengan value receiver
func (p Person) Introduce() {
	fmt.Printf("Hi, I'm %s, %d years old from %s\n", p.Name, p.Age, p.City)
}

// Method yang memodifikasi struct (tetapi tidak permanen dengan value receiver)
func (p Person) HaveBirthday() {
	p.Age++ // ini tidak akan mengubah original struct
	fmt.Printf("Happy birthday! New age: %d (not permanent)\n", p.Age)
}

// Method dengan pointer receiver (bisa modifikasi struct)
func (p *Person) UpdateCity(newCity string) {
	p.City = newCity
	fmt.Printf("City updated to: %s\n", p.City)
}

// Constructor function
func NewPerson(name string, age int, city string) *Person {
	return &Person{
		Name: name,
		Age:  age,
		City: city,
	}
}

// Struct untuk contoh calculation
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}
