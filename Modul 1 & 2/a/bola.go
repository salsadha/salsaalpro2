package main

import "fmt"

func main() {
	var jarijari int
	var volume, luas float64

	fmt.Print("Jari jari : ")
	fmt.Scan(&jarijari)

	volume = 4.0/3.0 * 3.14 * float64(jarijari) * float64(jarijari) * float64(jarijari)
	luas = 4.0 * 3.14 * float64(jarijari) *float64(jarijari)

	fmt.Print("Bola diatas memiliki jari jari ", jarijari, " volume ", volume, " dan luas bola ", luas)
}