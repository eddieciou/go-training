package main

import "fmt"

func addValue(number int) {
	number += 10
	fmt.Println("add function:", number)
}

func addPointer(numberPointer *int) {
	*numberPointer += 10
	fmt.Println("addPointer function:", *numberPointer)
}

func main() {
	// Pass by value
	var number = 10
	addValue(number)
	fmt.Println("main function:", number)

	// Pass by pointer
	addPointer(&number)
	fmt.Println("main function:", number)
}
