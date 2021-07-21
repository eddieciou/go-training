package main

import "fmt"

func main() {
	// 算術運算： +, -, *, /
	var x int
	x = 3*3 + 10
	fmt.Println(x)
	// 指定運算： =. +=, -=, *=, /=
	x = 5
	// x = x + 2 簡化成 x += 2
	x += 2
	fmt.Println(x)
	// 單元運算： ++, __
	x++
	fmt.Println(x)
	// 比較運算： >, <, >=, <=, ==
	var result bool = 4 > 3
	fmt.Println(result)
	// 邏輯運算： !, && , ||
	result = !false
	fmt.Println(result)
	result = true && false
	fmt.Println(result)
	result = true || false
	fmt.Println(result)
}
