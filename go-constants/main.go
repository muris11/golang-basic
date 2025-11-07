package main

import "fmt"

const LENGTH int = 10 // deklarasi konstanta dengan tipe data eksplisit
const WIDTH = 5 // deklarasi konstanta dengan tipe data implisit

func main() {

	const PI = 3.14 // deklarasi konstanta lokal dalam fungsi main
	// PI = 20 // akan menyebabkan error karena konstanta tidak dapat diubah nilai nya

	const ( 
		X int = 11 // deklarasi konstanta dalam grup
		Y 	= 3.14 // deklarasi konstanta dalam grup dengan tipe data implisit
		Z 	= "Hello World" // deklarasi konstanta dalam grup dengan tipe data string
	)

	fmt.Println(LENGTH) // penggunaan konstanta
	fmt.Println(WIDTH)  // penggunaan konstanta
	fmt.Println(PI) // penggunaan konstanta lokal
	fmt.Println(X) // penggunaan konstanta dalam grup
	fmt.Println(Y) // penggunaan konstanta dalam grup
	fmt.Println(Z) // penggunaan konstanta dalam grup
}
