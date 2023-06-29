package main

import "fmt"

func main() {
	var x int
	var y int
	var z int

	fmt.Print("indique o primeiro numero ")
	fmt.Scan(&x)
	fmt.Print("indique segundo numero ")
	fmt.Scan(&y)
	fmt.Print("indique o terceiro numero")
	fmt.Scan(&z)
	fmt.Print("seu resultado é ", x+y+z)

}
