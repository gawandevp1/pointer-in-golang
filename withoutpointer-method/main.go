// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type Student struct {
	Name string
}

func a(s Student) Student {
	s.Name = "Mr. Rahul Gandhi"
	return s
}

func main() {
	object := Student{
		Name: "P.M. Narendra Modi",
	}

	a(object)
	fmt.Println(object.Name)
}
