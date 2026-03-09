// package main 
// import "fmt" 
 
// func main() { 
//     var k int 
//     var hasil float64 
 
//     fmt.Print("Nilai K = ") 
//     fmt.Scan(&k) 
 
//     pembilang := (4*k + 2) * (4*k + 2) 
//     penyebut := (4*k + 1) * (4*k + 3) 
 
//     hasil = float64(pembilang) / float64(penyebut) 
 
//     fmt.Printf("Nilai f(K) = %.10f\n", hasil) 
// } 

package main
import "fmt"

func main() {
	var k, i int
	var hasil, fk float64

	fmt.Print("Nilai K = ")
	fmt.Scan(&k)

	hasil = 1

	for i = 0; i <= k; i++ {
		pembilang := float64((4*i + 2) * (4*i + 2))
		penyebut := float64((4*i + 1) * (4*i + 3))

		fk = pembilang / penyebut
		hasil = hasil * fk
	}

	fmt.Printf("Nilai akar 2 = %.10f\n", hasil)
}