package main

import "fmt"

func main() {
	name := "Dimas"

	if name == "Taufik" {
		fmt.Println("Halo Taufik")
	} else if name == "Ilham" {
		fmt.Println("Hai Ilham, apa kabar")
	} else {
		fmt.Println("Hai, boleh kenalan ga?", name)
	}

	// short statement if
	// length := len(name) // contoh yang biasanya digunakan

	// menggunakan short statement
	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah sesuai")
	}
}