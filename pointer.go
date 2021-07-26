package main

import "fmt"

func main() {
	// 1. 建立存放資料的變數
	var number = 5
	fmt.Println("原來的資料:", number)
	// 2. 取得變數的記憶體位置：&變數名稱
	fmt.Println("資料的記憶體位置:", &number)
	// 3. 存到指標變數。注意指標資料型態： *資料型態
	var numberPointer = &number
	fmt.Println("資料的記憶體位置:", numberPointer)
	// 4. 反解指標變數：*指標變數名稱
	fmt.Println("反解指標回原來的資料:", *numberPointer)

	//var message = "Hello"
	//fmt.Println(message)
	//var messagePointer = &message
	//fmt.Println(messagePointer)
	//fmt.Println(*messagePointer)
}
