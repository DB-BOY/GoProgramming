package main

import "fmt"

func main() {
	selectExm(1)
	selectExm(2)
	selectExm(5)
}

//fallthrough关键字，继续执行下一个case;
func selectExm(x int) {
	switch x {
	case 0:
		fmt.Println("0")
	case 1:
		fmt.Println(1)
	case 2:
		fallthrough
	case 3:
		fmt.Println("3")

	case 4, 5, 6:
		fmt.Println("4.5.6")

	default:
		fmt.Println("default")

	}
	//表达式形式
	switch {
	case 0 <= x && x <= 3:
		fmt.Println("0-3")
	case 4 <= x && x <= 6:
		fmt.Println("4-6")
	default:
		fmt.Println("other")
	}

	fmt.Println("*************")
}
