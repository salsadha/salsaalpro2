// package main 

// import "fmt" 

// func main() { 
// 	var kiri, kanan float64 
	
// 	for { 
// 		fmt.Print("Masukan berat belanjaan di kedua kantong: ") 
// 		fmt.Scan(&kiri, &kanan) 
// 	if kiri >= 9 || kanan >= 9 { 
// 		fmt.Println("Proses selesai.") 
// 		break 
// 		} 
// 	} 
// } 

package main 
import "fmt" 
func main() { 
var kiri, kanan float64 
 
    for { 
        fmt.Print("Masukan berat belanjaan di kedua kantong: ") 
        fmt.Scan(&kiri, &kanan) 
 
        total := kiri + kanan 
 
        if total > 150 || kiri < 0 || kanan < 0 { 
            fmt.Println("Proses selesai.") 
            break 
        } 
 
        selisih := kiri - kanan 
        if selisih < 0 { 
            selisih = -selisih 
        } 
 
        oleng := selisih >= 9 
 
        fmt.Println("Sepeda motor pak Andi akan oleng:", oleng) 
    } 
} 