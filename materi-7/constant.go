package main

import "fmt"

func main() {
	const firstName = "Taufik"
	const lastName = "Hidayat"
	// fullName := "Taufik Hidayat" : adalah bentuk var, const tidak memiliki instance (:=)

	// deklarasi multiple const
	// boleh untuk tidak dipakai, jikalau dalam bentuk variabel constant
	const (
		hobby = "Membaca"
		negara = "Indonesia"
		umur = 24 
	)

	// akan terjadi error, ketika variabel berusaha untuk diubah kembali
	// firstName = "Leon"
	// lastName = "Khenndedy"

	// const : diperbolehkan untuk tidak dipanggil di dalam program, tidak seperti var
	fmt.Println(firstName, lastName)
	fmt.Print(hobby, negara)
}
