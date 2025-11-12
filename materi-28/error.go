package main

import (
	"errors"
	"fmt"
)

// error
// dalam go lang package error bisa digunakan untuk alternatif tindak lanjut kesalahan dari program yang sudah kita buat

// membuat function pembagian
// di dalam function dibawah ini, terdapat 2 return, pertama value (int) dan juga kemungkinan error (interface error)
// berhubung 2 return, maka jika tidak terjadi error, gunakan 'nil'
func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Kesalahan, Pembagian dengan Nol") // memanggil fungsi milik package error
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	result, err := Pembagian(100, 0)

	// melakukan pengecekan
	if err == nil {
		fmt.Println("Hasil =", result) // jika berhasil
	} else {
		fmt.Println("Error:", err.Error()) // jika error, memanggil fungsi error
	}
}