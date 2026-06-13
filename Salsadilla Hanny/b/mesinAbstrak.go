package main

import "fmt"

type Domino struct {
	sisi1 int
	sisi2 int
	nilai int
	balak bool
}

type Dominoes struct {
	kartu  [28]Domino
	jumlah int
}

func kocokKartu(D *Dominoes) {
	var temp Domino

	if D.jumlah > 1 {
		temp = D.kartu[0]
		D.kartu[0] = D.kartu[1]
		D.kartu[1] = temp
	}
}

func ambilKartu(D *Dominoes) Domino {
	var hasil Domino

	if D.jumlah > 0 {
		D.jumlah--
		hasil = D.kartu[D.jumlah]
	}

	return hasil
}

func gambarKartu(d Domino, suit int) int {
	if suit == 1 {
		return d.sisi1
	}
	return d.sisi2
}

func nilaiKartu(d Domino) int {
	return d.sisi1 + d.sisi2
}

func galiKartu(D *Dominoes, target Domino) {
	var kartu Domino

	for D.jumlah > 0 {
		kartu = ambilKartu(D)

		if kartu.sisi1 == target.sisi1 ||
			kartu.sisi1 == target.sisi2 ||
			kartu.sisi2 == target.sisi1 ||
			kartu.sisi2 == target.sisi2 {
			return
		}
	}
}

func sepasangKartu(d1 Domino, d2 Domino) bool {
	return nilaiKartu(d1)+nilaiKartu(d2) == 12
}

func main() {
	var D Dominoes
	var i int
	var d1, d2 Domino
	var target Domino

	fmt.Print("Jumlah kartu: ")
	fmt.Scan(&D.jumlah)

	for i = 0; i < D.jumlah; i++ {
		fmt.Printf("Sisi 1 kartu %d: ", i+1)
		fmt.Scan(&D.kartu[i].sisi1)

		fmt.Printf("Sisi 2 kartu %d: ", i+1)
		fmt.Scan(&D.kartu[i].sisi2)

		D.kartu[i].nilai = D.kartu[i].sisi1 + D.kartu[i].sisi2
		D.kartu[i].balak = D.kartu[i].sisi1 == D.kartu[i].sisi2
	}

	kocokKartu(&D)
	_ = ambilKartu(&D)

	target.sisi1 = D.kartu[0].sisi1
	target.sisi2 = D.kartu[0].sisi2
	galiKartu(&D, target)

	fmt.Println("\nMasukkan kartu pertama")
	fmt.Print("Sisi 1: ")
	fmt.Scan(&d1.sisi1)
	fmt.Print("Sisi 2: ")
	fmt.Scan(&d1.sisi2)

	fmt.Println("\nMasukkan kartu kedua")
	fmt.Print("Sisi 1: ")
	fmt.Scan(&d2.sisi1)
	fmt.Print("Sisi 2: ")
	fmt.Scan(&d2.sisi2)

	if sepasangKartu(d1, d2) {
		fmt.Println("\nPasangan bernilai 12")
	} else {
		fmt.Println("\nBukan pasangan bernilai 12")
	}

	fmt.Println("\nSalsadilla Hanny Azizah_109082500014")
}

