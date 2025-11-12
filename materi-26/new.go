package main

import "fmt"

// new
// digunakan sebagai pengganti kata kunci '&' pada saat membuat object dari struct
// operator new biasanya memberikan data yang kosong pada saat object dibuat

type Address struct {
	City, Province, Country string
}

func main() {
	// membuat object dari struct address
	// var address1 *Address = &Address{} // kalau kita membuat seperti ini hasil datanya akan kosong
	var alamat1 *Address = new(Address) // ini lebih umum digunakan dari pada yang diatas
	var alamat2 *Address = alamat1 // alamat 1 dan 2 menuju pointer ke struct Address

	// mengubah data alamat 2
	alamat2.City = "Banyuwangi"

	// menampilkan hasil
	fmt.Println(alamat1) // data di alamat1 akan berubah, karena pointer nya sama
	fmt.Println(alamat2)
}