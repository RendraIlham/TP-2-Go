package main

import "fmt"

func adaBilanganM(n,m int) string {
	// Fungsi menerima masukan berupa bilanga bulat n dan m, serta mengembalikan string "YA" jika m terdapat pada n. Mengembalikan string "TIDAK" jika pada n tidak terdapat m
	var i,angka int
	
	for i = n; i <= 999999; i++ {
		
		angka = n % 10
		
		if (m >= 0) && (m <= 9) && (angka == m) {
			return "YA"
		}
		n = n / 10	
	} 
	return "TIDAK"
}



func main () {
	var n,m int
	fmt.Scan(&n,&m)
	fmt.Print(adaBilanganM(n,m))
}	