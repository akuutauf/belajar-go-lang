package main

import "fmt"

func main() {
	var nilai32 int32 = 32769
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	var name = "Taufik Hidayat"
	var t uint8 = name[0] // mendapatkan nilai 't' dalam bentuk byte (uint8)
	var tString = string(t)

	fmt.Println(name)
	fmt.Println(t)
	fmt.Println(tString)

}