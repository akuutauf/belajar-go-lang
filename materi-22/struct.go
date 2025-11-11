package main

import "fmt"

// struct
// adalah suatu cara yang digunakan untuk membuat template yang berfungsi sebagai acuan suatu data
// struct bisa di buat dengan menggunakan keywoard 'type' dan 'struct'
// struct terdiri dari beberapa atribut yang tipe datanya berbeda beda

// contoh pembuatan struct
type Customer struct {
	Name, Address string
	Age int
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
}