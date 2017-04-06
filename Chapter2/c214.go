package main

import "fmt"

func GetName() (firstName, lastName, nickName string) {
	return "may", "chan", "chibi maruko"
}

func main() {
	_, _, nickName := GetName()
	fmt.Println(nickName)

}
