package main

import "fmt"

func main() {
	const C int = 10

	var b int = C * 20
	fmt.Println(b)
	// C = 20
	// fmt.Println(&C)
}
