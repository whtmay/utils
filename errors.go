package utils

import (
	"fmt"
)

func Err(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
func Assert(ok bool, info string) {
	if !ok {
		fmt.Println(info)
	}
}

