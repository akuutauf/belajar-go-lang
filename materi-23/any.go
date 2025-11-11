package main

import "fmt"

// interface kosong
// dapat digunkana untuk mengisi apapun dan mereturn tipe data apapun
// seperti string, int, struct, maupun array
// interface kosong juga bisa di definisikan dengan keywoard 'any'
// function println, panic menggunakan interface kosong/any
// tipe data yang paling tinggi, semuanya mengikuti interface kosong/any

func Ups() any {
	// return 1
	// return "Hai"
	return true
}

func main() {
	// contoh pemanggilan interface kosong
	var kosong any = Ups()
	fmt.Println(kosong)
}