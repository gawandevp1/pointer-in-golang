package main

import "fmt"

func main() {
	i, j := 20, 1000

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 25   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
