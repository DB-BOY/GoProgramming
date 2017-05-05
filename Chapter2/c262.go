package main

import (
	"os"
	"io"
)

/*
defer语句的调用是先进后出的原则，
 最后一个defer语句将最先被执行。
 只不过，当你需要为defer语句到底哪个先执行这种细节而的时候，
 说明你的代码 构可能需要调整一下了。
 */
func main() {
	CopyFile("/Users/pro/test/src.ttt", "/Users/pro/test/src.t")
}

func  CopyFile(det,src string)  (w int64,err error){
	srcFile, err := os.Open(src)
	if err !=nil {
		return
	}
	defer srcFile.Close()

	dstFile,err :=os.Create(det)
	if err !=nil{
		return
	}
	defer dstFile.Close()
	return io.Copy(dstFile,srcFile)

}
