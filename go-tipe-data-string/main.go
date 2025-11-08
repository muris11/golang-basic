package main

import "fmt"

func main(){
	var myName string = "Rifqy Saputra" // deklarasi variabel dengan tipe data string
	var myAddress string // deklarasi variabel dengan tipe data string
	nickName := "Rifqy" // deklarasi variabel dengan tipe data string menggunakan short declaration

	fmt.Printf("Value: %v, Tipe data: %T\n", myName, myName) // mencetak nilai dan tipe data variabel myName
	fmt.Printf("Value: %v, Tipe data: %T\n", myAddress, myAddress) // mencetak nilai dan tipe data variabel myAddress
	fmt.Printf("Value: %v, Tipe data: %T\n", nickName, nickName) // mencetak nilai dan tipe data variabel nickName	

}