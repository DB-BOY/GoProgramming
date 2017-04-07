package main

import "fmt"

//不定参数

func main() {
	myFunc(1, 2, 4, 5, 6)
	myFunc(1, 2, 2, 2)

	myFunc2([]int{1, 2, 4, 5, 6})
	myFunc2([]int{1, 2, 2, 2})

}

func myFunc(args ...int) {
	for _, arg := range args {
		fmt.Print(arg, " ")

	}
	fmt.Println()

}

func myFunc2(args []int) {

	for _, arg := range args {
		fmt.Print(arg, " ")
	}
	fmt.Println()
}
