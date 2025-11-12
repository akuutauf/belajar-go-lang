package main

import "fmt"

// pointer method
// method pada struct bisa diaktifkan pointer dengan menggunakan '*' pada parameternya
// direkomendasikan selalu gunakan pointer pada method
// pada pointer method, object tidak perlu didefinisikan sebagai pointer juga-
// cukup tanda bintang '*' pada parameter method, maka semua data yang dikirimkan ke method otomatis di atur ke pointer

type Man struct {
	Name string
}

// method
func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	// pembuatan object
	taufik  := Man{"Taufik"}
	taufik.Married()

	fmt.Println(taufik.Name)

}