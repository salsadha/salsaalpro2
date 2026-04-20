package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Jumlah data: ")
	fmt.Scan(&n)

	arr := make([]int, n)

	// input
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	// a. tampil semua
	fmt.Println("Semua:", arr)

	// b. indeks ganjil
	fmt.Print("Indeks ganjil: ")
	for i := 1; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	// c. indeks genap
	fmt.Print("Indeks genap: ")
	for i := 0; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	// d. kelipatan x
	var x int
	fmt.Print("Masukkan x: ")
	fmt.Scan(&x)

	fmt.Print("Kelipatan indeks x: ")
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	// e. hapus indeks
	var idx int
	fmt.Print("Hapus indeks: ")
	fmt.Scan(&idx)

	arr = append(arr[:idx], arr[idx+1:]...)
	fmt.Println("Setelah hapus:", arr)

	// f. rata-rata
	total := 0
	for _, v := range arr {
		total += v
	}
	mean := float64(total) / float64(len(arr))
	fmt.Println("Rata-rata:", mean)

	// g. standar deviasi
	var sum float64
	for _, v := range arr {
		sum += math.Pow(float64(v)-mean, 2)
	}
	std := math.Sqrt(sum / float64(len(arr)))
	fmt.Println("Std Dev:", std)

	// h. frekuensi
	var cari int
	fmt.Print("Cari angka: ")
	fmt.Scan(&cari)

	count := 0
	for _, v := range arr {
		if v == cari {
			count++
		}
	}
	fmt.Println("Frekuensi:", count)
}