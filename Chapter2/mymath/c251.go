package mymath

import "errors"

func Add(a int, b int) (ret int, err error) {

	if a < 0 || b < 0 {
		err = errors.New("支持吃非负数相加")
		return
	}
	return a + b, nil
}
