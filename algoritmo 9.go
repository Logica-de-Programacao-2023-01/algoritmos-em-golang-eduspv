package main

import "fmt"

func main() {
	var p float32
	fmt.Print("valor:")
	fmt.Scan(&p)
	fmt.Print("o valor com desconto será de ", p-(p*0.10))

}
