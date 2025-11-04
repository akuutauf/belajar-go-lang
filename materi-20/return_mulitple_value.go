package main

import "fmt"

// return multiple value
// bisa diberikan dengan memberikan tipe data return yang lebih dari satu
// dan wajib memberikan return
func getFullname() (string, string) {
	return "Taufik", "Hidayat"
}

func main() {
	// pemanggilan return multiple value
	// pemanggilan di simpan ke dalam variabel
	firstName, lastName := getFullname()

	fmt.Println(firstName, lastName)

	// pemanggilan return mutiple value mewajibkan semua nilai diambil
	// namun bisa dilakukan pengabaian untuk data yang tidak kita butuhkan, seperti dibawah ini
	namaDepan, _ := getFullname()

	fmt.Println(namaDepan)
}