package main

import "fmt"

func main() {
	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	for i, fruit := range fruits {
		fmt.Printf("elemen %d :%s\n", i+1, fruit)
	}
}
