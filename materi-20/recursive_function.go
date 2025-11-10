package main

import "fmt"

// recursive function
// digunakan untuk memanggil fungsi nya sendiri
// sebagai contoh studi kasus: factorial

// fungsi factorial (tanpa recursive)
func factorialLoop(value int) int {
	// memberikan nilai awal
	result := 1
	
	// perulangan faktorial
	for i := value; i > 0; i-- {
		result *= i
	}

	return result
}

// fungsi factorial (dengan recursive)
func factorialRecursive(value int) int {
	// melakukan pengecekan, ketika value sudah mencapai 1, makan recursive dihentikan
	if value == 1 {
		return 1 
	} else {
		// jika value tidak mencapai nilai selain 1, maka lanjutkan
		// proses rekursiv dilakukan dengan mengkalikan nilai saat ini, dengan nilai dibawahnya (value - 1)
		return value * factorialRecursive(value - 1)
	}
}

func main() {
	// pemanggilan fungsi faktorial (tanpa recursive)
	fmt.Println(factorialLoop(10))

	// pemanggilan fungsi faktorial (dengan recursive)
	fmt.Println(factorialRecursive(10))
}