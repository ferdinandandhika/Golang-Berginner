package main

import "fmt"

var nilai1, nilai2, nilai3, nilai4, nilai5 float64

func main() {
	defer fmt.Println(" =============== SELESAI =============== ")
	fmt.Println("==========================================")
	fmt.Println("     PROGRAM HITUNG NILAI OPERATOR")
	fmt.Println("==========================================")
	fmt.Print("Masukkan nilai 1 = ")
	fmt.Scan(&nilai1)
	fmt.Print("Masukkan nilai 2 = ")
	fmt.Scan(&nilai2)
	fmt.Print("Masukkan nilai 3 = ")
	fmt.Scan(&nilai3)
	fmt.Print("Masukkan nilai 4 = ")
	fmt.Scan(&nilai4)

	fmt.Println("==========================================")
	fmt.Println("                  Hasil")
	fmt.Println("==========================================")
	fmt.Println("Nilai 1 * Nilai 2 = ", nilai1*nilai2)
	fmt.Println("Nilai 2 + Nilai 3 = ", nilai2+nilai3)
	fmt.Println("Nilai 3 / Nilai 4 = ", nilai3/nilai4)
	fmt.Println("Nilai 4 - Nilai 1 = ", nilai4-nilai1)
}
