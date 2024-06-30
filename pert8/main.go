package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Lepkom struct {
	Nama   string `json:"nama_mahasiswa"`
	Kursus string `json:"kursus_mahasiswa"`
	Foto   string `json:"foto_mahasiswa"`
}

var data_mahasiswa = []Lepkom{
	{
		Nama:   "Ferdinand Andhika",
		Kursus: "Beginner Berginner",
		Foto:   "img/gambar1.jpg",
	},
	{
		Nama:   "Nedianti Widhiyan",
		Kursus: "Fundamental Web",
		Foto:   "img/gambar2.jpg",
	},
}

func user(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodGet {
		result, err := json.Marshal(data_mahasiswa)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}
}

func main() {
	port := ":8081" //diganti dengan NPM
	http.HandleFunc("/api/mahasiswa", user)

	fmt.Println("web berjalan di port localhost" + port + "/api/mahasiswa")
	http.ListenAndServe(port, nil)
}
