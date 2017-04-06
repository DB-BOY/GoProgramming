package main

import "fmt"

func main() {

	fmt.Println(example(3))
}

func example(x int) int {
	if x == 0 {
		return 5
	} else {
		return x
	}
}
