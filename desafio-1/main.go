package main

import "fmt"

const ebulicaoK float64 = 373.15

func main() {
	tempK := ebulicaoK
	tempC := tempK - 273.15

	fmt.Printf("A temperatura de ebulição da água em K é: %.2f \nA temperatura de ebulição da água em C é: %.2f", tempK, tempC)
}
