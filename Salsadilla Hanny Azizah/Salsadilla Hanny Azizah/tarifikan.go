package main

import "fmt"

const NMAX = 1000

type arrFloat [NMAX]float64

func main() {
	var data arrFloat
	var x, y int

	fmt.Print("Masukkan jumlah ikan (x): ")
	fmt.Scan(&x)
	fmt.Print("Masukkan kapasitas wadah (y): ")
	fmt.Scan(&y)

	for i := 0; i < x; i++ {
		fmt.Printf("Masukkan berat ikan ke-%d: ", i+1)
		fmt.Scan(&data[i])
	}

	var totalWadah float64
	var count int

	fmt.Println("Total berat tiap wadah:")

	for i := 0; i < x; i++ {
		totalWadah += data[i]
		count++

		if count == y || i == x-1 {
			fmt.Printf("%.2f ", totalWadah)
			totalWadah = 0
			count = 0
		}
	}

	var total float64
	for i := 0; i < x; i++ {
		total += data[i]
	}

	rata := total / float64(x)
	fmt.Printf("\nRata-rata berat ikan: %.2f\n", rata)
}