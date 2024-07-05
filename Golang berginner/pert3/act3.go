package main
import "fmt"
func main() {
	var i int32
	for j := 1; j <= 10; j++ {
		fmt.Println("======================")
		fmt.Print("Input Nilai = ")
		fmt.Scanln(&i)
		if i <= 10 {
			if i%2 == 0 { //i%2 sisa bagi (modulus)
				fmt.Println(i, " bilangan genap\n")
			} else {
				fmt.Println(i, " bilangan ganjil\n")
			}
		} else { 
			j = int(i)
			fmt.Println("Pertanyaan selesai")
		}}}
