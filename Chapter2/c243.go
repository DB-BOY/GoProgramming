package main

import "fmt"

/*
只支持for
不支持while ，do-while
 */
func main() {

	sum := 0
	//for(;;)简化
	for {
		sum++
		if sum > 80 {
			break;
		}
	}

	a := []int{1, 2, 3, 4, 5, 6}
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	//for i,j:=0,len(a);i<j ;i=i+1  {
	//	fmt.Println(a[i])
	//}

	//break 中断循环
JLoop:
	for j := 0; j < 5; j++ {

	FLoop:
		for i := 0; i < 10; i++ {

			fmt.Print(j, " ", i, "\n")
			if i > 5 {
				break FLoop
			}
			if j > 4 {
				break JLoop
			}
		}
	}
	fmt.Println("d")
}
