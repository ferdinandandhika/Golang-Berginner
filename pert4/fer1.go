package main

import "fmt"

func main() {
	var ekspetasiGaji, gajiSekarang int

	fmt.Print("Masukkan gaji Anda : ")
	fmt.Scan(&gajiSekarang)

	fmt.Print("Masukkan gaji yang Anda inginkan : ")
	fmt.Scan(&ekspetasiGaji)

	naikanGaji(&gajiSekarang, ekspetasiGaji)

	fmt.Printf("\nGaji Anda sekarang %d\n", gajiSekarang)
}

func naikanGaji(gajiAwal *int, gajiAkhir int) {
	*gajiAwal = gajiAkhir
}
