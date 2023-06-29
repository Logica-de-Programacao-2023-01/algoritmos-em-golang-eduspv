package main

import "fmt"

func main() {
	var x float32

	fmt.Print("quantos anos voce tem?")
	fmt.Scan(&x)
	fmt.Print("voce tem aproximadamente ", x*365, " dias vividos")
}
