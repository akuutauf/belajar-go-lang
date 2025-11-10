package main

import "fmt"

// recover (menggunakan kode sebelumnya yaitu panic)
// digunakan untuk merekover dari operasi panic
// dimana ketika program terhenti karena panic, disarankan menggunakan recover untuk tetap melanjutkan program
// agar program tidak berhenti begitu saja
// recover menangkap isi dari panic(...), dan melanjutkan program agar tetap berjalan serta mencatat error ke dalam log

func endApp() {
	fmt.Println("End app")

	// contoh penerapan recover yang benar
	// karena di fungsi utama diterapkan operasi defer,-
	// maka ketika terjadi error di fungsi utama dengan penerapan panic. 
	// kita bisa recover di fungsi defer dengan menangkap isi panic dan mencatat ke log
	// juga tetap melanjutkan program
	message := recover()
	fmt.Println("Error Message: ", message)
}

func runApp(error bool) {
	// mengimplementasikan endApp sebagai defer
	defer endApp()

	if error {
		panic("ERROR...")

		// contoh penerapan recover yang salah
		// message := recover()
		// fmt.Println("Error Message: ", message)
		
		// usaha recover diatas salah, dan tidak bisa mencatat dan melanjutkan program-
		// karena ketika panic sudah dieksekusi, maka seluruh program akan berhenti total
		// sehingga perlu diletakkan ke dalam function defer/defer
	} else {
		fmt.Println("Application is running")
	}
}

func main() {
	// memanggil fungsi run app
	runApp(true) // jika true, maka akan terjadi error

	// program berikutnya
	fmt.Println("Program berlanjut...")
}