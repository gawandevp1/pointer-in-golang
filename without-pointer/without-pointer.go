package main

import "fmt"

func b(funcVar int) int {
	funcVar = 9
	return funcVar
}

func main() {
	mainVariable := 10

	b(mainVariable)
	fmt.Println(mainVariable)
}