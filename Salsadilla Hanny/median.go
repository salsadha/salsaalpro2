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

func median(T arrInt, n int) int {

	if n%2 == 1 {
		return T[n/2]
	}

	return (T[(n/2)-1] + T[n/2]) / 2
}

func main() {
	var data arrInt
	var x, n int

	fmt.Println("Masukkan data bilangan:")
	fmt.Println("(Masukkan 0 untuk mencari median)")
	fmt.Println("(Masukkan -5313 untuk mengakhiri program)")

	fmt.Scan(&x)

	for x != -5313 {

		if x == 0 {

			selectionSort(&data, n)

			fmt.Println("Median data:", median(data, n))

		} else {

			data[n] = x
			n++

		}

		fmt.Scan(&x)
	}
}