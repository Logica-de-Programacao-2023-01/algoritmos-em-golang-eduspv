package main

import "fmt"

func main() {
	var x int
	var y int
	var z int

	fmt.Print("primeiro numero ")
	fmt.Scan(&x)
	fmt.Print("segundo numero ")
	fmt.Scan(&y)
	fmt.Print("terceiro numero ")
	fmt.Scan(&z)

	fmt.Print("a sua media ponderada é ", (x*2+y*3+z*5)/(x+y+z))
}
