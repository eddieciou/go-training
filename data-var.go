package main // 可執行的程式必須使用 main 封包
import "fmt"

// 載入內建的 fmt 封包，用來做基本的輸入輸出

func main() { // 建立 main 函式， 程式的進入點
	// 程式執行時，從這邊開始
	// 輸出訊息到終端機: fmt.Println(資料,資料,資料, ...)
	/*fmt.Println(3)      // 整數 int
	fmt.Println(3.1415) // 浮點數 float64
	fmt.Println("測試中文") // 字串 string
	fmt.Println(true)   // 布林值 bool
	fmt.Println('a')    // 字符 rune
	fmt.Println(3, 3.1415, "測試中文", true, 'a')*/

	// 變數的使用
	var x int      // 宣告變數 x , 資料型態 int
	x = 4          // 指定資料：把資料4放進變數x裡
	fmt.Println(x) // 使用變數： 用變數的名稱代替資料，做操作
	x = 10         // 指定新的資料：把資料10放進變數，覆蓋原來的資料
	fmt.Println(x)
	//x = "Hello" // 此行會顯示錯誤，變數只能存放原本指定的資料型態的資料
	var f = 3.1415 // 宣告變數時，可以順便把資料放進變數
	fmt.Println(f)
	var s = "哈摟 Golang"
	fmt.Println(s)
	var test = true
	//goland:noinspection GoBoolExpressions
	fmt.Println(test)
	var c = 'b'
	fmt.Println(c)
}
