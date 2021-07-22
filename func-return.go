package main

import "fmt"

func show(message string) {
	if message == "Hello" {
		return
	}
	fmt.Println(message)
}

func add(number1 int, number2 int) int {
	return number1 + number2
}

func test() (int, string) {
	return 2, "Hello"
}

func main() {
	// 基本的return 運作方式
	//show("Hello")
	//show("你好")

	// 單一的回傳值
	//var result = add(1,2)
	//fmt.Println(result)

	// 多個回傳值
	var number int
	var message string
	number, message = test()
	fmt.Println(number, message)
}
