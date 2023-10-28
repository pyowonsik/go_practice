package main

import "fmt"

func getMyAge() int {
	return 22
}

func main() {

	// 변수 초기후 비교
	switch age := getMyAge(); age {
	case 10:
		fmt.Println("10대 입니다.")
	case 33:
		fmt.Println("Pair 3")
	default:
		fmt.Println("My age is", age)

	}

	// fmt.Println("age is",age)

}
