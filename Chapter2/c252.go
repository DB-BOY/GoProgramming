package main

//函数调用
import (
	"github.com/user/GoProgramming/Chapter2/mymath"
	"fmt"
)

func main() {

	c, error := mymath.Add(-1, 2);
	if error == nil {
		fmt.Println(c)
	} else {
		fmt.Println(error)
	}
}
