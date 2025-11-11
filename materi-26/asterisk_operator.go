package main

import "fmt"

// asterisk operator
// digunakan untuk menandai variabel yang bersifat pointer
// jika variabel sudah di buat menjadi bentuk pointer (menuju ke refference variabel lain yang sama)-
// maka, ketika di inisiasi ulang untuk datanya harus di definisikan sebagai pointer juga dengan tanda '&'
// jika di inisiasi dengan tanda '&' pada STRUCT maka akan membuat refference baru,
// namun jika di inisiasi dengan tanda '*' pada variabel nya, maka seluruh variabel yang mengacu refference yang sama,-
// akan berpindah ke refference yang baru saja dibuat

type Address struct {
	City, Province, Country string
}

func main() {
	// membuat object dari struct address
	var address1 Address = Address{"Banyuwangi", "Jawa Timur", "Indonesia"}
	var address2 *Address = &address1 // address2: adalah sebagai pointer

	// mengubah data milik address2
	address2.City = "Jember" // data address1 juga ikut berubah

	// menampilkan isi data dari masing masing object address
	fmt.Println(address1) // ikut berubah
	fmt.Println(address2) // data berubah

	// namun jika kita ingin mengubah pointer milik variabel address 2,
	// maka definisikan ulang pointer nya dengan tanda '&'
	// akan mengubah data address nya (address2 punya address sendiri dari address sebelumnya)
	// address2 = &Address{"Malang", "Jawa Timur", "Indonesia"}
	// fmt.Println(address1)
	// fmt.Println(address2) // data baru

	// memindah seluruh refference dari address lama ke address yang baru dengann tanda '*'
	// ketika kode ini dijalankan, maka seluruh variabel yang mengacu ke address ini akan beralih ke address yang baru
	*address2 = Address{"Kediri", "Jawa Timur", "Indonesia"}
	fmt.Println(address1)
	fmt.Println(address2) // semua variabel yang mengacu ke data sebelumnya berpincah acuan ke data yang baru ini
}