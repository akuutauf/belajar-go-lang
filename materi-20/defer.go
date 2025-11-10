package main

import "fmt"

// defer
// merupakan suatu metode yang digunakan untuk menjadwalkan/menjalankan -
// suatu fungsi dalam fungsi lain setelah fungsi utama selesai dijalkankan.
// defer akan mengabaikan error pada fungsi utama, sehingga masih tetap dijalankan ketika terjadi error

func logging() {
	fmt.Println("Fungsi selesai dijalankan")
}

func runApplication() {
	// memanggil function logging sebagai defer
	defer logging()
	fmt.Println("Application is running...")
}

func main() {
	// memanggil fungsi utama
	runApplication()
}