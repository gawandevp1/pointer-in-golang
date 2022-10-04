// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type Student struct {
	Name string
}

func a(s *Student) *Student {
	s.Name = "C.M. Eknath Shinde"
	return s
}

func main() {
	object := Student{
		Name: "C.M. Uddhav Thakare",
	}

	a(&object)
	fmt.Println(object.Name)
}