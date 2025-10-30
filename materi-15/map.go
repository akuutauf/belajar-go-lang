package main

import "fmt"

func main() {
	// cara pertama membuat map
	// person := map[string]string{} // [string] : tipe data key, string{} : tipe data value
	// person["name"] = "Taufik"
	// person["address"] = "Banyuwangi"

	// cara kedua membuat map
	person := map[string]string{
		"name":    "Taufik",
		"address": "Banyuwangi",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	
	 // jika mengakses key yang tidak ada akan menghasilkan kosong (string) atau 0 (number)
	fmt.Println(person["hoby"])

	person["hoby"] = "Reading"
	fmt.Println(person)
	fmt.Println(person["hoby"])

	// membuat map dengan cara baru dan implementasi delete
	// delete memerlukan object map dan juga key, ditujukan untuk menghapus data key dan valuenya

	book := make(map[string]string)
	book["title"] = "Sang Alkemis"
	book["author"] = "Paullo Coelho"
	book["pages"] = "15"

	fmt.Println(book)
	
	delete(book, "pages") // menghapus key dan value dari map
	fmt.Println(book)
}