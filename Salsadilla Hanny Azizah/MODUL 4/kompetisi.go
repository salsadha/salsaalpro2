package main

import "fmt"

func hitungSkor(soal *int, skor *int) {
	var waktu int
	*soal = 0
	*skor = 0


	for i := 0; i < 8; i++{
		fmt.Scan(&waktu)

		if waktu <= 300{
			*soal = *soal + 1
			*skor = *soal + waktu
		}
	}
}

func main() {
		var nama, pemenang string
		var soal, skor int
		var maxSoal, minSkor int

		for {
			fmt.Scan(&nama)

			if nama == "Selesai" {
				break
			}

			hitungSkor(&soal, &skor)
			if soal > maxSoal || (soal == maxSoal && skor < minSkor) || maxSoal == 0 {
				pemenang = nama
				maxSoal = soal
				minSkor = skor
			}
		}

		fmt.Println(pemenang, maxSoal, minSkor)
	}