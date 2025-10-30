package main

import "fmt"

func main() {
	fmt.Println("Nama Saya Adalah :")
	fmt.Print("Taufik ")
	fmt.Println("Hidayat")

	// len : digunakan untuk menghitung jumlah karakter string
	fmt.Println("Jumlah Huruf pada = 'Taufik' adalah:", len("Taufik"))

	// "string"[4] : digunakan untuk mengambil abjad/huruf pada suatu string
	fmt.Println("Hidayat"[4]) // akan muncul dalam bentuk byte (y = 121)
}