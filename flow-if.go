package main

import "fmt"

func main() {
	// 基本語法練習
	//var userSelect string
	//fmt.Print("是否執行Go(y/n)：")
	//fmt.Scanln(&userSelect)
	//if userSelect == "y" {
	//	fmt.Println("Go")
	//}else if userSelect == "n"  {
	//	fmt.Println("Not Go")
	//}else {
	//	fmt.Println("Valid Input")
	//}

	// 簡易情境： ATM 檢測使用者輸入的金額是否適當
	var money int
	fmt.Print("請問想領多少錢?：")
	fmt.Scanln(&money)
	if money < 100 {
		fmt.Println("Too Few")
	} else if money < 100000 {
		fmt.Println("OK")
	} else {
		fmt.Println("Too Much")
	}
	fmt.Println("執行完畢")
}
