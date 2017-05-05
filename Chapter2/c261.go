package main


///异常内容详见第三章

func main()  {

}

type PathError struct {
	Op string
	Path string
	Err error
}

func (e *PathError) Error() string{
	return e.Op+" "+e.Path+":"+e.Err.Error()
}
