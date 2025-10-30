package main

import "fmt"

func main() {
	var name [3]string
	name[0] = "Taufik"
	name[1] = "Ilham"
	name[2] = "Dimas"

	fmt.Println(name[0])
	fmt.Println(name[1])
	fmt.Println(name[2])

	// Array juga bisa di inisialisasi secara langsung
	// jika nilai index array tidak di isi, maka akan terdapat nilai default yaitu 0
	var values = [3]int {24,17,15,}
	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	// menghitung panjang array, meskipun array tidak disi, akan tetap bisa dihitung panjangnya
	fmt.Println(len(values))
	
	// array bisa di atur untuk jumlah yang tidak menentu menggunakan tanda '...'
	// namun operasi ini hanya boleh di panggildan langsung di isikan nilainya
	// jika tidak di isi maka akan 0, dan juga tidak bisa menambahkan data baru setelah dilakukan pengisian data diawal
	var number = [...]int {1,2,3,4,5}
	fmt.Println(number)
	fmt.Println(len(number))
	
	// mengubah isi array dapat menggunakan pemanggilan langsung dan operator assignment
	number[0] = 0
	fmt.Println(number)
}