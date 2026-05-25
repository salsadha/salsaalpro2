package main

import "fmt"

type arrInt [1000000]int

func selectionSort(T *arrInt, n int) {
	var i, j, idxMin, temp int

	i = 1

	for i <= n-1 {

		idxMin = i - 1
		j = i

		for j < n {

			if T[idxMin] > T[j] {
				idxMin = j
			}

			j = j + 1
		}

		temp = T[idxMin]
		T[idxMin] = T[i-1]
		T[i-1] = temp

		i = i + 1
	}
}

func main() {
	var daerah int
	var m, i, j int
	var angka arrInt

	fmt.Print("Masukkan jumlah daerah: ")
	fmt.Scan(&daerah)

	for i = 0; i < daerah; i++ {

		fmt.Println("\nDaerah ke-", i+1)

		fmt.Print("Masukkan jumlah nomor rumah: ")
		fmt.Scan(&m)

		fmt.Println("Masukkan nomor rumah kerabat:")

		for j = 0; j < m; j++ {
			fmt.Scan(&angka[j])
		}

		selectionSort(&angka, m)

		fmt.Println("Urutan rumah kerabat:")

		for j = 0; j < m; j++ {
			if angka[j]%2 == 1 {
				fmt.Print(angka[j], " ")
			}
		}

		for j = m - 1; j >= 0; j-- {
			if angka[j]%2 == 0 {
				fmt.Print(angka[j], " ")
			}
		}

		fmt.Println()
	}
}