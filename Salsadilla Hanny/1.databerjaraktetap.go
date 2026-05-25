package main

import "fmt"

type arrInt [1000]int

func insertionSort(T *arrInt, n int) {
	var i, j, temp int

	i = 1

	for i <= n-1 {

		j = i
		temp = T[j]

		for j > 0 && temp < T[j-1] {

			T[j] = T[j-1]
			j = j - 1
		}

		T[j] = temp
		i = i + 1
	}
}

func main() {
	var data arrInt
	var n, x, i int
	var selisih int
	var tetap bool = true

	fmt.Println("Masukkan bilangan bulat:")
	fmt.Println("(Masukkan bilangan negatif untuk berhenti)")

	fmt.Scan(&x)

	for x >= 0 {

		data[n] = x
		n++

		fmt.Scan(&x)
	}

	insertionSort(&data, n)

	fmt.Println("Data setelah diurutkan:")

	for i = 0; i < n; i++ {
		fmt.Print(data[i], " ")
	}

	fmt.Println()

	selisih = data[1] - data[0]

	for i = 2; i < n; i++ {

		if data[i]-data[i-1] != selisih {
			tetap = false
		}
	}

	if tetap {
		fmt.Println("Data berjarak", selisih)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}