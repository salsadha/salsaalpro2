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
	var rumah arrInt

	fmt.Print("Masukkan banyak daerah kerabat: ")
	fmt.Scan(&daerah)

	for i = 0; i < daerah; i++ {

		fmt.Println("\nDaerah ke-", i+1)

		fmt.Print("Masukkan banyak rumah kerabat: ")
		fmt.Scan(&m)

		fmt.Println("Masukkan nomor rumah kerabat:")

		for j = 0; j < m; j++ {
			fmt.Scan(&rumah[j])
		}

		selectionSort(&rumah, m)

		fmt.Println("Nomor rumah setelah diurutkan:")

		for j = 0; j < m; j++ {
			fmt.Print(rumah[j], " ")
		}

		fmt.Println()
	}
}