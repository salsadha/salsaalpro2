package main

import "fmt"

func bintang(i, n int){
	if i > n {
		return
	}

	for j:= 0 ; j < i; j++{
		fmt.Print("*")
	}
	fmt.Println()

	bintang (i+1, n)
} 

func main() {
	var n int
	fmt.Scan(&n)

	bintang(1, n)
}