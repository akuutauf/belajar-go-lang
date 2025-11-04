package main

import "fmt"

// function dengan parameter
// function parameter berisikan nama variabel, dan tipe data yang digunakan
func sayHelloTo(firstName string, lastname string) {
	fmt.Println("Hello", firstName, lastname)
}

func main() {
	// pemanggilan function parameter
	// jika terdapat parameter maka wajib mengisi semua parameter 
	// isi parameter yang dikirimkan juga harus sesuai tipe data nya
	sayHelloTo("Taufik", "Hidayat")
	sayHelloTo("Ilham", "Nurizky")
}