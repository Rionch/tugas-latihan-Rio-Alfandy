package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var hargaBaju, hargaCelana, hargaSepatu int
	hargaBaju = 100000
	hargaCelana = 75000
	hargaSepatu = 200000

	fmt.Println("Silahkan pilih barang yang ingin dibeli!")
	fmt.Println("Daftar Barang:")
	fmt.Println("1. Baju - Harga: Rp", hargaBaju)
	fmt.Println("2. Celana - Harga: Rp", hargaCelana)
	fmt.Println("3. Sepatu - Harga: Rp", hargaSepatu)

	var totalBiaya int
	var selesai bool

	for !selesai {
		fmt.Print("Pilih barang yang akan dibeli (1/2/3) atau 0 untuk selesai: ")
		pilihanBrg := input()

		pilihan, err := strconv.Atoi(pilihanBrg)
		if err != nil {
			fmt.Println("Pilihan tidak valid. Harap masukkan nomor yang valid.")
		} else if pilihan == 0 {
			selesai = true
		} else if pilihan != 1 && pilihan != 2 && pilihan != 3 {
			fmt.Println("Pilihan tidak valid. Harap masukkan nomor yang valid.")
		} else {
			fmt.Print("Masukkan jumlah barang yang dibeli: ")
			jumlahBrg := input()

			jumlah, err := strconv.Atoi(jumlahBrg)
			if err != nil {
				fmt.Println("Jumlah tidak valid. Harap masukkan angka yang valid.")
			} else {
				harga := 0
				if pilihan == 1 {
					harga = hargaBaju
				} else if pilihan == 2 {
					harga = hargaCelana
				} else if pilihan == 3 {
					harga = hargaSepatu
				}

				totalBiaya += harga * jumlah

				fmt.Printf("Total biaya: Rp %d\n", totalBiaya)
			}
		}
	}

	if totalBiaya > 0 {
		fmt.Print("Masukkan nama Struk untuk menyimpan invoice: ")
		namaFile := input()

		file, err := os.Create(namaFile)
		if err != nil {
			fmt.Println("Gagal membuat file.")
			return
		}
		file.Close()

		file, err = os.OpenFile(namaFile, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			fmt.Println("Gagal membuka file.")
			return
		}

		_, err = file.WriteString(fmt.Sprintf("Total biaya: Rp %d\n", totalBiaya))
		if err != nil {
			fmt.Println("Gagal menulis ke file.")
			return
		}

		fmt.Println("Invoice telah disimpan ke dalam file", namaFile)
		file.Close()

	}
}

func input() string {
	var input string
	fmt.Scanln(&input)
	return input
}
