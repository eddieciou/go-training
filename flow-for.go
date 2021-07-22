package main

import "fmt"

func main() {
	// 基本迴圈使用
	//var x int = 0
	//for x < 3 {
	//	fmt.Println(x)
	//	x++
	//}

	//var x int
	//for x = 0; x < 3; x++ {
	//	fmt.Println(x)
	//}

	//// 實際範例，計算1+2+3+4+...50 的結果
	var result = 0
	var counter int
	for counter = 1; counter <= 50; counter++ {
		result += counter
	}
	fmt.Println(result)
}
