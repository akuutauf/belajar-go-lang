package main

import "fmt"

// panic
// digunakan untuk menghentikan program, bilamana terjadi eror
// jadi program benar benar berhenti ketika kita menjalankan panic

func endApp() {
	fmt.Println("End app")
}

func runApp(error bool) {
	// mengimplementasikan endApp sebagai defer
	defer endApp()

	if error {
		panic("ERROR...")
	} else {
		fmt.Println("Application is running")
	}
}

func main() {
	// memanggil fungsi run app
	// runApp(false)
	runApp(true) // jika true, maka akan terjadi error
}