package main

import "fmt"

func lowToUpper(kar rune) rune {
	// Fungsi menerima karakter alfabet dalam huruf kecil dan mengembalikan karakter alfabet dalam huruf besar
	if (kar == 'a') {
		return 'A' 
	} else if (kar == 'b') {
		return 'B' 
	} else if (kar == 'c') {
		return 'C' 
	} else if (kar == 'd') {
		return 'D' 
	} else if (kar == 'e') {
		return 'E' 
	} else if (kar == 'f') {
		return 'F' 
	} else if (kar == 'g') {
		return 'G' 
	} else if (kar == 'h') {
		return 'H'
	} else if (kar == 'i') {
		return 'I' 
	} else if (kar == 'j') {
		return 'J'
	} else if (kar == 'k') {
		return 'K' 
	} else if (kar == 'l') {
		return 'L' 
	} else if (kar == 'm') {
		return 'M' 
	} else if (kar == 'n') {
		return 'N' 
	} else if (kar == 'o') {
		return 'O'
	} else if (kar == 'p') {
		return 'P' 
	} else if (kar == 'q') {
		return  'Q'
	} else if (kar == 'r') {
		return 'R' 
	} else if (kar == 's') {
		return 'S' 
	} else if (kar == 't') {
		return 'T' 
	} else if (kar == 'u') {
		return 'U'		
	} else if (kar == 'v') {
		return 'V' 
	} else if (kar == 'w') {
		return 'W'
	} else if (kar == 'x') {
		return 'X'
	} else if (kar == 'y') {
		return 'Y'
	} else if (kar == 'z') {
		return 'Z'
	} else {
		return kar
	}
}	

func main () {
	var hurufBesar,hasil rune
	fmt.Scanf("%c",&hurufBesar)
	hasil = lowToUpper(hurufBesar)
	fmt.Printf("%c\n",hasil)
}