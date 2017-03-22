package main

import "C"
import "unsafe"

//后续会介绍，这个没跑通

//# command-line-arguments
//could not determine kind of name for C.free
//could not determine kind of name for C.puts
func main() {
	cstr := C.CString("hello world")
	C.puts(cstr)
	C.free(unsafe.Pointer(cstr))
}
