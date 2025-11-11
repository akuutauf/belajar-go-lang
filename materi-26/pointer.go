package main

import "fmt"

// pointer
// go lang bekerja secara pass by value (copy)
// pointer digunakan untuk menyimpan data berdasarkan refference (satu data yang sama dalam satu memori yang sama) -
// tanpa menduplikasi data
// pointer ditujukan untuk pass by refference
// untuk membuat variabel menggunakan referensi memori yang sama pada variabel lainnya, bisa mengunakan simbol '&'

type Address struct {
	City, Province, Country string
}

func main() {
	// membuat data dengan pass by value
	// membuat object dari struct address
	// address1 := Address{"Banyuwangi", "Jawa Timur", "Indonesia"}
	// address2 := address1 // copy value dari address1 ke address2

	// // mengubah data milik address2
	// address2.City = "Jember"

	// // menampilkan isi data dari masing masing object address
	// fmt.Println(address1) // data tidak berubah
	// fmt.Println(address2) // data berubah

	// membuat data dengan pointer: pass by refference
		// membuat object dari struct address
	var address1 Address = Address{"Banyuwangi", "Jawa Timur", "Indonesia"}
	var address2 *Address = &address1 // pointer ke address 1: pass by refference

	// mengubah data milik address2
	address2.City = "Jember"

	// menampilkan isi data dari masing masing object address
	fmt.Println(address1) // data tidak berubah
	fmt.Println(address2) // data berubah
}