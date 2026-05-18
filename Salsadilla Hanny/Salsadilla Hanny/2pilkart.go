// Nama : Salsadilla Hanny Azizah
// NIM  : 109082500014

package main

import "fmt"

func main() {
	var suara int
	var totalMasuk int = 0
	var suaraSah int = 0

	var calon [21]int

	for {
		fmt.Scan(&suara)
		totalMasuk++

		if suara == 0 {
			break
		}

		if suara >= 1 && suara <= 20 {
			calon[suara]++
			suaraSah++
		}
	}

	ketua := 1
	wakil := 1

	for i := 2; i <= 20; i++ {
		if calon[i] > calon[ketua] {
			ketua = i
		}
	}

	for i := 1; i <= 20; i++ {
		if i != ketua {
			wakil = i
			break
		}
	}

	for i := 1; i <= 20; i++ {
		if i != ketua {
			if calon[i] > calon[wakil] {
				wakil = i
			} else if calon[i] == calon[wakil] && i < wakil {
				wakil = i
			}
		}
	}

	fmt.Println("Suara masuk:", totalMasuk)
	fmt.Println("Suara sah:", suaraSah)
	fmt.Println("Ketua RT:", ketua)
	fmt.Println("Wakil ketua:", wakil)
}