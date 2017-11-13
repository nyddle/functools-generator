package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
)

type Person struct {
	name   string
	age    int
	gender bool //0 - female, 1 - male
}

func filter(list []Person, predicate func(p *Person) bool) []Person {

	result := []Person{}
	for _, person := range list {
		if predicate(&person) {
			result = append(result, person)
		}
	}
	return result
}

func male(p *Person) bool {
	return p.gender
}

func adult(p *Person) bool {
	return p.age > 18
}

func main() {
	fmt.Print("Hello there!\n")

	niceComany := []Person{
		{gender: true, age: 20},
		{gender: false, age: 30},
		{gender: true, age: 10},
		{gender: true, age: 10},
	}

	spew.Dump("Adults: ", filter(niceComany, adult))
	spew.Dump("Males: ", filter(niceComany, male))
}
