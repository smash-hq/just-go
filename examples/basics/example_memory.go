package basics

import (
	"fmt"
	"unsafe"
)

type T struct {
	a int8
	b int64
	c int32
	d int16
}

type T2 struct {
	a int8
	b int16
	c int32
	d int64
}

func Memory() {
	t := T{}
	t2 := T2{}
	fmt.Println(unsafe.Sizeof(t))
	fmt.Println(unsafe.Sizeof(t2))
}
