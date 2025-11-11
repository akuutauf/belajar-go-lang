package main

import "fmt"

// struct method
// merupakan suatu function yang di tanamkan pada struct (class)
// ketika ingin membuat struct method, kita harus membuat struct terlebih dahulu
// begitu juga ketika ingin menjalankan method pada struct, perlu dibuat object terlebih dahulu
// (kode di bawah ini adalah lanjutkan dari kode sebelumnya)

// contoh pembuatan struct
type Customer struct {
	Name, Address string
	Age int
}

// contoh struct method
func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
}

func main() {
	// pembuatan data struct
	var taufik Customer
	fmt.Println(taufik) // jika data atribut belum di-isi, maka akan menampilkan default value

	// mengubah / mengisi nilai atribut struct
	taufik.Name = "Taufik Hidayat"
	taufik.Address = "Banyuwangi"
	taufik.Age = 25

	fmt.Println(taufik)
	fmt.Println(taufik.Name)
	fmt.Println(taufik.Address)
	fmt.Println(taufik.Age)

	// cara ke 2: membuat data dari struct
	dimas := Customer{
		Name: "Dimas",
		Address: "Indonesia",
		Age: 15,
	}

	fmt.Println(dimas)

	// cara ke 3: membuat data dari struct
	// cara ini harus urut sesuai dengan urutan kolom pada struct
	ilham := Customer{"Ilham", "Songgon", 18}
	fmt.Println(ilham)

	// pemanggilan method dari object yang sudah dibuat dari struct (class)
	taufik.sayHello("Ilham")
	dimas.sayHello("Ilham")
	ilham.sayHello("dimas")
}