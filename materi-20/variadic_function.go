package main

import "fmt"

// variadic function
// biasanya menggunakan parameter (args) berupa array/slice
func sumAll(numbers ...int) int {
	total := 0

	// menggunakan perulangan (for each)
	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	// jika menggunakan variadic function, cukup mengisikan nilai parameter nya secara langsung
	fmt.Println(sumAll(10,10,10,10,10))
	fmt.Println(sumAll(10,10,10,10,10,10,10))

	// namun jika function menggunakan parameter array, maka pemanggilan nya seperti dibawah ini
	// fmt.Println(sumAll([]int{10,10,10,10,10,10,10,10,10,10}))

	// namun, apabila kita sudah mempunyai variabel slice
	// kita juga bisa memanggil variadic function, dengan menambahkan '...'
	numbers := []int{20,20,20,20}
	fmt.Println(sumAll(numbers...))
}