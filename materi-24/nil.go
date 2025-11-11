package main

import "fmt"

// nil
// nil adalah alternatif dari nilai yang kosong (null)
// nill digunakan pada beberapa tipe data seperti, interface, function, map, slice, pointer dan channel

// membuat function untuk menerapkan jikalau atribut nama tidak di isi maka akan diberikan default value nil
func NewMap(name string) map[string]string{
	// memberikan pengkondisian
	if name == "" {
		return nil
	} else {
		return map[string] string{
			"name": name,
		}
	}
}


func main() {
	// membuat data yang kosong dengan memanggil function NewMap
	data := NewMap("")

	if data == nil {
		fmt.Println("Data map masih kosong")
	} else {
		fmt.Println(data["name"])
	}

	// fmt.Println(data)
}