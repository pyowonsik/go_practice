package main

import "fmt"

func main() {

	// 범위가 큰타입에서 작은 타입으로 변환
	var aaa int16 = 3456
	var ccc int8 = int8(aaa)
	fmt.Println("범위가 큰타입에서 작은 타입으로 변환")
	fmt.Println(aaa, ccc)
}
