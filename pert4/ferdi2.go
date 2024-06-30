package main

import "fmt"

func main() {
	// Mendefinisikan sebuah slice bernama 'kursus' yang berisi beberapa mata kuliah.
	var kursus = []string{"dbms", "server os", "networking", "web", "desktop", "erp"} //Slice
	// Membuat slice baru 'kursus_saya' dari slice 'kursus' dengan indeks 1 hingga 5.
	kursus_saya := kursus[2:4]
	// Menambahkan elemen baru "Manajemen Informatika" ke slice 'kursus_saya' menggunakan fungsi append.
	kursus_saya = append(kursus_saya, "Manajemen Informatika")
	// Mencetak isi dari slice 'kursus'.
	fmt.Println("\nIsi dari kursus adalah ", kursus)
	// Mencetak panjang dan kapasitas dari slice 'kursus'.
	fmt.Printf("Panjang kursus %d dan kapasitas %d\n", len(kursus), cap(kursus))
	// Mencetak isi dari slice 'kursus_saya'.
	fmt.Println("\nIsi dari kursus_saya adalah ", kursus_saya)
	// Mencetak panjang dan kapasitas dari slice 'kursus_saya'.
	fmt.Printf("Panjang kursus_saya %d dan kapasitas %d\n", len(kursus_saya), cap(kursus_saya))
}
