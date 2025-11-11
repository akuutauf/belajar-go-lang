package main

import "fmt"

// interface
// digunakan sebagai kontrak pada suatu struct, agar dapat digunakan sesuai aturan dari interface
// interface dapat digunakan lebih dari satu struct

// contoh pembuatan interface
type HasName interface {
	 // dimana semua struct yang ingin mengimplementasikan interface HasName harus mempunyai method GetName()
	GetName() string
}

// contoh membuat method yang mengambil atribut 'name' dari struct yang mengimplementasikan interface
func sayHello (value HasName) {
	fmt.Println("Haii", value.GetName())
}

// kemudian setelah membuat interface, aturan method, dan method yang mengimplementasikan aturan interface
// selanjutnya membuat struct
type Person struct {
	Name string
}

// struct 2
type Animal struct {
	Name string
}

// kemudian kita bisa membuat method milik struct yang menerapkan aturan interface
// method GetName() sifatnya wajib ketika struct ingin mengimplementasikan dari interface HasName
func (person Person) GetName() string {
	return person.Name
}

// struct method 2: milik animal
func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	// pembuatan object struct
	person := Person{"Taufik H"}
	
	// method sayHello adalah method global yang mengimplementasikan interface
	sayHello(person) 

	// pembuatan object untuk struct 2: Animal
	animal := Animal{"Penguine"}
	sayHello(animal) 
}