package main

import "fmt"

func main() {
	var x int
	fmt.Print("escreva um numero ")
	fmt.Scan(&x)
	fmt.Print("O sucessor do seu numero é ", x+1, " ,O antecessor é ", x-1)
}
