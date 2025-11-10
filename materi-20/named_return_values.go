package main

import "fmt"

// named return values
// named return value dapat mengembalikan beberapa nilai
// bisa di definisikan parameter variabel dengan tipe data
// sehingga isi function tidak perlu mendeklarasikan parameter variabel ulang
// jika parameter fungsi memiliki tipe data yang sama, maka cukup diberikan tipe data diakhir parameter

func getCompleteName() (firstName, middleName, lastName string) {
	// tidak perlu mendefinisikan variabel baru dengan simbol ':='
	firstName = "Muhammad"
	middleName = "Ilham"
	// lastName = "Nurizky"

	// jika salah satu variabel tidak di isi, maka disi nilai default string ("")
	return firstName, middleName, lastName
}

func main() {
	a, b, c := getCompleteName()

	fmt.Println(a, b, c)
}