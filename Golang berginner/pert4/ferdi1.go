package main

import "fmt"

func main() {
	var ekspetasiGaji, gajiSekarang int

	fmt.Print("Masukkan gaji Anda:")
	fmt.Scan(&gajiSekarang)

	fmt.Print("Masukkan gaji yang Anda inginkan: ")
	fmt.Scan(&ekspetasiGaji)

	/* Pada fungsi main, parameter gajiSekarang
	   diubah menjadi pointer &gajiSekarang  */

	naikanGaji(&gajiSekarang, ekspetasiGaji)

	fmt.Printf("\nGaji Anda sekarang %d\n", gajiSekarang)

}

/*
Pada definisi fungsi naikanGaji, parameter gajiAwal
diubah menjadi pointer *int dan dalam fungsi
*/
func naikanGaji(gajiAwal *int, gajiAkhir int) {
	*gajiAwal = gajiAkhir

	// saat memanggil fungsi naikanGaji. nilai dari gajiAwal
	// diberikan nilai gajiAkhir dengan *gajiAwal = gajiAkhir.
}
