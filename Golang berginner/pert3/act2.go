package main

import "fmt"

func main() {
	var uts, uas float32

	defer fmt.Println("================================")
	fmt.Println("================================")
	fmt.Print("Masukkan nilai UTS: ")
	fmt.Scanln(&uts)
	fmt.Print("Masukkan nilai UAS: ")
	fmt.Scanln(&uas)

	total := uts*0.3 + uas*0.7 //UTS 30% dan UAS 70%

	var grade string
	switch {
	case total >= 81:
		grade = "A"
	case total >= 61:
		grade = "B"
	case total >= 41:
		grade = "C"
	case total >= 21:
		grade = "D"
	default:
		grade = "E"
	}

	fmt.Println("\n======= Hasil Nilai =======")
	fmt.Printf("Nilai UTS: %.2f\n", uts)
	fmt.Printf("Nilai UAS: %.2f\n", uas)
	fmt.Printf("Nilai Total: %.2f\n", total)
	fmt.Printf("Grade: %s\n", grade)
}
