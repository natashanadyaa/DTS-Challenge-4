package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	biodata := map[int]map[string]string{
		1: {
			"nama":      "Natasha",
			"alamat":    "Jeruk Manis V no.92",
			"pekerjaan": "Mahasiswa",
			"alasan":    "Karena di kampus saya tidak belajar bahasa pemrograman golang",
		},
		2: {
			"nama":      "Nadya",
			"alamat":    "Jeruk Nipis III no.3",
			"pekerjaan": "Mahasiswa",
			"alasan":    "Mau sertifikat",
		},
		3: {
			"nama":      "Harjanto",
			"alamat":    "Jeruk Kecut VI no.8",
			"pekerjaan": "Mahasiswa",
			"alasan":    "Tambah pengalaman",
		},
		4: {
			"nama":      "Natnut",
			"alamat":    "Jeruk Asam VII no.10",
			"pekerjaan": "Penyanyi",
			"alasan":    "Mempercantik CV",
		},
	}
	if len(os.Args) < 2 {
		fmt.Println("Anda belum memasukkan nomor identitas")
		return
	}

	id, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Nomor identitas harus berupa angka")
		return
	}

	if data, ok := biodata[id]; ok {
		fmt.Println("Nama:      ", data["nama"])
		fmt.Println("Alamat:    ", data["alamat"])
		fmt.Println("Pekerjaan: ", data["pekerjaan"])
		fmt.Println("Alasan:    ", data["alasan"])
	} else {
		fmt.Println("Data biodata dengan nomor identitas tersebut tidak ditemukan")
	}
}
