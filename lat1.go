package main

import "fmt"

func reamur (C float64) float64 {
	var r float64
	r = 0.80 * C
	return r 
}

func fahrenheit (C float64) float64 {
	var f float64
	f = (1.80 * C) + 32.00
	return  f 
}

func kelvin (C float64) float64 {
	var k float64
	k = C + 273.00
	return k 
}

func main () {
	var Cawal,Cakhir,i,R,F,K,step float64
	fmt.Scan(&Cawal,&Cakhir,&step)
	fmt.Println("Celcius\tReamur\tFahrenheit\tKelvin")
	for i = Cawal; i <= Cakhir; i = i + step {
		R = reamur(Cawal)
		F = fahrenheit(Cawal)
		K = kelvin(Cawal)

	
		fmt.Printf("%.2f\t%.2f\t%.2f\t%.2f\n",Cawal,R,F,K)
		Cawal = Cawal + step
	}
}