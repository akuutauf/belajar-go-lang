package main

import "fmt"

func main() {
	var name string

	name = "Taufik Hidayat"
	fmt.Println(name)

	name = "Leon Kennedy"
	fmt.Println("Nama yang baru :", name)

	/* jika membuat variabel lalu mengisikan secara langsung,
	 maka tidak perlu mendefinisikan tipe data variabel
	*/
	var hobby = "Membaca"
	fmt.Println("Hobi Saya adalah :", hobby)

	var age uint8 = 23
	fmt.Println("Usia saya adalah :", age)

	/*
	penggunaan var tidak wajib, atau bisa digantikan :=
	pada saat variabel langsung dilakukan pengisian nilai
	:= simbol ini digunakan hanya untuk deklarasi awal
	*/
	lastSchool := "Politeknik Negeri Banyuwangi"
	fmt.Println("Pendidikan Terakhir :", lastSchool)

	/*
	go juga dapat menyimpan banyak variabel untuk dalam penggunaan var
	*/

	var (
		parrentFisrtName = "Abu"
		parrentLastName = "Hasan"
	)

	fmt.Println("Nama Wali :", parrentFisrtName, parrentLastName)
}