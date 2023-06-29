package main

import "fmt"

func main() {
	var peso, libras float32
	fmt.Print("coloque o peso em kg para ser convertido ")
	fmt.Scan(&peso)
	if peso >= 0 {
		fmt.Print(peso*2.205, "libras")
	} else {
		fmt.Print("coloque o peso em libras para ser convertido ")
		fmt.Scan(&libras)
		fmt.Println(libras/2.205, "kg")
	}
}
