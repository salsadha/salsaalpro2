package main

import "fmt"

func main() {
	var klubA, klubB string
	fmt.Print("Klub A: ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B: ")
	fmt.Scan(&klubB)

	var skorA, skorB int
	var hasil []string
	i := 1

	for {
		fmt.Printf("Pertandingan %d: ", i)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			hasil = append(hasil, klubA)
			fmt.Println("Hasil:", klubA)
		} else if skorB > skorA {
			hasil = append(hasil, klubB)
			fmt.Println("Hasil:", klubB)
		} else {
			fmt.Println("Hasil: Draw")
		}

		i++
	}

	fmt.Println("Pertandingan selesai")
}