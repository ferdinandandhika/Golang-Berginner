package main
import "fmt"

const pi = 22 / 7
var jari float64

func main() {
	defer fmt.Println(" =============== SELESAI =============== ") //tambahin baris ini
	fmt.Println("==========================================")
	fmt.Println("    PROGRAM PERHITUNGAN LUAS LINGAKRAN")
	fmt.Println("==========================================")
	fmt.Print("Masukan jari-jari lingkaran \t= ")
	fmt.Scan(&jari)
	luas := pi * (jari * jari)
	fmt.Println("Luas Lingkaran = ", luas)
}
