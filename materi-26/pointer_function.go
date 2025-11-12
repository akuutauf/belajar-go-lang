package main

import "fmt"

// pointer function
// function bisa mengirimkan parameter, dimana sifatnya adalah pass by value (copy)
// sehingga ketika mengubah data diluar function, maka data tersebut tidak bisa berubah
// karena data yang dimasukkan sebelumnya disalin sehingga data yang original tidak berpengaruh
// akan tetapi ketika parameter menerapkan simbol '*', maka ketika diubah, data original juga berubah

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	// membuat object address
	// address := Address{} // object normal

	// namun jika kita menggunakan object pointer seperti yang diharapkan di function,
	// maka kita harus mendefinisikan object pointer pula
	var address1 *Address = &Address{}
	address2 := Address{}

	// memanggil function untuk mengubah country
	ChangeCountryToIndonesia(address1)

	// tapi kalau sudah terlanjut membuat object ke bentuk tanpa pointer,
	// maka cukup tambahkan '&'
	ChangeCountryToIndonesia(&address2) 

	// menampilkannya
	fmt.Println(address1)
	fmt.Println(address2)
}