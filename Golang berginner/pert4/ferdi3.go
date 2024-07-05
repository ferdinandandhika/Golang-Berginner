package main

import "fmt"

func main() {
	// Informasi mahasiswa disimpan dalam map 'data' dengan NPM sebagai kunci
	var data = map[string]mahasiswa{
		"50421513": {
			"Ferdinand Andhika Widhiyan",
			"3IA20",
		},
		"51421134": {
			"Nedianti Rellya",
			"3IA21",
		},
		"51421587": {
			"Cindy Maharani",
			"3IA20",
		},
	}

	for {
		// Meminta pengguna untuk memasukkan NPM
		var search string
		fmt.Print("Masukkan NPM Anda : ")
		fmt.Scan(&search)

		// Mencari NPM dalam map 'data'
		value, ok := data[search]
		if ok {
			// Jika NPM ditemukan, mencetak informasi nama dan kelas mahasiswa
			fmt.Printf("Nama Anda %s \nKelas Anda %s\n", value.Nama, value.Kelas)
		} else {
			// Jika NPM tidak ditemukan, mencetak pesan kesalahan
			fmt.Println("Data tidak ditemukan!")
		}

		// Meminta pengguna apakah ingin melanjutkan program atau tidak
		var lanjut string
		fmt.Println("==================================")
		fmt.Print("Apakah mau lanjutkan program Y/N : ")
		fmt.Scan(&lanjut)

		// Jika pengguna memilih untuk tidak melanjutkan, keluar dari loop
		if lanjut != "Y" && lanjut != "y" {
			fmt.Println("Program Selesai")
			break
		}
		fmt.Println("==================================")
	}
}

// Struktur 'mahasiswa' digunakan untuk menyimpan informasi nama dan kelas mahasiswa
type mahasiswa struct {
	Nama  string
	Kelas string
}
