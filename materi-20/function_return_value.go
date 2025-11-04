package main

import "fmt"

// pembuatan function return value
// function return value mewajibkan untuk memberikan tipe data return (contoh: string)
// jika tipe data return sudah diberikan, maka function wajib terdapat return
func getHello(name string) string {
	hello := "Hello " + name
	return hello
}

func main() {
	// pemanggilan function return value
	// jika function return value dipanggil, maka dibutuhkan variabel menyimpannya
	result := getHello("Taufik")
	fmt.Println(result)

	// atau bisa diberikan langsung pada fungsi print seperti berikut ini
	fmt.Println("Ilham")
	fmt.Println("Dimas")
}