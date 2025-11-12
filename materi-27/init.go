package main

import (
	"belajar-go-lang/database"

	// jika kita ingin memanggil function init pada package internal tanpa memanggil fungsi di dalamnya,-
	// maka gunakan tanda blank identifyer atau simbol '_'
	_"belajar-go-lang/internal"
	"fmt"
)

func main() {
	fmt.Println(database.GetDatabase())
}