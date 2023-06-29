package main

import "fmt"

func main() {
	var x float32
	var y float32
	var z float32
	var l float32

	fmt.Print("qual o valor da hora do seu trabalho? ")
	fmt.Scan(&x)

	fmt.Print("quantas horas você trabalha por dia? ")
	fmt.Scan(&y)

	fmt.Print("quantos dias você trabalha por semana? ")
	fmt.Scan(&z)

	fmt.Print("voce faltou qnts dias no mes? ")
	fmt.Scan(&l)

	fmt.Print("o seu salario do mes será ", ((x*y*z)*4)-(l*x*y))

}
