package main

import "fmt"

var cnt int = 0

func IncreaseAndReturn() int {
	fmt.Println("IncreaseAndReturn()", cnt)
	cnt++
	return cnt

}

func main() {

	// && 연산자 이기 때문에 false 바로 반환 함수 호출 x , 출력 x
	if false && IncreaseAndReturn() < 5 {
		fmt.Println("1증가")
	}

	// || 연산자 이기 때문에 true 바로 반환 함수 호출 x , 출력 o
	if true || IncreaseAndReturn() < 5 {
		fmt.Println("2증가")
	}

	fmt.Println("cnt:", cnt)

}
