package main

import "fmt"

func adaBilanganM (n,m int) string {
	var i,bil int
	for i = n; i <= 999999; i++ {

		bil = n % 10

		if (m >= 0) && (m <= 9) && (bil == m) {
			return "YA"
		}

		n = n / 10
	}
	return "TIDAK"
}

func main () {
	var n,m int
	var hasil string
	fmt.Scan(&n,&m)
	hasil = adaBilanganM(n,m)
	fmt.Print(hasil)

}