package main

import "fmt"

func main() {
	a := 3
	switch a {
	case 1:
		fmt.Println("a == 1")
	case 2:
		fmt.Println("a == 2")
	case 3:
		fmt.Println("a == 3")
		fallthrough
	case 4:
		fmt.Println("a == 4")
		fallthrough
	default:
		fmt.Println("a > 4")
	}
}
