package main // 可執行的程式必須使用 main 封包
import "fmt"

// 載入內建的 fmt 封包，用來做基本的輸入輸出

func main() { // 建立 main 函式， 程式的進入點
	// 程式執行時，從這邊開始
	// 輸出訊息到終端機: fmt.Println(資料,資料,資料, ...)
	fmt.Println(3)      // 整數 int
	fmt.Println(3.1415) // 浮點數 float64
	fmt.Println("測試中文") // 字串 string
	fmt.Println(true)   // 布林值 bool
	fmt.Println('a')    // 字符 rune
	fmt.Println(3, 3.1415, "測試中文", true, 'a')
}
