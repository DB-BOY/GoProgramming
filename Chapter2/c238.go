package main

import "fmt"

func main() {
	//mySlice1 := make([]int, 5)
	//mySlice2 := make([]int, 5, 20)
	mySlice3 := []int{1, 2, 3, 4, 5}

	//遍历

	for i := 0; i < len(mySlice3); i++ {
		fmt.Println("mySlice[", i, "] = ", mySlice3[i])
	}
	fmt.Println()
	for i, v := range mySlice3 {
		fmt.Println("mySlice[", i, "] = ", v)
	}

	oldSlice := []int{1, 2, 3, 4, 5}
	newSlice := oldSlice[:3]

	fmt.Println("len; ", len(newSlice))
	fmt.Println("cap; ", cap(newSlice))

}
