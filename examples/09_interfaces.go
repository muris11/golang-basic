package main

import (
	"fmt"
	"math"
)

// Interface mendefinisikan kontrak (set of methods)
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Struct Rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// Rectangle mengimplementasikan Shape interface
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Struct Circle
type Circle struct {
	Radius float64
}

// Circle mengimplementasikan Shape interface
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Interface untuk animal
type Animal interface {
	Speak() string
	Move() string
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return "Woof!"
}

func (d Dog) Move() string {
	return "Running"
}

type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return "Meow!"
}

func (c Cat) Move() string {
	return "Walking"
}

// Empty interface
type Any interface{}

// Interface composition
type Reader interface {
	Read() string
}

type Writer interface {
	Write(data string)
}

// ReadWriter adalah komposisi dari Reader dan Writer
type ReadWriter interface {
	Reader
	Writer
}

type File struct {
	Content string
}

func (f File) Read() string {
	return f.Content
}

func (f *File) Write(data string) {
	f.Content = data
}

func main() {
	fmt.Println("=== Contoh Interface di Golang ===\n")

	// Menggunakan interface
	fmt.Println("=== Shape Interface ===")
	rect := Rectangle{Width: 10, Height: 5}
	circle := Circle{Radius: 7}

	printShapeInfo(rect)
	printShapeInfo(circle)

	// Slice of interface
	fmt.Println("\n=== Slice of Interface ===")
	shapes := []Shape{
		Rectangle{Width: 5, Height: 3},
		Circle{Radius: 4},
		Rectangle{Width: 8, Height: 6},
	}

	totalArea := 0.0
	for i, shape := range shapes {
		area := shape.Area()
		fmt.Printf("Shape %d - Area: %.2f, Perimeter: %.2f\n", i+1, area, shape.Perimeter())
		totalArea += area
	}
	fmt.Printf("Total area: %.2f\n", totalArea)

	// Animal interface
	fmt.Println("\n=== Animal Interface ===")
	dog := Dog{Name: "Buddy"}
	cat := Cat{Name: "Whiskers"}

	animals := []Animal{dog, cat}
	for _, animal := range animals {
		fmt.Printf("Animal speaks: %s, moves by %s\n", animal.Speak(), animal.Move())
	}

	// Type assertion
	fmt.Println("\n=== Type Assertion ===")
	var shape Shape = Rectangle{Width: 10, Height: 5}
	
	// Type assertion dengan check
	if rect, ok := shape.(Rectangle); ok {
		fmt.Printf("It's a rectangle with width: %.2f and height: %.2f\n", rect.Width, rect.Height)
	}

	// Type assertion tanpa check (bisa panic jika salah)
	rectangle := shape.(Rectangle)
	fmt.Printf("Rectangle: %+v\n", rectangle)

	// Type switch
	fmt.Println("\n=== Type Switch ===")
	describeShape(Rectangle{Width: 5, Height: 3})
	describeShape(Circle{Radius: 4})
	describeShape("This is a string")

	// Empty interface
	fmt.Println("\n=== Empty Interface ===")
	var anything interface{}
	
	anything = 42
	fmt.Printf("Value: %v, Type: %T\n", anything, anything)
	
	anything = "Hello"
	fmt.Printf("Value: %v, Type: %T\n", anything, anything)
	
	anything = true
	fmt.Printf("Value: %v, Type: %T\n", anything, anything)
	
	anything = Rectangle{Width: 10, Height: 5}
	fmt.Printf("Value: %v, Type: %T\n", anything, anything)

	// Function dengan empty interface parameter
	fmt.Println("\n=== Function with Empty Interface ===")
	printAny(42)
	printAny("Golang")
	printAny(3.14)
	printAny([]int{1, 2, 3})

	// Interface composition
	fmt.Println("\n=== Interface Composition ===")
	file := &File{Content: "Initial content"}
	fmt.Printf("Read: %s\n", file.Read())
	file.Write("New content")
	fmt.Printf("Read after write: %s\n", file.Read())

	// Interface values
	fmt.Println("\n=== Interface Values ===")
	var animal Animal
	fmt.Printf("Initial - Value: %v, Type: %T\n", animal, animal)
	
	animal = Dog{Name: "Rex"}
	fmt.Printf("After assignment - Value: %v, Type: %T\n", animal, animal)
	fmt.Printf("Speaks: %s\n", animal.Speak())

	// Nil interface vs interface with nil value
	fmt.Println("\n=== Nil Interface ===")
	var nilInterface Animal
	fmt.Printf("Nil interface: %v (is nil: %t)\n", nilInterface, nilInterface == nil)
	
	var nilDog *Dog
	var interfaceWithNil Animal = nilDog
	fmt.Printf("Interface with nil value: %v (is nil: %t)\n", interfaceWithNil, interfaceWithNil == nil)

	// Polymorphism
	fmt.Println("\n=== Polymorphism ===")
	processAnimals([]Animal{
		Dog{Name: "Max"},
		Cat{Name: "Luna"},
		Dog{Name: "Charlie"},
	})

	// Interface method with different receivers
	fmt.Println("\n=== Different Receivers ===")
	counter := &Counter{Value: 0}
	demonstrateCounter(counter)
}

// Function yang menerima interface
func printShapeInfo(s Shape) {
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

// Type switch
func describeShape(i interface{}) {
	switch v := i.(type) {
	case Rectangle:
		fmt.Printf("Rectangle with area: %.2f\n", v.Area())
	case Circle:
		fmt.Printf("Circle with area: %.2f\n", v.Area())
	case string:
		fmt.Printf("String: %s\n", v)
	case int:
		fmt.Printf("Integer: %d\n", v)
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}
}

// Function dengan empty interface
func printAny(value interface{}) {
	fmt.Printf("Value: %v, Type: %T\n", value, value)
}

// Polymorphism example
func processAnimals(animals []Animal) {
	for i, animal := range animals {
		fmt.Printf("Animal %d: %s - %s\n", i+1, animal.Speak(), animal.Move())
	}
}

// Counter untuk demonstrasi method dengan pointer receiver
type Counter struct {
	Value int
}

func (c *Counter) Increment() {
	c.Value++
}

func (c Counter) GetValue() int {
	return c.Value
}

type Incrementer interface {
	Increment()
}

func demonstrateCounter(inc Incrementer) {
	fmt.Println("Demonstrating counter:")
	for i := 0; i < 3; i++ {
		inc.Increment()
		// inc harus berupa pointer karena Increment() memiliki pointer receiver
		if counter, ok := inc.(*Counter); ok {
			fmt.Printf("  Counter value: %d\n", counter.Value)
		}
	}
}
