package main

import "fmt"

func main() {
	for i := 1; i <= 15; i++ {
		if i%2 == 0 {
			continue
		}

		if i > 8 {
			break

		}

		fmt.Println("Angka ganjil", i)
	}
}
