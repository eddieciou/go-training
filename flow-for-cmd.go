package main

import "fmt"

func main() {
	// break
	//var x int
	//for x = 0; x < 5; x++ {
	//	if x == 3 {
	//		break
	//	}
	//	fmt.Println(x)
	//}

	// continue
	//var x int
	//for x=0; x<5 ; x++{
	//	if x%2==0 { // x 是偶數
	//		continue
	//	}
	//	fmt.Println(x)
	//}

	// 實際範例，持續讓使用者輸入文字，計算總和。直到使用者輸入0為止。
	var result = 0
	var input int
	for true {
		fmt.Print("請輸入一個數字(輸入0停止輸入)：")
		fmt.Scanln(&input)
		if input == 0 {
			break
		}
		result += input
		fmt.Println("總合為:", result)
	}
	fmt.Println("執行結束，總和為:", result)
}
