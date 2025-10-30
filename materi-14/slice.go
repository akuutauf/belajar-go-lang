package main

import "fmt"

func main() {
	names := [...]string{"Taufik", "Hidayat", "Ilham", "Dimas", "Risky", "Prayoga"}

	slice1 := names[4:6]
	fmt.Println(slice1)
	
	slice2 := names[:3]
	fmt.Println(slice2)
	
	slice3 := names[3:]
	fmt.Println(slice3)
	
	slice4 := names[:]
	fmt.Println(slice4)

	// append: menggabungkan slice untuk array lama ke array yang baru
	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	daySlice1 := days[5:]
	fmt.Println()
	fmt.Println(daySlice1)

	// mengubah nilai index dayslice1 akan berpengaruh ke array days
	// karena slice hanya berfungsi sebagai penunjuk ke array utama saja (days)
	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"

	fmt.Println(daySlice1)
	fmt.Println(days)

	// membuat dayslice2 dengan menggunakan fungsi append
	// ketika array utama (days) memiliki ruang untuk menambahkan array baru,
	// maka elemen array yang baru akan ditambahkan, namun jika sudah penuh, akan dibuatkan array baru
	// array baru di simpan di dalam logic dayslice2, namun tidak bisa diakses secara langsung
	
	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "Sabtu Lama"
	// daysBaru := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu Lama", "Minggu", "Libur Baru"}

	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)

	// membuat slice dari awal
	fmt.Println()
	newSlice := make([]string, 2, 5) // panjang array = 2, kapasitas array = 5
	newSlice[0] = "Taufik"
	newSlice[1] = "Hidayat"
	// newSlice[2] = "Ilham" // error, harusnya menggunakan append

	fmt.Println(newSlice)
	fmt.Println(len(newSlice)) // menghitung panjang newSlice (panjang maksimal nya adalah = 2)
	fmt.Println(cap(newSlice)) // menghitung kapasitas newSlice (kapasitas maksimal nya adalah = 5)
	// panjang array bisa ditentukan diawal, berbeda dengan kapasitas. karena kapasitas adalah batas keseluruhan array

	// menambahkan index baru pada slice/array dengan menggunakan append
	newSlice2 := append(newSlice, "Ilham", "Deny")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))
	
	newSlice2[0] = "Dimas" // ketika mengubah nilai indeks slice 2, maka nilai newSlice pertama akan berubah
	fmt.Println(newSlice) // karena masih menggunakan array yang sama
	fmt.Println(newSlice2)

	// copy slice
	fromSlice := days[:] // menampung array yang akan di copy
	daySliceCopy := make([]string, len(fromSlice), cap(fromSlice))
	// daySliceCopy := make([]string, len(days), cap(days)) // atau bisa pakai cara ini

	copy(daySliceCopy, fromSlice) // melakukan copy array

	fmt.Println()
	fmt.Println(fromSlice)
	fmt.Println(daySliceCopy)
	
	daySliceCopy[0] = "Wawan"
	fmt.Println(daySliceCopy) // array baru, jika di ubah tidak mempengaruhi array utama (days)
	fmt.Println(fromSlice) // array utama
	fmt.Println(days) // array utama

	// perbedaan pembuatan array dan slice
	iniArray := [...]int{1,2,3,4,5}
	iniSlice := []int{1,2,3,4,5}

	fmt.Println()
	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
