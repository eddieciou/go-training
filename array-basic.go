package main

import "fmt"

func main() {

	/* // 整數陣列
	var numberArray [3]int
	numberArray[0] = 15
	numberArray[1] = 30
	numberArray[2] = 18
	// numberArray[3] = 18 不能超過陣列的長度，陣列的長度在宣告時就已固定
	fmt.Println(numberArray)
	fmt.Println(numberArray[1]*10)

	// 字串陣列
	var nameArray = [2]string{"Jhon", "Mary"}
	fmt.Println(nameArray)
	// 取得陣列長度
	fmt.Println(len(numberArray))
	fmt.Println(len(nameArray)) */

	// 逐一取得陣列中的資料(遍歷陣列中的資料)
	/* // 計算成績的總和
	var gradeArray = [4]int{60, 90, 75, 10}
	var sum int
	var index int

	for index=0; index < len(gradeArray); index++ {
		sum += gradeArray[index]
	}

	fmt.Println("成績總和:", sum)
	fmt.Println("成績平均:", sum / len(gradeArray)) */

	// 由使用者自行輸入分數，計算平均
	var gradeArray [4]int
	var index int
	var sum int

	for index = 0; index < len(gradeArray); index++ {
		fmt.Print("請輸入第", index+1, "位學生的成績:")
		fmt.Scanln(&gradeArray[index])
		sum += gradeArray[index]
	}

	fmt.Println("成績總和:", sum)
	fmt.Println("成績平均:", sum/len(gradeArray))

}
