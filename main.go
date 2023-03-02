package main

import (
	"fmt"
)


type rect struct {
	width, height int
}


// func area() read a pointer to rect and changes its values
// because the variable passed is a pointer, the real object will also be modified
func (r *rect) area() int {
	r.height = -10
	r.width = 6
	return r.height * r.width
}


// func perim gets a copy of the object, within its own scope,
// the changes in this function will not be applied to the main object
func (r rect) perim() int {
	r.width = 2
	r.height = 8
	fmt.Println(r)
	return 2*r.width * r.height
}

type person struct {
	name string
	age uint32
}

func newPerson(name string, age uint32) person {
	var p person
	p.name = name
	p.age = age
	return p
}

func newPointerOfPerson(name string, age uint32) *person {
	someGuy := person{name, age}
	return &someGuy
}

func newPointerOfPersonWithVar(name string, age uint32) *person {
	var p person
	p.name = name
	p.age = age
	return &p
}


func main() {

	// allocating two variables to view their sequential addresses in stack
	a := 1
	b := 2
	fmt.Printf("addr a: %p\naddr b: %p\n", &a, &b)

	c := make([]int, 1)
	c[0] = 3
	fmt.Printf("addr c: %p\n", &c)
	
	// allocating an array of struct in stack
	myRect := rect{3, 5}
	fmt.Printf("addr of rect object allocated in stack: %p\n", &myRect)


	// allocating the same struct in heap with make
	anotherRect := make([]rect, 1)
	anotherRect[0].width = 3
	anotherRect[0].height = 5
	fmt.Println(anotherRect)
	fmt.Printf("addr of rect object allocated with make in heap: %p\n", &anotherRect)

	fmt.Println(myRect)
	fmt.Println(myRect.area())
	fmt.Printf("The struct sent to perim(): ")
	fmt.Println(myRect)
	fmt.Printf("The struct in perim(): ")
	fmt.Println(myRect.perim())
	fmt.Printf("The struct after being editted in perim(): ")
	fmt.Println(myRect)
	
	x := newPerson("mahdi", 22)
	y := newPointerOfPerson("mahdi", 22)
	z := newPointerOfPersonWithVar("mahdi", 22)
	fmt.Printf("%T    %p\n", x, &x)
	fmt.Printf("%T    %p\n", y, &y)
	fmt.Printf("%T    %p\n", z, &z)
}