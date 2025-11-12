package main

import (
	"belajar-go-lang/helper"
	"fmt"
)

// import
// di go lang, jika kita ingin menggunakan kode dalam package yang berbeda bisa menggunakan import

func main() {

	// pemanggilan biasa
	fmt.Println("Taufik")

	// pemanggilan dengan fungsi yang sudah di import dari helper
	result := helper.SayHello("Taufik")
	fmt.Println(result)

	// menampilkan reource lainnya
	// result2 := helper.sayGoodBye("Taufik") // tidak bisa diakses
	// fmt.Println(helper.version) // tidak bisa diakses

	fmt.Println(helper.Application)
}