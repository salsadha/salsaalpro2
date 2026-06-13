// Salsadilla Hanny Azizah_109082500014


package main

import "fmt"

var pita string
var idx int
var karakter byte

func start() {
	idx = 0
	karakter = pita[idx]
}

func maju() {
	idx++

	if idx < len(pita) {
		karakter = pita[idx]
	}
}

func eop() bool {
	return karakter == '.'
}

func cc() byte {
	return karakter
}

func main() {
	var jumlahKarakter int
	var jumlahA int
	var jumlahLE int
	var frekuensi float64
	var prev byte

	fmt.Print("Masukkan karakter (akhiri dengan .): ")
	fmt.Scan(&pita)

	start()

	prev = ' '

	for !eop() {

		jumlahKarakter++

		if cc() == 'A' {
			jumlahA++
		}

		if prev == 'L' && cc() == 'E' {
			jumlahLE++
		}

		prev = cc()

		maju()
	}

	if jumlahKarakter > 0 {
		frekuensi = float64(jumlahA) / float64(jumlahKarakter)
	}

	fmt.Println("Jumlah karakter =", jumlahKarakter)
	fmt.Println("Jumlah huruf A =", jumlahA)
	fmt.Println("Frekuensi A =", frekuensi)
	fmt.Println("Jumlah LE =", jumlahLE)

	fmt.Println("Salsadilla Hanny Azizah_109082500014")

}
