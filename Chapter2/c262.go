package main

import (
	"os"
	"io"
	"fmt"
)

/*
defer语句的调用是先进后出的原则，
 最后一个defer语句将最先被执行。
 只不过，当你需要为defer语句到底哪个先执行这种细节而的时候，
 说明你的代码 构可能需要调整一下了。
 */
func main() {
	//CopyFile("/Users/pro/test/src.ttt", "/Users/pro/test/src.t")
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
}

func CopyFile(det, src string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return
	}
	defer srcFile.Close()

	dstFile, err := os.Create(det)
	if err != nil {
		return
	}
	defer dstFile.Close()
	return io.Copy(dstFile, srcFile)

}

func f1() (result int) {
	defer func() {
		result++
	}()
	return 0
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}