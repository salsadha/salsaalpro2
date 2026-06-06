package main

import "fmt"

func main() {
	var x, data string
	var n, i int
	var jumlah int
	var posisi int
	var ditemukan bool

	fmt.Scan(&x)
	fmt.Scan(&n)

	posisi = -1

	for i = 1; i <= n; i++ {
		fmt.Scan(&data)

		if data == x {
			jumlah++
			ditemukan = true

			if posisi == -1 {
				posisi = i
			}
		}
	}

	if ditemukan {
		fmt.Println("String ditemukan")
		fmt.Println("Posisi pertama:", posisi)
	} else {
		fmt.Println("String tidak ditemukan")
	}

	fmt.Println("Jumlah kemunculan:", jumlah)

	if jumlah >= 2 {
		fmt.Println("Sedikitnya terdapat dua string", x)
	} else {
		fmt.Println("Kurang dari dua string", x)
	}
}