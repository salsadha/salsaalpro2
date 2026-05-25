package main

import "fmt"

const nMax int = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating int
}

type DaftarBuku [nMax]Buku

func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	var i int

	for i = 0; i < n; i++ {

		fmt.Println("\nData Buku ke-", i+1)

		fmt.Print("Masukkan ID Buku        : ")
		fmt.Scan(&pustaka[i].id)

		fmt.Print("Masukkan Judul Buku     : ")
		fmt.Scan(&pustaka[i].judul)

		fmt.Print("Masukkan Penulis Buku   : ")
		fmt.Scan(&pustaka[i].penulis)

		fmt.Print("Masukkan Penerbit Buku  : ")
		fmt.Scan(&pustaka[i].penerbit)

		fmt.Print("Masukkan Jumlah Buku    : ")
		fmt.Scan(&pustaka[i].eksemplar)

		fmt.Print("Masukkan Tahun Terbit   : ")
		fmt.Scan(&pustaka[i].tahun)

		fmt.Print("Masukkan Rating Buku    : ")
		fmt.Scan(&pustaka[i].rating)
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	var i, idx int

	idx = 0

	for i = 1; i < n; i++ {

		if pustaka[i].rating > pustaka[idx].rating {
			idx = i
		}
	}

	fmt.Println("\nData Buku Terfavorit")
	fmt.Println("Judul     :", pustaka[idx].judul)
	fmt.Println("Penulis   :", pustaka[idx].penulis)
	fmt.Println("Penerbit  :", pustaka[idx].penerbit)
	fmt.Println("Tahun     :", pustaka[idx].tahun)
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	var i, j int
	var temp Buku

	i = 1

	for i <= n-1 {

		j = i
		temp = pustaka[j]

		for j > 0 && temp.rating > pustaka[j-1].rating {

			pustaka[j] = pustaka[j-1]
			j = j - 1
		}

		pustaka[j] = temp
		i = i + 1
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	var i int

	fmt.Println("\n5 Buku Dengan Rating Tertinggi")

	if n > 5 {
		n = 5
	}

	for i = 0; i < n; i++ {
		fmt.Println(i+1, ".", pustaka[i].judul)
	}
}

func CariBuku(pustaka DaftarBuku, n, r int) {
	var kiri, kanan, tengah int
	var found bool = false

	kiri = 0
	kanan = n - 1

	for kiri <= kanan && !found {

		tengah = (kiri + kanan) / 2

		if pustaka[tengah].rating == r {

			found = true

		} else if r < pustaka[tengah].rating {

			kiri = tengah + 1

		} else {

			kanan = tengah - 1
		}
	}

	if found {

		fmt.Println("\nData Buku Ditemukan")
		fmt.Println("Judul      :", pustaka[tengah].judul)
		fmt.Println("Penulis    :", pustaka[tengah].penulis)
		fmt.Println("Penerbit   :", pustaka[tengah].penerbit)
		fmt.Println("Tahun      :", pustaka[tengah].tahun)
		fmt.Println("Eksemplar  :", pustaka[tengah].eksemplar)
		fmt.Println("Rating     :", pustaka[tengah].rating)

	} else {

		fmt.Println("\nTidak ada buku dengan rating seperti itu")
	}
}

func main() {
	var pustaka DaftarBuku
	var n, ratingCari int

	fmt.Print("Masukkan jumlah buku: ")
	fmt.Scan(&n)

	DaftarkanBuku(&pustaka, n)

	CetakTerfavorit(pustaka, n)

	UrutBuku(&pustaka, n)

	Cetak5Terbaru(pustaka, n)

	fmt.Print("\nMasukkan rating buku yang dicari: ")
	fmt.Scan(&ratingCari)

	CariBuku(pustaka, n, ratingCari)
}