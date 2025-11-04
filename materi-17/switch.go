package main

import (
	"fmt"
)

func main() {
	name := "Dimas"

	switch name {

	case "Taufik":
		fmt.Println("Hai Taufik")

	case "Ilham":
		fmt.Println("Hai Ilham")

	default: 
		fmt.Println("Hai boleh kenalan?")
	}	

	// short statement switch
	// length := len(name) // yang biasanya digunakan

	switch length := len(name); length > 5 {
	case true:		
		fmt.Println("Nama terlalu panjang")
	
	case false:
		fmt.Println("Nama sudah benar")
	}

	// switch juga bisa tanpa parameter kondisi, sehingga kondisinya cukup diletakkan di setiap case
	// namun hal ini sangat tidak direkomendasikan, karena lebih baik menggukan if else expression
	length := len(name)

	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}