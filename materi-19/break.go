package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {

		// penggunaan break 
		if i == 5 {
			break
		}

		fmt.Println("Perulangan ke-", i)
	}
}