package main

import "fmt"

func main() {
	// contoh perulangan biasa
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke ", counter)
	// 	counter++
	// }

	// contoh perulangan dengan menggunakan statement
	// init statement adalah 'counter := 1'
	// post statement adalah 'counter++'

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke ", counter)
	}
	
	fmt.Println("Selesai")

	// contoh perulangan manual dengan array
	names := []string{"Taufik", "Ilham", "Dimas"}

	for i := 0; i<len(names);i++ {
		fmt.Println(names[i])
	}

	// contoh for dengan range (foreach)
	for index, name:= range names {
		fmt.Println("index : ", index, "nama : ", name)
	}

	// contoh perulangan dengan range tanpa index
	for _, name:= range names {
		fmt.Println(name)
	}
}