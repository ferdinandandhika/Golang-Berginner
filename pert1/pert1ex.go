package main

import "fmt" // Mengimpor package fmt untuk input/output

func main() {
    // Deklarasi variabel string name1 dengan nilai "ferdi"
    var name1 string = "ferdi"

    // Deklarasi variabel string name2 dengan nilai "Cindy"
    var name2 string = " Cindy"

    // Deklarasi variabel string name3 menggunakan pendekatan penugasan pendek, dengan nilai "Belajar Golang"
    name3 := "Belajar Golang"

    // Deklarasi variabel integer age dengan nilai 20
    age := 20

    // Mencetak nilai variabel name1 ke console tanpa baris baru
    fmt.Print(name1)

    // Mencetak nilai variabel name2 ke console dan memberikan baris baru
    fmt.Println(name2)

    // Mencetak string "Cakra " ke console tanpa baris baru
    fmt.Print("Cakra ")

    // Mencetak nilai variabel name3 dan age ke console dengan baris baru
    fmt.Println(name3, age)
}
