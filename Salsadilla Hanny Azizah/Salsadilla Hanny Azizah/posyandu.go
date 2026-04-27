package main

import "fmt"

const NMAX = 100

type arrBalita [NMAX]float64

func hitungMinMax(arr arrBalita, n int, min *float64, max *float64) {
	*min = arr[0]
	*max = arr[0]

	for i := 1; i < n; i++ {
		if arr[i] < *min {
			*min = arr[i]
		}
		if arr[i] > *max {
			*max = arr[i]
		}
	}
}

func rerata(arr arrBalita, n int) float64 {
	var total float64
	for i := 0; i < n; i++ {
		total += arr[i]
	}
	return total / float64(n)
}

func main() {
	var data arrBalita
	var n int
	var min, max float64

	fmt.Print("Masukkan banyak data: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan berat ke-%d: ", i+1)
		fmt.Scan(&data[i])
	}

	hitungMinMax(data, n, &min, &max)
	rata := rerata(data, n)

	fmt.Printf("Berat minimum: %.2f kg\n", min)
	fmt.Printf("Berat maksimum: %.2f kg\n", max)
	fmt.Printf("Rata rata: %.2f kg\n", rata)
}