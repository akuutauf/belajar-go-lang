package main

import "fmt"

// closure
// adalah suatu metode yang memungkinkan function dapat mengakses data data disekitarnya (diatasnya)
// closure mampu dilakukan, dan menguntung untuk proses debug
// namun perlu diperhatikan hal ini jangan terlalu sering digunakan pada pengembangan aplikasi

func main() {
	counter := 0

	// membuat fungsi anonymous yang di simpan ke dalam variabel
	increment := func() {
		fmt.Println("Counter Increment")

		// operasi increment dapat dilakukan, namun sangat berbahaya jika fungsi terlalu kompleks
		// dimana memanipulasi data diluar fungsi, disarankan menggunakan parameter
		counter++
	}

	increment()
	increment()

	// atau pemanggilannya bisa dilakukan seperti ini, agar tidak membingungkan
	// counter = increment()

	// atau menggunakan parameter pada fungsi nya
	fmt.Println(counter)
}