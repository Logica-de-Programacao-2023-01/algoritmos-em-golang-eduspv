package main

import "fmt"

func main() {
	var x int

	fmt.Print("insira o seu numero ")
	fmt.Scan(&x)
	fmt.Print("o dobro do seu numero é ", x*2, " o triplo é ", x*3, " o quadruplo do seu numero é ", x*4)

}
