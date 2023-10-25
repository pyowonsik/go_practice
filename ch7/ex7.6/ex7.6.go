package main

import "fmt"

// 재귀 호출 함수 안에서 자기 자신 함수를 다시 호출

func printNo(n int) {
	if n == 0 {
		return
	}
	fmt.Println(n)
	printNo(n - 1)
	fmt.Println("After", n)
}

func main() {
	printNo(3)
}
