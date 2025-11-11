package main

import "fmt"

// type assertions
// digunakan untuk mengkonversi isi dari variabel dari suatu tipe data ke tipe data tertentu
// namun perlu diperhatikan, untuk melakukan konversi suatu value semisal dari string ke int

// function random yang memiliki keluaran any atau interface kosong
func random() any {
	return true
}

func main() {
	// memanggil function random
	var result any = random()

	// melakukan konversi dengan type assetions
	// var resultString string = result.(string)

	// fmt.Println(resultString)

	// jikalau di konversi dari string ke int akan mengakibatkan panic, jika memang value nya tidak sesuai
	// var resultInt int = result.(int)

	// fmt.Println(resultInt)

	// namun agar tidak error, disarankan menggunakan switch expression (pencegahan)
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	case bool:
		fmt.Println("Boolean", value)
	default:
		fmt.Println("Unknown", value)
	}

}