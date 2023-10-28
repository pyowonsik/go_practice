package main

import "fmt"

func getMyAge() int {
	return 22
}

func main() {

	// 변수 초기후 비교
	switch age := getMyAge(); true {
	case age < 10:
		fmt.Println("Child")
	case age < 20:
		fmt.Println("Teenager")
	case age < 30:
		fmt.Println("20s")
	default:
		fmt.Println("My age is", age)

	}

	// fmt.Println("age is",age)

}
