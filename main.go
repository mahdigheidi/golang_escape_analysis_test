package main

import (
	"fmt"
)

type rect struct {
	width, height int
}

func (r *rect) area() int {
	r.height = -10
	r.width = 6
	return r.height * r.width
}

func (r rect) perim() int {
	r.width = 2
	r.height = 8
	return 2*r.width * r.height
}

type person struct {
	name string
	age int
}

func newPerson(name string) person {
	var p person
	p.name = name
	p.age = 28
	return p
}

func main() {
	a := 1
	b := 2
	fmt.Printf("%p\n%p\n", &a, &b)
	myRect := rect{width: 3, height: 5}
	fmt.Printf("%p\n", &myRect)

	anotherRect := make([]rect, 1)
	anotherRect[0].width = 3
	anotherRect[0].height = 5
	fmt.Println(anotherRect[0])
	fmt.Printf("%p\n", &anotherRect)

}