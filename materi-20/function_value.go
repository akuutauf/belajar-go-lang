package main

import "fmt"

// function as value
// menyimpan fungsi ke dalam sebuah variabel

func getGoodBye(name string) string {
	return "Good Bye " + name
}


func main() {
	// menyimpan function ke dalam variabel
	// jadi ketika hendak menyimpan fungsi ke dalam variabel, maka jangan menggunakan tanda '()'
	contoh := getGoodBye

	fmt.Println(contoh("Taufik"))
	fmt.Println(contoh("Hidayat"))
}