package main

import "fmt"

func main() {

	noName(3, 3)
}

//匿名函数
func noName(x, y int) {
	f := func(x, y int) int {
		return x + y
	}
	fmt.Println(f(x, y))

	//func(ch chan int) {
	//	ch <- ACK
	//}(reply_chan)  //花括号后直接跟参数列表表示函数调用
}
