package main

import "fmt"

// anonymous function
// adalah suatu metode untuk membuat function tanpa nama
// anonymous function bisa di simpan ke dalam variabel, maupun di deklarasikan langsung tanpa nama
// teknik ini digunakan ketika terdapat studi kasus yang penerapannya sederhana

// type declaration
type Blacklist func(string) bool

// function utama
func registerUser(name string, blacklist Blacklist) {
	// melakukan pengecekan dari nilai return fugsi dari parameter blacklist
	if blacklist(name) {
		fmt.Println("Your are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	// anonymous function di simpan ke dalam variabel
	blacklist := func(name string) bool {
		return name == "anjing"
	}

	// pemanggilan fungsi utama
	registerUser("Hidayat", blacklist)

	// anonymous function di deklarasikan langsung pada saat pemanggilan fungsi utama
	registerUser("anjing", func(name string) bool {
		return name == "anjing"
	})
}