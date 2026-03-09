package main

import "fmt"

func main() {
	var bunga, pita string
	var i, jumlah int

	i = 1

	for {
		fmt.Print("Bunga ", i, ": ")
		fmt.Scan(&bunga)

		if bunga == "SELESAI" {
			break
		}

		pita = pita + bunga + " - "
		jumlah++
		i++
	}

	fmt.Println("Pita :", pita)
	fmt.Println("Bunga:", jumlah)
}