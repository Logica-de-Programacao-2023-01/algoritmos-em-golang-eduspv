package main

import "fmt"

func main() {
	var x float32

	fmt.Print("qual o seu salario?")
	fmt.Scan(&x)
	fmt.Print("o seu salario com aumento será de ", x*1.15)

}
