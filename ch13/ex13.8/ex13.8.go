package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	A int8
	B int8
	C int8
	D int
	E int
}

func main() {

	var user = User{1, 2, 3, 4, 5}

	fmt.Println(unsafe.Sizeof(user))

}
