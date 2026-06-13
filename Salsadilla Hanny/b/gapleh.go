// Salsadilla Hanny Azizah_109082500014

package main

import "fmt"

type Domain struct {
	sisi1 int
	sisi2 int
}

func bisaDisambung(a Domain, b Domain) bool {
	return a.sisi1 == b.sisi1 ||
		a.sisi1 == b.sisi2 ||
		a.sisi2 == b.sisi1 ||
		a.sisi2 == b.sisi2
}

func main() {
	var kartu1, kartu2 Domain

	fmt.Println("Kartu Pertama")
	fmt.Print("Sisi 1: ")
	fmt.Scan(&kartu1.sisi1)
	fmt.Print("Sisi 2: ")
	fmt.Scan(&kartu1.sisi2)

	fmt.Println("Kartu Kedua")
	fmt.Print("Sisi 1: ")
	fmt.Scan(&kartu2.sisi1)
	fmt.Print("Sisi 2: ")
	fmt.Scan(&kartu2.sisi2)

	if bisaDisambung(kartu1, kartu2) {
		fmt.Println("Kartu dapat disambungkan")
	} else {
		fmt.Println("Kartu tidak dapat disambungkan")
	}

	fmt.Println("Salsadilla Hanny Azizah_109082500014")

}
