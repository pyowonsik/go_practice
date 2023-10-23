package main

import "fmt"

func main() {
	var a int
	var b int
	n, err := fmt.Scan(&a, &b) // a,b 주소를 입력으로 넘겨야함
	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b)
	}
}
