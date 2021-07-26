package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	var person1 = Person{"Eddie", 29}
	var person2 = Person{age: 28, name: "Connie"}
	fmt.Println(person1)
	fmt.Println(person1.name, person1.age)

	fmt.Println(person2)
	fmt.Println(person2.name, person2.age)

	person2.name = "康尼"
	fmt.Println(person2)
	fmt.Println(person2.name, person2.age)
}
