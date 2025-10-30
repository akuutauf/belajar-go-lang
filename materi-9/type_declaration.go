package main

import "fmt"

func main() {
	type NoKTP string // membuat tipe data baru (NoKTP) yang berjenis string/alias string

	var ktpTaufik NoKTP = "1111111111111111" // NoKTP sebagai tipe data baru sebagai alias string
	var contoh string = "2222222222222222"

	var contohKtp = NoKTP(contoh) // tipe data baru juga bisa digunakan sebagai konversi

	fmt.Println(ktpTaufik)
	fmt.Println(contohKtp)
}