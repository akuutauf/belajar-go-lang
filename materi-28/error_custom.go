package main

import (
	"fmt"
)

// error custom
// digunakan untuk kustomisasi dari function error yang sudah ada, dengan memperhatikan kontrak dari interface error

// contoh function validation error (yang sesuai kontrak error)
// membuat struct
type validationError struct{
	Message string
}

// membuat method struct validation error
// method struct usahakan menggunakan pointer
// memanggil interfae error dengan 'Error()'
func (v *validationError) Error() string {
	return v.Message
}

// contoh not found error (yang sesuai kontrak error)
// membuat struct
type notFoundError struct {
	Message string
}

// membuat method struct not found error
func (n *notFoundError) Error() string {
	return n.Message
}

// membuat function save data
func saveData(id string, data any) error {
	if id == "" {
		return &validationError{Message: "validation error"}
	} else if id != "taufik" {
		return &notFoundError{Message: "data not found"}
	}

	// ok: jika tidak selain pengkondisian diatas, maka hasinya mengembalikan nil: atau tidak ada error
	return nil
}

func main() {
	// memanggil save data
	err := saveData("ilham", nil)

	// memastikan error yang terjadi dengan memberikan pengecekan
	// maksud dari 'ok' di sini adalah bernilai boolean, true/false
	// ok juga digunakan untuk 'apakah dia bisa dikonversi errornya?' bernilai boolean true/false
	if err != nil {
		// jika ada eror maka,
		// melakukan pengecekan ulang, dicek apakah error yang diterima adalah error validasi, not found atau error lain
		if validationErr, ok := err.(*validationError); ok {
			fmt.Println("Ini adalah error jenis validasi error:", validationErr.Error())
		} else if notFoundErr, ok := err.(*notFoundError); ok {
			fmt.Println("Ini adalah error jenis not found error:", notFoundErr.Error())
		} else {
			fmt.Println("Ini jenis error yang lain", err.Error())
		}

	// atau pengecekan nya bisa menggunakan switch case, seperti ini
	switch finalErr := err.(type) {
		case *validationError:
			fmt.Println("Ini adalah error jenis validasi error:", finalErr.Error())
		case *notFoundError:
			fmt.Println("Ini adalah error jenis not found error:", finalErr.Error())
		default:
			fmt.Println("Ini jenis error yang lain", finalErr.Error())
	}

	} else {
		// jika tidak ada eror, maka sukses
		fmt.Println("Sukses, tidak ada error")
	}
}