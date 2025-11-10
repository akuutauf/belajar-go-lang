package main

import "fmt"

// function as parameter
// digunakan untuk mengambil parameter dari fungsi lain yang tipe datanya serupa/sesuai

// type declaration: untuk function
// Filter: nama lain dari function nya
type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	// menyimpan nama yang sudah di filter pada fungsi yang lain, melalui parameter function yaitu filter
	// memanggil fungsi yang lain melalui parameter filter
	filterredName := filter(name)
	
	// menampilkan teks
	fmt.Println("Hello " + filterredName)
}

// membuat function yang lain sebagai acuan parameter pada fungsi utama (sayHelloWithFilter)
func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	// memanggil fungsi sayHelloWithFilter dan mengisikan function spamFilter sebagai parameternya
	// jangan gunakan tanda '()' pada spamFilter agar tidak mengembalikan nilai (hanya butuh fungsinya saja)
	sayHelloWithFilter("Taufik", spamFilter)

	// meyinpan fungsi ke dalam variabel filter (function as value)
	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)
}