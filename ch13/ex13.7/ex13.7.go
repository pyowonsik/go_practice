package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	A int8
	B int // int64
	C int8
	D int // int64
	E int8
}

func main() {

	user := User{1, 2, 3, 4, 5}
	fmt.Println(unsafe.Sizeof(user))

}
