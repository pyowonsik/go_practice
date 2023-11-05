package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	Age   int32   // 4바이트
	Score float64 // 8바이트
}

func main() {

	user := User{23, 77.2}

	fmt.Println(unsafe.Sizeof(user))
	// 메모리 정렬 : 64비트 컴퓨터이기 때문에 8바이트씩 읽는것이 효과적임 따라서, 4바이트의 메모리 패딩을 주고 8의 배수에 맞춰줌.
}
