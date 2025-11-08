package main

import "fmt"

func main() {

	var person = map[string]string{ // deklarasi map dengan tipe data string
		"name":    "Rifqy Saputra",
		"job": "Programmer",
	} 

	stock := map[string]int{ // deklarasi map dengan tipe data integer
		"apple":  100,
		"handprhone": 200,
	}

	fmt.Println(person) // mencetak nilai map person
	fmt.Println(person["name"]) // mencetak elemen dengan key "name" dari map person 
	fmt.Println(person["job"]) // mencetak elemen dengan key "job" dari map person

	fmt.Println(stock) // mencetak nilai map stock
	fmt.Println(stock["apple"]) // mencetak elemen dengan key "apple" dari map stock
	fmt.Println(stock["handprhone"]) // mencetak elemen dengan key "handprhone" dari map stock


	var laptop = make(map[string]string) // deklarasi map dengan tipe data string menggunakan fungsi make
	laptop["merk"] = "Acer"
	laptop["model"] = "Nitro 5"

	computer := make(map[string]int) // deklarasi map dengan tipe data integer menggunakan fungsi make
	computer["dell"] = 1500
	computer["lenovo"] = 2500

	fmt.Println(laptop) // mencetak nilai map laptop
	fmt.Println(laptop["merk"]) // mencetak elemen dengan key "merk" dari map laptop
	fmt.Println(laptop["model"]) // mencetak elemen dengan key "model" dari map laptop

	fmt.Println(computer) // mencetak nilai map computer	
	fmt.Println(computer["dell"]) // mencetak elemen dengan key "dell" dari map computer
	fmt.Println(computer["lenovo"]) // mencetak elemen dengan key "lenovo" dari map computer


	fmt.Println(len(person)) // mencetak panjang map person
	fmt.Println(len(stock)) // mencetak panjang map stock

	var mapKosong1 = make(map[string]string) // deklarasi map kosong dengan tipe data string menggunakan fungsi make
	var mapKosong2 map[string]string // deklarasi map kosong dengan tipe data integer menggunakan fungsi make

	fmt.Println(mapKosong1 == nil) // mencetak nilai mapKosong1
	fmt.Println(mapKosong2 == nil) // mencetak nilai mapKosong2

	laptop["tahun"] = "2020" // menambahkan elemen pada map laptop
	computer["asus"] = 3000 // menambahkan elemen pada map computer

	fmt.Println(laptop) // mencetak nilai map laptop
	fmt.Println(laptop["tahun"]) // mencetak nilai map laptop
	fmt.Println(computer) // mencetak nilai map computer
	fmt.Println(computer["asus"]) // mencetak nilai map computer

	delete(laptop, "model") // menghapus elemen dengan key "model" dari map laptop
	fmt.Println(laptop) // mencetak nilai map laptop

	person2 := person // membuat salinan map person ke person2
	fmt.Println(person) // mencetak nilai map person
	fmt.Println(person2)  // mencetak nilai map person2

	person2["name"] = "Rifqy S" // mengubah elemen dengan key "name" dari map person2

	fmt.Println(person) // mencetak nilai map person setelah diubah
	fmt.Println(person2) // mencetak nilai map person2 setelah diubah

}