# Golang Basic - Penjelasan Dasar Golang

Repository ini berisi penjelasan dan contoh-contoh dasar pemrograman Golang (Go).

## Daftar Isi

1. [Pengenalan Golang](#pengenalan-golang)
2. [Instalasi](#instalasi)
3. [Konsep Dasar](#konsep-dasar)
4. [Tipe Data](#tipe-data)
5. [Variabel dan Konstanta](#variabel-dan-konstanta)
6. [Operator](#operator)
7. [Kontrol Alur](#kontrol-alur)
8. [Fungsi](#fungsi)
9. [Array dan Slice](#array-dan-slice)
10. [Map](#map)
11. [Struct](#struct)
12. [Interface](#interface)
13. [Goroutine dan Channel](#goroutine-dan-channel)
14. [Error Handling](#error-handling)

## Pengenalan Golang

Go (atau Golang) adalah bahasa pemrograman yang dikembangkan oleh Google pada tahun 2007 dan dirilis sebagai open source pada tahun 2009. Go dirancang untuk:

- **Sederhana dan mudah dipelajari**: Sintaks yang bersih dan minimal
- **Cepat**: Kompilasi yang cepat dan performa yang tinggi
- **Concurrent**: Dukungan bawaan untuk pemrograman konkuren dengan goroutine
- **Scalable**: Cocok untuk membangun aplikasi yang scalable
- **Produktif**: Tools yang lengkap dan standard library yang kuat

## Instalasi

Untuk menginstal Golang, kunjungi [https://golang.org/dl/](https://golang.org/dl/) dan ikuti instruksi untuk sistem operasi Anda.

Verifikasi instalasi:
```bash
go version
```

## Konsep Dasar

### Hello World

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

Untuk menjalankan:
```bash
go run main.go
```

## Tipe Data

Go memiliki beberapa tipe data dasar:

### Tipe Data Numerik

- **Integer**: `int`, `int8`, `int16`, `int32`, `int64`
- **Unsigned Integer**: `uint`, `uint8`, `uint16`, `uint32`, `uint64`
- **Float**: `float32`, `float64`
- **Complex**: `complex64`, `complex128`

### Tipe Data Lainnya

- **Boolean**: `bool` (true/false)
- **String**: `string`
- **Byte**: `byte` (alias untuk uint8)
- **Rune**: `rune` (alias untuk int32, merepresentasikan Unicode code point)

Lihat contoh: [examples/01_datatypes.go](examples/01_datatypes.go)

## Variabel dan Konstanta

### Deklarasi Variabel

Ada beberapa cara mendeklarasikan variabel:

```go
// Cara 1: Dengan tipe data eksplisit
var name string = "John"

// Cara 2: Type inference
var age = 25

// Cara 3: Short declaration (hanya di dalam fungsi)
city := "Jakarta"

// Cara 4: Multiple variables
var x, y int = 1, 2
```

### Konstanta

```go
const Pi = 3.14159
const (
    StatusOK = 200
    StatusNotFound = 404
)
```

Lihat contoh: [examples/02_variables.go](examples/02_variables.go)

## Operator

### Operator Aritmatika
- `+` (penjumlahan)
- `-` (pengurangan)
- `*` (perkalian)
- `/` (pembagian)
- `%` (modulus)

### Operator Perbandingan
- `==` (sama dengan)
- `!=` (tidak sama dengan)
- `<` (lebih kecil)
- `>` (lebih besar)
- `<=` (lebih kecil atau sama dengan)
- `>=` (lebih besar atau sama dengan)

### Operator Logika
- `&&` (AND)
- `||` (OR)
- `!` (NOT)

Lihat contoh: [examples/03_operators.go](examples/03_operators.go)

## Kontrol Alur

### If-Else

```go
if age >= 18 {
    fmt.Println("Dewasa")
} else {
    fmt.Println("Anak-anak")
}

// If dengan statement pendek
if score := 85; score >= 80 {
    fmt.Println("Lulus dengan baik")
}
```

### Switch

```go
switch day {
case "Senin":
    fmt.Println("Awal minggu")
case "Jumat":
    fmt.Println("Akhir minggu kerja")
default:
    fmt.Println("Hari biasa")
}
```

### For Loop

```go
// For loop standar
for i := 0; i < 5; i++ {
    fmt.Println(i)
}

// While-style loop
for count < 10 {
    count++
}

// Infinite loop
for {
    // loop tanpa henti
    break // gunakan break untuk keluar
}

// Range loop
numbers := []int{1, 2, 3, 4, 5}
for index, value := range numbers {
    fmt.Printf("Index: %d, Value: %d\n", index, value)
}
```

Lihat contoh: [examples/04_controlflow.go](examples/04_controlflow.go)

## Fungsi

```go
// Fungsi sederhana
func greet(name string) {
    fmt.Printf("Hello, %s!\n", name)
}

// Fungsi dengan return value
func add(a int, b int) int {
    return a + b
}

// Multiple return values
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}

// Named return values
func calculate(a, b int) (sum int, product int) {
    sum = a + b
    product = a * b
    return // naked return
}

// Variadic function
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}
```

Lihat contoh: [examples/05_functions.go](examples/05_functions.go)

## Array dan Slice

### Array

Array memiliki ukuran tetap:

```go
var arr [5]int
arr[0] = 1
arr[1] = 2

// Inisialisasi array
numbers := [5]int{1, 2, 3, 4, 5}
```

### Slice

Slice memiliki ukuran dinamis:

```go
// Membuat slice
slice := []int{1, 2, 3}

// Append ke slice
slice = append(slice, 4, 5)

// Slicing
subset := slice[1:3] // elemen index 1 hingga 2

// Make slice dengan kapasitas
s := make([]int, 5, 10) // length 5, capacity 10
```

Lihat contoh: [examples/06_arrays_slices.go](examples/06_arrays_slices.go)

## Map

Map adalah struktur data key-value:

```go
// Membuat map
ages := make(map[string]int)
ages["Alice"] = 25
ages["Bob"] = 30

// Map literal
scores := map[string]int{
    "Math": 90,
    "Physics": 85,
}

// Mengakses value
age := ages["Alice"]

// Check key existence
value, exists := ages["Charlie"]
if exists {
    fmt.Println(value)
}

// Delete key
delete(ages, "Bob")

// Iterate map
for key, value := range scores {
    fmt.Printf("%s: %d\n", key, value)
}
```

Lihat contoh: [examples/07_maps.go](examples/07_maps.go)

## Struct

Struct adalah tipe data komposit yang mengelompokkan field:

```go
type Person struct {
    Name string
    Age  int
    City string
}

// Membuat instance
p1 := Person{
    Name: "Alice",
    Age:  25,
    City: "Jakarta",
}

// Akses field
fmt.Println(p1.Name)

// Pointer ke struct
p2 := &Person{Name: "Bob", Age: 30}
p2.Age = 31 // otomatis dereference

// Method pada struct
func (p Person) Introduce() {
    fmt.Printf("Hi, I'm %s from %s\n", p.Name, p.City)
}

// Method dengan pointer receiver
func (p *Person) Birthday() {
    p.Age++
}
```

Lihat contoh: [examples/08_structs.go](examples/08_structs.go)

## Interface

Interface mendefinisikan set method:

```go
type Shape interface {
    Area() float64
    Perimeter() float64
}

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

// Rectangle mengimplementasikan Shape interface
func printShapeInfo(s Shape) {
    fmt.Printf("Area: %.2f\n", s.Area())
    fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
}
```

Lihat contoh: [examples/09_interfaces.go](examples/09_interfaces.go)

## Goroutine dan Channel

### Goroutine

Goroutine adalah lightweight thread:

```go
func sayHello() {
    fmt.Println("Hello from goroutine!")
}

// Menjalankan goroutine
go sayHello()

// Anonymous goroutine
go func() {
    fmt.Println("Anonymous goroutine")
}()
```

### Channel

Channel digunakan untuk komunikasi antar goroutine:

```go
// Membuat channel
ch := make(chan int)

// Mengirim data ke channel
go func() {
    ch <- 42
}()

// Menerima data dari channel
value := <-ch

// Buffered channel
buffered := make(chan int, 5)

// Channel dengan range
for value := range ch {
    fmt.Println(value)
}

// Select untuk multiple channel
select {
case msg1 := <-ch1:
    fmt.Println(msg1)
case msg2 := <-ch2:
    fmt.Println(msg2)
default:
    fmt.Println("No message")
}
```

Lihat contoh: [examples/10_concurrency.go](examples/10_concurrency.go)

## Error Handling

Go menggunakan explicit error handling:

```go
import "errors"

// Membuat error
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// Menggunakan error
result, err := divide(10, 0)
if err != nil {
    fmt.Println("Error:", err)
    return
}
fmt.Println("Result:", result)

// Custom error type
type ValidationError struct {
    Field string
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// Defer, panic, recover
func safeDivide(a, b int) (result int) {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
            result = 0
        }
    }()
    
    return a / b // akan panic jika b = 0
}
```

Lihat contoh: [examples/11_errorhandling.go](examples/11_errorhandling.go)

## Struktur Project

```
golang-basic/
├── README.md
└── examples/
    ├── 01_datatypes.go
    ├── 02_variables.go
    ├── 03_operators.go
    ├── 04_controlflow.go
    ├── 05_functions.go
    ├── 06_arrays_slices.go
    ├── 07_maps.go
    ├── 08_structs.go
    ├── 09_interfaces.go
    ├── 10_concurrency.go
    └── 11_errorhandling.go
```

## Menjalankan Contoh

Untuk menjalankan contoh-contoh:

```bash
# Menjalankan satu file
go run examples/01_datatypes.go

# Build executable
go build examples/01_datatypes.go

# Run semua test (jika ada)
go test ./...
```

## Resources Tambahan

- [Official Go Documentation](https://golang.org/doc/)
- [Go by Example](https://gobyexample.com/)
- [A Tour of Go](https://tour.golang.org/)
- [Effective Go](https://golang.org/doc/effective_go)

## Lisensi

Repository ini dibuat untuk tujuan pembelajaran.

---

**Selamat belajar Golang! 🚀**