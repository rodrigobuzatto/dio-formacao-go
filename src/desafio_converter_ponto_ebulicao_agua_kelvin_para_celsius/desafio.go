package main

import "fmt"

const pontoEbulicaoKelvin = 373

func main() {
	pontoEbulicaoCelsius := pontoEbulicaoKelvin - 273
	fmt.Printf("A temperatura de ebulição da água em K é %d e a temperatura de ebulição da água em C é %d", pontoEbulicaoKelvin, pontoEbulicaoCelsius)
}
