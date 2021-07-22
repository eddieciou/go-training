package main

import "fmt"

func test(message string) {
	fmt.Println(message)
}

func add(number1 int, number2 int) {
	var result = number1 + number2
	fmt.Println(result)
}

func sum(maxNumber int) {
	var result = 0
	var counter int
	for counter = 1; counter <= maxNumber; counter++ {
		result += counter
	}
	fmt.Println(result)
}

func main() {
	// 基礎韓式語法演練
	//test("Hello")
	//test("Go")
	//add(1, 4)
	//add(1, -4)

	// 計算 1+2+3+...+10 的函式包裝
	sum(10)
	sum(50)
}
