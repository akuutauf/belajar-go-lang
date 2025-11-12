package database

// package initialization
// digunakan untuk menjalankan program pada saat fungsi dipanggil pertama kali

var connection string

// fungsi init
func init() {
	connection = "MySql"
}

func GetDatabase() string {
	return connection
}