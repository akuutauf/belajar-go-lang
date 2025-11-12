package helper

import "fmt"

// package
// pakage dalam golang dapat disebut juga sebagai folder
// file di dalam package biasanya mengikuti nama package, seperti: file helper nama package nya adalah helper

// access modifier
// digunakan untuk membatasi akses untuk resource baik variabel function, maupun lain lain
// bisanya jika diawali dengan huruf kecil, hanya bisa diakses di package yang sama (helper)
// namun jika diawali dengan huruf besar, maka bisa diakses di package yang lain (main)

var version = "1.0.0" // hanya bisa diakses di package yang sama (helper)
var Application = "golang" // bisa diakses diluar package

func SayHello(name string) string {
	return "Hello " + name
}

// contoh function yang tidak bisa diakses diluar package, karena menggunakan awalan huruf kecil
func sayGoodBye(name string) string {
	return "Good Bye " + name
}

// function pemanggil untuk resource yang satu package
func Contoh(name string) {
	sayGoodBye(name)
	fmt.Println(version)
}