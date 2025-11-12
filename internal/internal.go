package internal

import "fmt"

// fungsi utama, yang akan dijalankan pertama kali saat package dipanggil
func init() {
 fmt.Println("This is for internal only")
}